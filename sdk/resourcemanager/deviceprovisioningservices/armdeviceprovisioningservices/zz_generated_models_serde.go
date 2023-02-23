//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceprovisioningservices

import (
	"encoding/json"
	"fmt"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CertificateProperties.
func (c CertificateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateByteArray(objectMap, "certificate", c.Certificate, runtime.Base64StdFormat)
	populateTimeRFC1123(objectMap, "created", c.Created)
	populateTimeRFC1123(objectMap, "expiry", c.Expiry)
	populate(objectMap, "isVerified", c.IsVerified)
	populate(objectMap, "subject", c.Subject)
	populate(objectMap, "thumbprint", c.Thumbprint)
	populateTimeRFC1123(objectMap, "updated", c.Updated)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type CertificateProperties.
func (c *CertificateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificate":
			err = runtime.DecodeByteArray(string(val), &c.Certificate, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "created":
			err = unpopulateTimeRFC1123(val, "Created", &c.Created)
			delete(rawMsg, key)
		case "expiry":
			err = unpopulateTimeRFC1123(val, "Expiry", &c.Expiry)
			delete(rawMsg, key)
		case "isVerified":
			err = unpopulate(val, "IsVerified", &c.IsVerified)
			delete(rawMsg, key)
		case "subject":
			err = unpopulate(val, "Subject", &c.Subject)
			delete(rawMsg, key)
		case "thumbprint":
			err = unpopulate(val, "Thumbprint", &c.Thumbprint)
			delete(rawMsg, key)
		case "updated":
			err = unpopulateTimeRFC1123(val, "Updated", &c.Updated)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type IotDpsPropertiesDescription.
func (i IotDpsPropertiesDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allocationPolicy", i.AllocationPolicy)
	populate(objectMap, "authorizationPolicies", i.AuthorizationPolicies)
	populate(objectMap, "deviceProvisioningHostName", i.DeviceProvisioningHostName)
	populate(objectMap, "enableDataResidency", i.EnableDataResidency)
	populate(objectMap, "idScope", i.IDScope)
	populate(objectMap, "ipFilterRules", i.IPFilterRules)
	populate(objectMap, "iotHubs", i.IotHubs)
	populate(objectMap, "privateEndpointConnections", i.PrivateEndpointConnections)
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "publicNetworkAccess", i.PublicNetworkAccess)
	populate(objectMap, "serviceOperationsHostName", i.ServiceOperationsHostName)
	populate(objectMap, "state", i.State)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProvisioningServiceDescription.
func (p ProvisioningServiceDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", p.Etag)
	populate(objectMap, "id", p.ID)
	populate(objectMap, "location", p.Location)
	populate(objectMap, "name", p.Name)
	populate(objectMap, "properties", p.Properties)
	populate(objectMap, "sku", p.SKU)
	populate(objectMap, "systemData", p.SystemData)
	populate(objectMap, "tags", p.Tags)
	populate(objectMap, "type", p.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Resource.
func (r Resource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", r.ID)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, "CreatedAt", &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, "CreatedBy", &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, "CreatedByType", &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, "LastModifiedAt", &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, "LastModifiedBy", &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, "LastModifiedByType", &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type TagsResource.
func (t TagsResource) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "tags", t.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VerificationCodeResponseProperties.
func (v *VerificationCodeResponseProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "certificate":
			err = runtime.DecodeByteArray(string(val), &v.Certificate, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "created":
			err = unpopulate(val, "Created", &v.Created)
			delete(rawMsg, key)
		case "expiry":
			err = unpopulate(val, "Expiry", &v.Expiry)
			delete(rawMsg, key)
		case "isVerified":
			err = unpopulate(val, "IsVerified", &v.IsVerified)
			delete(rawMsg, key)
		case "subject":
			err = unpopulate(val, "Subject", &v.Subject)
			delete(rawMsg, key)
		case "thumbprint":
			err = unpopulate(val, "Thumbprint", &v.Thumbprint)
			delete(rawMsg, key)
		case "updated":
			err = unpopulate(val, "Updated", &v.Updated)
			delete(rawMsg, key)
		case "verificationCode":
			err = unpopulate(val, "VerificationCode", &v.VerificationCode)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
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

func populateByteArray(m map[string]interface{}, k string, b []byte, f runtime.Base64Encoding) {
	if azcore.IsNullValue(b) {
		m[k] = nil
	} else if len(b) == 0 {
		return
	} else {
		m[k] = runtime.EncodeByteArray(b, f)
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
