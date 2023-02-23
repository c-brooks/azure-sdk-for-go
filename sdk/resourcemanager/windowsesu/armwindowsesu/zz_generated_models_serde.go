//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armwindowsesu

import (
	"encoding/json"
	"fmt"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type MultipleActivationKey.
func (m MultipleActivationKey) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "location", m.Location)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "tags", m.Tags)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MultipleActivationKeyProperties.
func (m MultipleActivationKeyProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "agreementNumber", m.AgreementNumber)
	populateTimeRFC3339(objectMap, "expirationDate", m.ExpirationDate)
	populate(objectMap, "installedServerNumber", m.InstalledServerNumber)
	populate(objectMap, "isEligible", m.IsEligible)
	populate(objectMap, "multipleActivationKey", m.MultipleActivationKey)
	populate(objectMap, "osType", m.OSType)
	populate(objectMap, "provisioningState", m.ProvisioningState)
	populate(objectMap, "supportType", m.SupportType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MultipleActivationKeyProperties.
func (m *MultipleActivationKeyProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "agreementNumber":
			err = unpopulate(val, "AgreementNumber", &m.AgreementNumber)
			delete(rawMsg, key)
		case "expirationDate":
			err = unpopulateTimeRFC3339(val, "ExpirationDate", &m.ExpirationDate)
			delete(rawMsg, key)
		case "installedServerNumber":
			err = unpopulate(val, "InstalledServerNumber", &m.InstalledServerNumber)
			delete(rawMsg, key)
		case "isEligible":
			err = unpopulate(val, "IsEligible", &m.IsEligible)
			delete(rawMsg, key)
		case "multipleActivationKey":
			err = unpopulate(val, "MultipleActivationKey", &m.MultipleActivationKey)
			delete(rawMsg, key)
		case "osType":
			err = unpopulate(val, "OSType", &m.OSType)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &m.ProvisioningState)
			delete(rawMsg, key)
		case "supportType":
			err = unpopulate(val, "SupportType", &m.SupportType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MultipleActivationKeyUpdate.
func (m MultipleActivationKeyUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", m.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TrackedResource.
func (t TrackedResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
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

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
