//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdelegatednetwork

import (
	"encoding/json"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ControllerResource.
func (c ControllerResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", c.ID)
	populate(objectMap, "location", c.Location)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "tags", c.Tags)
	populate(objectMap, "type", c.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ControllerResourceUpdateParameters.
func (c ControllerResourceUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", c.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DelegatedController.
func (d DelegatedController) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DelegatedSubnet.
func (d DelegatedSubnet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DelegatedSubnetResource.
func (d DelegatedSubnetResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", d.ID)
	populate(objectMap, "location", d.Location)
	populate(objectMap, "name", d.Name)
	populate(objectMap, "tags", d.Tags)
	populate(objectMap, "type", d.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Orchestrator.
func (o Orchestrator) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", o.ID)
	populate(objectMap, "identity", o.Identity)
	populate(objectMap, "kind", o.Kind)
	populate(objectMap, "location", o.Location)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "properties", o.Properties)
	populate(objectMap, "tags", o.Tags)
	populate(objectMap, "type", o.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OrchestratorResource.
func (o OrchestratorResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", o.ID)
	populate(objectMap, "identity", o.Identity)
	populate(objectMap, "kind", o.Kind)
	populate(objectMap, "location", o.Location)
	populate(objectMap, "name", o.Name)
	populate(objectMap, "tags", o.Tags)
	populate(objectMap, "type", o.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OrchestratorResourceUpdateParameters.
func (o OrchestratorResourceUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", o.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceUpdateParameters.
func (r ResourceUpdateParameters) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", r.Tags)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}
