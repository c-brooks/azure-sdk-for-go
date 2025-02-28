//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armelasticsan

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type ElasticSan.
func (e ElasticSan) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", e.ID)
	populate(objectMap, "location", e.Location)
	populate(objectMap, "name", e.Name)
	populate(objectMap, "properties", e.Properties)
	populate(objectMap, "systemData", e.SystemData)
	populate(objectMap, "tags", e.Tags)
	populate(objectMap, "type", e.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ElasticSan.
func (e *ElasticSan) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &e.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &e.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &e.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &e.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &e.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &e.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &e.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type IscsiTargetInfo.
func (i IscsiTargetInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "status", i.Status)
	populate(objectMap, "targetIqn", i.TargetIqn)
	populate(objectMap, "targetPortalHostname", i.TargetPortalHostname)
	populate(objectMap, "targetPortalPort", i.TargetPortalPort)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type IscsiTargetInfo.
func (i *IscsiTargetInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &i.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &i.Status)
			delete(rawMsg, key)
		case "targetIqn":
			err = unpopulate(val, "TargetIqn", &i.TargetIqn)
			delete(rawMsg, key)
		case "targetPortalHostname":
			err = unpopulate(val, "TargetPortalHostname", &i.TargetPortalHostname)
			delete(rawMsg, key)
		case "targetPortalPort":
			err = unpopulate(val, "TargetPortalPort", &i.TargetPortalPort)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type List.
func (l List) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type List.
func (l *List) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &l.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &l.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRuleSet.
func (n NetworkRuleSet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "virtualNetworkRules", n.VirtualNetworkRules)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkRuleSet.
func (n *NetworkRuleSet) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "virtualNetworkRules":
			err = unpopulate(val, "VirtualNetworkRules", &n.VirtualNetworkRules)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationDisplay.
func (o OperationDisplay) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", o.Description)
	populate(objectMap, "operation", o.Operation)
	populate(objectMap, "provider", o.Provider)
	populate(objectMap, "resource", o.Resource)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationDisplay.
func (o *OperationDisplay) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &o.Description)
			delete(rawMsg, key)
		case "operation":
			err = unpopulate(val, "Operation", &o.Operation)
			delete(rawMsg, key)
		case "provider":
			err = unpopulate(val, "Provider", &o.Provider)
			delete(rawMsg, key)
		case "resource":
			err = unpopulate(val, "Resource", &o.Resource)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type OperationListResult.
