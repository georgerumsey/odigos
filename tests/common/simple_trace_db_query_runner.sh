#!/bin/bash

# Ensure the script fails if any command fails
set -e

# Function to verify the YAML schema
function verify_yaml_schema() {
  local file=$1
  local query=$(yq e '.query' "$file")
  local expected_count=$(yq e '.expected.count' "$file")
  local minimum_count=$(yq e '.expected.minimum' "$file")

  if [[ -z "$query" || ("$expected_count" == "null" && "$minimum_count" == "null") || (-z "$minimum_count" && -z "$expected_count") ]]; then
    echo "Invalid YAML schema in file: $file"
    exit 1
  fi
}

function urlencode() (
  local length="${#1}"
  for ((i = 0; i < length; i++)); do
    local c="${1:i:1}"
    case $c in
    [a-zA-Z0-9.~_-]) printf "$c" ;;
    *) printf '%%%02X' "'$c" ;;
    esac
  done
)

# Function to process a YAML file
function process_yaml_file() {
  # The syntax of "${1:-traces}" means that if the first argument is not provided, it will use the default value "traces"
  local dest_namespace="${3:-traces}"
  local dest_service="${4:-simple-trace-db}"
  local dest_port="4318"
  local verbose="${2}"

  local file=$1
  file_name=$(basename "$file")
  echo "Running test $file_name"
  echo "Dest namespace: $dest_namespace"
  echo "Dest service: $dest_service"
  echo "Dest port: $dest_port"

  query=$(yq '.query' "$file")
  encoded_query=$(urlencode "$query")
  expected_count=$(yq e '.expected.count' "$file")
  minimum_count=$(yq e '.expected.minimum' "$file")

  response=$(kubectl get --raw /api/v1/namespaces/$dest_namespace/services/$dest_service:$dest_port/proxy/v1/traces\?jmespath=$encoded_query)

  if [ "$verbose" == "true" ]; then
    echo "============== Raw response from trace db ===================="
    echo "$response" | jq .
    echo "========================================================="
  fi

  num_of_traces=$(echo "$response" | jq 'keys | length')

  if [ "$expected_count" != "null" ]; then
    # if num_of_traces not equal to expected_count
    if [ "$num_of_traces" -ne "$expected_count" ]; then
      echo "Test FAILED: expected $expected_count got $num_of_traces"
      exit 1
    else
      echo "Test PASSED: expected $expected_count got $num_of_traces"
      exit 0
    fi
  fi

  if [ "$minimum_count" != "null" ]; then
    if [ "$num_of_traces" -lt "$minimum_count" ]; then
      echo "Test FAILED: expected at least $minimum_count got $num_of_traces"
      exit 1
    else
      echo "Test PASSED: expected at least $minimum_count got $num_of_traces"
      exit 0
    fi
  fi
}

# Check if the first argument is provided
if [ -z "$1" ]; then
  echo "Usage: $0 <simple-trace-db-test-file> [--verbose]"
  exit 1
fi

# Test file path
TEST_FILE=$1
VERBOSE=false
if [ "$2" == "--verbose" ]; then
  VERBOSE=true
fi

# Check if yq is installed
if ! command -v yq &>/dev/null; then
  echo "yq command not found. Please install yq."
  exit 1
fi

DESTINATION_NAMESPACE=$3
DESTINATION_SERVICE=$4

verify_yaml_schema $TEST_FILE
process_yaml_file $TEST_FILE $VERBOSE $DESTINATION_NAMESPACE $DESTINATION_SERVICE
