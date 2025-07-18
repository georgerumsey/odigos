/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	actionsv1alpha1 "github.com/odigos-io/odigos/api/actions/v1alpha1"
	common "github.com/odigos-io/odigos/common"
)

// ActionSpecApplyConfiguration represents a declarative configuration of the ActionSpec type for use
// with apply.
type ActionSpecApplyConfiguration struct {
	ActionName      *string                                `json:"actionName,omitempty"`
	Notes           *string                                `json:"notes,omitempty"`
	Disabled        *bool                                  `json:"disabled,omitempty"`
	Signals         []common.ObservabilitySignal           `json:"signals,omitempty"`
	AddClusterInfo  *actionsv1alpha1.AddClusterInfoConfig  `json:"addClusterInfo,omitempty"`
	DeleteAttribute *actionsv1alpha1.DeleteAttributeConfig `json:"deleteAttribute,omitempty"`
	RenameAttribute *actionsv1alpha1.RenameAttributeConfig `json:"renameAttribute,omitempty"`
	PiiMasking      *actionsv1alpha1.PiiMaskingConfig      `json:"piiMasking,omitempty"`
	K8sAttributes   *actionsv1alpha1.K8sAttributesConfig   `json:"k8sAttributes,omitempty"`
}

// ActionSpecApplyConfiguration constructs a declarative configuration of the ActionSpec type for use with
// apply.
func ActionSpec() *ActionSpecApplyConfiguration {
	return &ActionSpecApplyConfiguration{}
}

// WithActionName sets the ActionName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ActionName field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithActionName(value string) *ActionSpecApplyConfiguration {
	b.ActionName = &value
	return b
}

// WithNotes sets the Notes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Notes field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithNotes(value string) *ActionSpecApplyConfiguration {
	b.Notes = &value
	return b
}

// WithDisabled sets the Disabled field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Disabled field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithDisabled(value bool) *ActionSpecApplyConfiguration {
	b.Disabled = &value
	return b
}

// WithSignals adds the given value to the Signals field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Signals field.
func (b *ActionSpecApplyConfiguration) WithSignals(values ...common.ObservabilitySignal) *ActionSpecApplyConfiguration {
	for i := range values {
		b.Signals = append(b.Signals, values[i])
	}
	return b
}

// WithAddClusterInfo sets the AddClusterInfo field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AddClusterInfo field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithAddClusterInfo(value actionsv1alpha1.AddClusterInfoConfig) *ActionSpecApplyConfiguration {
	b.AddClusterInfo = &value
	return b
}

// WithDeleteAttribute sets the DeleteAttribute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeleteAttribute field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithDeleteAttribute(value actionsv1alpha1.DeleteAttributeConfig) *ActionSpecApplyConfiguration {
	b.DeleteAttribute = &value
	return b
}

// WithRenameAttribute sets the RenameAttribute field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RenameAttribute field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithRenameAttribute(value actionsv1alpha1.RenameAttributeConfig) *ActionSpecApplyConfiguration {
	b.RenameAttribute = &value
	return b
}

// WithPiiMasking sets the PiiMasking field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PiiMasking field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithPiiMasking(value actionsv1alpha1.PiiMaskingConfig) *ActionSpecApplyConfiguration {
	b.PiiMasking = &value
	return b
}

// WithK8sAttributes sets the K8sAttributes field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the K8sAttributes field is set to the value of the last call.
func (b *ActionSpecApplyConfiguration) WithK8sAttributes(value actionsv1alpha1.K8sAttributesConfig) *ActionSpecApplyConfiguration {
	b.K8sAttributes = &value
	return b
}