func (o OperationListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", o.NextLink)
	populate(objectMap, "value", o.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type OperationListResult.
func (o *OperationListResult) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", o, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &o.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &o.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", o, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Properties.
func (p Properties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "availabilityZones", p.AvailabilityZones)
	populate(objectMap, "baseSizeTiB", p.BaseSizeTiB)
	populate(objectMap, "extendedCapacitySizeTiB", p.ExtendedCapacitySizeTiB)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "sku", p.SKU)
	populate(objectMap, "totalIops", p.TotalIops)
	populate(objectMap, "totalMBps", p.TotalMBps)
	populate(objectMap, "totalSizeTiB", p.TotalSizeTiB)
	populate(objectMap, "totalVolumeSizeGiB", p.TotalVolumeSizeGiB)
	populate(objectMap, "volumeGroupCount", p.VolumeGroupCount)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Properties.
func (p *Properties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", p, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "availabilityZones":
			err = unpopulate(val, "AvailabilityZones", &p.AvailabilityZones)
			delete(rawMsg, key)
		case "baseSizeTiB":
			err = unpopulate(val, "BaseSizeTiB", &p.BaseSizeTiB)
			delete(rawMsg, key)
		case "extendedCapacitySizeTiB":
			err = unpopulate(val, "ExtendedCapacitySizeTiB", &p.ExtendedCapacitySizeTiB)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &p.ProvisioningState)
			delete(rawMsg, key)
		case "sku":
			err = unpopulate(val, "SKU", &p.SKU)
			delete(rawMsg, key)
		case "totalIops":
			err = unpopulate(val, "TotalIops", &p.TotalIops)
			delete(rawMsg, key)
		case "totalMBps":
			err = unpopulate(val, "TotalMBps", &p.TotalMBps)
			delete(rawMsg, key)
		case "totalSizeTiB":
			err = unpopulate(val, "TotalSizeTiB", &p.TotalSizeTiB)
			delete(rawMsg, key)
		case "totalVolumeSizeGiB":
			err = unpopulate(val, "TotalVolumeSizeGiB", &p.TotalVolumeSizeGiB)
			delete(rawMsg, key)
		case "volumeGroupCount":
			err = unpopulate(val, "VolumeGroupCount", &p.VolumeGroupCount)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", p, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RPOperation.
func (r RPOperation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "display", r.Display)
	populate(objectMap, "isDataAction", r.IsDataAction)
	populate(objectMap, "name", r.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RPOperation.
func (r *RPOperation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "display":
			err = unpopulate(val, "Display", &r.Display)
			delete(rawMsg, key)
		case "isDataAction":
			err = unpopulate(val, "IsDataAction", &r.IsDataAction)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &r.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKU.
func (s SKU) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", s.Name)
	populate(objectMap, "tier", s.Tier)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKU.
func (s *SKU) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "tier":
			err = unpopulate(val, "Tier", &s.Tier)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKUCapability.
func (s SKUCapability) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", s.Name)
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKUCapability.
func (s *SKUCapability) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKUInformation.
func (s SKUInformation) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capabilities", s.Capabilities)
	populate(objectMap, "locationInfo", s.LocationInfo)
	populate(objectMap, "locations", s.Locations)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "resourceType", s.ResourceType)
	populate(objectMap, "tier", s.Tier)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKUInformation.
func (s *SKUInformation) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "capabilities":
			err = unpopulate(val, "Capabilities", &s.Capabilities)
			delete(rawMsg, key)
		case "locationInfo":
			err = unpopulate(val, "LocationInfo", &s.LocationInfo)
			delete(rawMsg, key)
		case "locations":
			err = unpopulate(val, "Locations", &s.Locations)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "resourceType":
			err = unpopulate(val, "ResourceType", &s.ResourceType)
			delete(rawMsg, key)
		case "tier":
			err = unpopulate(val, "Tier", &s.Tier)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKUInformationList.
func (s SKUInformationList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "value", s.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKUInformationList.
func (s *SKUInformationList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "value":
			err = unpopulate(val, "Value", &s.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SKULocationInfo.
func (s SKULocationInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "location", s.Location)
	populate(objectMap, "zones", s.Zones)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SKULocationInfo.
func (s *SKULocationInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "location":
			err = unpopulate(val, "Location", &s.Location)
			delete(rawMsg, key)
		case "zones":
			err = unpopulate(val, "Zones", &s.Zones)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SourceCreationData.
func (s SourceCreationData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["createSource"] = "None"
	populate(objectMap, "sourceUri", s.SourceURI)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SourceCreationData.
func (s *SourceCreationData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createSource":
			err = unpopulate(val, "CreateSource", &s.CreateSource)
			delete(rawMsg, key)
		case "sourceUri":
			err = unpopulate(val, "SourceURI", &s.SourceURI)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
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

// MarshalJSON implements the json.Marshaller interface for type Update.
func (u Update) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", u.Properties)
	populate(objectMap, "tags", u.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Update.
func (u *Update) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &u.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &u.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type UpdateProperties.
func (u UpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "baseSizeTiB", u.BaseSizeTiB)
	populate(objectMap, "extendedCapacitySizeTiB", u.ExtendedCapacitySizeTiB)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UpdateProperties.
func (u *UpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "baseSizeTiB":
			err = unpopulate(val, "BaseSizeTiB", &u.BaseSizeTiB)
			delete(rawMsg, key)
		case "extendedCapacitySizeTiB":
			err = unpopulate(val, "ExtendedCapacitySizeTiB", &u.ExtendedCapacitySizeTiB)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VirtualNetworkRule.
func (v VirtualNetworkRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["action"] = "Allow"
	populate(objectMap, "state", v.State)
	populate(objectMap, "id", v.VirtualNetworkResourceID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VirtualNetworkRule.
func (v *VirtualNetworkRule) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "action":
			err = unpopulate(val, "Action", &v.Action)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &v.State)
			delete(rawMsg, key)
		case "id":
			err = unpopulate(val, "VirtualNetworkResourceID", &v.VirtualNetworkResourceID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Volume.
func (v Volume) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "systemData", v.SystemData)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type Volume.
func (v *Volume) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &v.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &v.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &v.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &v.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &v.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &v.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroup.
func (v VolumeGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "systemData", v.SystemData)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeGroup.
func (v *VolumeGroup) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &v.ID)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &v.Name)
			delete(rawMsg, key)
		case "properties":
			err = unpopulate(val, "Properties", &v.Properties)
			delete(rawMsg, key)
		case "systemData":
			err = unpopulate(val, "SystemData", &v.SystemData)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &v.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &v.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupList.
func (v VolumeGroupList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeGroupList.
func (v *VolumeGroupList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &v.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &v.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupProperties.
func (v VolumeGroupProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "encryption", v.Encryption)
	populate(objectMap, "networkAcls", v.NetworkACLs)
	populate(objectMap, "protocolType", v.ProtocolType)
	populate(objectMap, "provisioningState", v.ProvisioningState)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeGroupProperties.
func (v *VolumeGroupProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "encryption":
			err = unpopulate(val, "Encryption", &v.Encryption)
			delete(rawMsg, key)
		case "networkAcls":
			err = unpopulate(val, "NetworkACLs", &v.NetworkACLs)
			delete(rawMsg, key)
		case "protocolType":
			err = unpopulate(val, "ProtocolType", &v.ProtocolType)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &v.ProvisioningState)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupUpdate.
func (v VolumeGroupUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeGroupUpdate.
func (v *VolumeGroupUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &v.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &v.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeGroupUpdateProperties.
func (v VolumeGroupUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "encryption", v.Encryption)
	populate(objectMap, "networkAcls", v.NetworkACLs)
	populate(objectMap, "protocolType", v.ProtocolType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeGroupUpdateProperties.
func (v *VolumeGroupUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "encryption":
			err = unpopulate(val, "Encryption", &v.Encryption)
			delete(rawMsg, key)
		case "networkAcls":
			err = unpopulate(val, "NetworkACLs", &v.NetworkACLs)
			delete(rawMsg, key)
		case "protocolType":
			err = unpopulate(val, "ProtocolType", &v.ProtocolType)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeList.
func (v VolumeList) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", v.NextLink)
	populate(objectMap, "value", v.Value)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeList.
func (v *VolumeList) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "nextLink":
			err = unpopulate(val, "NextLink", &v.NextLink)
			delete(rawMsg, key)
		case "value":
			err = unpopulate(val, "Value", &v.Value)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeProperties.
func (v VolumeProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "creationData", v.CreationData)
	populate(objectMap, "sizeGiB", v.SizeGiB)
	populate(objectMap, "storageTarget", v.StorageTarget)
	populate(objectMap, "volumeId", v.VolumeID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeProperties.
func (v *VolumeProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "creationData":
			err = unpopulate(val, "CreationData", &v.CreationData)
			delete(rawMsg, key)
		case "sizeGiB":
			err = unpopulate(val, "SizeGiB", &v.SizeGiB)
			delete(rawMsg, key)
		case "storageTarget":
			err = unpopulate(val, "StorageTarget", &v.StorageTarget)
			delete(rawMsg, key)
		case "volumeId":
			err = unpopulate(val, "VolumeID", &v.VolumeID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeUpdate.
func (v VolumeUpdate) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeUpdate.
func (v *VolumeUpdate) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "properties":
			err = unpopulate(val, "Properties", &v.Properties)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &v.Tags)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", v, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type VolumeUpdateProperties.
func (v VolumeUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "sizeGiB", v.SizeGiB)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type VolumeUpdateProperties.
func (v *VolumeUpdateProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", v, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "sizeGiB":
			err = unpopulate(val, "SizeGiB", &v.SizeGiB)
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

func unpopulate(data json.RawMessage, fn string, v interface{}) error {
	if data == nil {
		return nil
	}
	if err := json.Unmarshal(data, v); err != nil {
		return fmt.Errorf("struct field %s: %v", fn, err)
	}
	return nil
}
