//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

import (
	"encoding/json"
	"fmt"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore/runtime"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AgentUpdatePatchProperties.
func (a AgentUpdatePatchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "maintenanceWindowTimeZone", a.MaintenanceWindowTimeZone)
	populate(objectMap, "maintenanceWindows", a.MaintenanceWindows)
	populate(objectMap, "type", a.Type)
	populate(objectMap, "useSessionHostLocalTime", a.UseSessionHostLocalTime)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type AgentUpdateProperties.
func (a AgentUpdateProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "maintenanceWindowTimeZone", a.MaintenanceWindowTimeZone)
	populate(objectMap, "maintenanceWindows", a.MaintenanceWindows)
	populate(objectMap, "type", a.Type)
	populate(objectMap, "useSessionHostLocalTime", a.UseSessionHostLocalTime)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationGroup.
func (a ApplicationGroup) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", a.Etag)
	populate(objectMap, "id", a.ID)
	populate(objectMap, "identity", a.Identity)
	populate(objectMap, "kind", a.Kind)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "managedBy", a.ManagedBy)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "plan", a.Plan)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "sku", a.SKU)
	populate(objectMap, "systemData", a.SystemData)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationGroupPatch.
func (a ApplicationGroupPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationPatch.
func (a ApplicationPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "tags", a.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationProperties.
func (a ApplicationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicationType", a.ApplicationType)
	populate(objectMap, "commandLineArguments", a.CommandLineArguments)
	populate(objectMap, "commandLineSetting", a.CommandLineSetting)
	populate(objectMap, "description", a.Description)
	populate(objectMap, "filePath", a.FilePath)
	populate(objectMap, "friendlyName", a.FriendlyName)
	populateByteArray(objectMap, "iconContent", a.IconContent, runtime.Base64StdFormat)
	populate(objectMap, "iconHash", a.IconHash)
	populate(objectMap, "iconIndex", a.IconIndex)
	populate(objectMap, "iconPath", a.IconPath)
	populate(objectMap, "msixPackageApplicationId", a.MsixPackageApplicationID)
	populate(objectMap, "msixPackageFamilyName", a.MsixPackageFamilyName)
	populate(objectMap, "objectId", a.ObjectID)
	populate(objectMap, "showInPortal", a.ShowInPortal)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ApplicationProperties.
func (a *ApplicationProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "applicationType":
			err = unpopulate(val, "ApplicationType", &a.ApplicationType)
			delete(rawMsg, key)
		case "commandLineArguments":
			err = unpopulate(val, "CommandLineArguments", &a.CommandLineArguments)
			delete(rawMsg, key)
		case "commandLineSetting":
			err = unpopulate(val, "CommandLineSetting", &a.CommandLineSetting)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &a.Description)
			delete(rawMsg, key)
		case "filePath":
			err = unpopulate(val, "FilePath", &a.FilePath)
			delete(rawMsg, key)
		case "friendlyName":
			err = unpopulate(val, "FriendlyName", &a.FriendlyName)
			delete(rawMsg, key)
		case "iconContent":
			err = runtime.DecodeByteArray(string(val), &a.IconContent, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "iconHash":
			err = unpopulate(val, "IconHash", &a.IconHash)
			delete(rawMsg, key)
		case "iconIndex":
			err = unpopulate(val, "IconIndex", &a.IconIndex)
			delete(rawMsg, key)
		case "iconPath":
			err = unpopulate(val, "IconPath", &a.IconPath)
			delete(rawMsg, key)
		case "msixPackageApplicationId":
			err = unpopulate(val, "MsixPackageApplicationID", &a.MsixPackageApplicationID)
			delete(rawMsg, key)
		case "msixPackageFamilyName":
			err = unpopulate(val, "MsixPackageFamilyName", &a.MsixPackageFamilyName)
			delete(rawMsg, key)
		case "objectId":
			err = unpopulate(val, "ObjectID", &a.ObjectID)
			delete(rawMsg, key)
		case "showInPortal":
			err = unpopulate(val, "ShowInPortal", &a.ShowInPortal)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DesktopPatch.
func (d DesktopPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", d.Properties)
	populate(objectMap, "tags", d.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DesktopProperties.
func (d DesktopProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", d.Description)
	populate(objectMap, "friendlyName", d.FriendlyName)
	populateByteArray(objectMap, "iconContent", d.IconContent, runtime.Base64StdFormat)
	populate(objectMap, "iconHash", d.IconHash)
	populate(objectMap, "objectId", d.ObjectID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DesktopProperties.
func (d *DesktopProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &d.Description)
			delete(rawMsg, key)
		case "friendlyName":
			err = unpopulate(val, "FriendlyName", &d.FriendlyName)
			delete(rawMsg, key)
		case "iconContent":
			err = runtime.DecodeByteArray(string(val), &d.IconContent, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "iconHash":
			err = unpopulate(val, "IconHash", &d.IconHash)
			delete(rawMsg, key)
		case "objectId":
			err = unpopulate(val, "ObjectID", &d.ObjectID)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExpandMsixImageProperties.
func (e ExpandMsixImageProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", e.DisplayName)
	populate(objectMap, "imagePath", e.ImagePath)
	populate(objectMap, "isActive", e.IsActive)
	populate(objectMap, "isRegularRegistration", e.IsRegularRegistration)
	populateTimeRFC3339(objectMap, "lastUpdated", e.LastUpdated)
	populate(objectMap, "packageAlias", e.PackageAlias)
	populate(objectMap, "packageApplications", e.PackageApplications)
	populate(objectMap, "packageDependencies", e.PackageDependencies)
	populate(objectMap, "packageFamilyName", e.PackageFamilyName)
	populate(objectMap, "packageFullName", e.PackageFullName)
	populate(objectMap, "packageName", e.PackageName)
	populate(objectMap, "packageRelativePath", e.PackageRelativePath)
	populate(objectMap, "version", e.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ExpandMsixImageProperties.
func (e *ExpandMsixImageProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &e.DisplayName)
			delete(rawMsg, key)
		case "imagePath":
			err = unpopulate(val, "ImagePath", &e.ImagePath)
			delete(rawMsg, key)
		case "isActive":
			err = unpopulate(val, "IsActive", &e.IsActive)
			delete(rawMsg, key)
		case "isRegularRegistration":
			err = unpopulate(val, "IsRegularRegistration", &e.IsRegularRegistration)
			delete(rawMsg, key)
		case "lastUpdated":
			err = unpopulateTimeRFC3339(val, "LastUpdated", &e.LastUpdated)
			delete(rawMsg, key)
		case "packageAlias":
			err = unpopulate(val, "PackageAlias", &e.PackageAlias)
			delete(rawMsg, key)
		case "packageApplications":
			err = unpopulate(val, "PackageApplications", &e.PackageApplications)
			delete(rawMsg, key)
		case "packageDependencies":
			err = unpopulate(val, "PackageDependencies", &e.PackageDependencies)
			delete(rawMsg, key)
		case "packageFamilyName":
			err = unpopulate(val, "PackageFamilyName", &e.PackageFamilyName)
			delete(rawMsg, key)
		case "packageFullName":
			err = unpopulate(val, "PackageFullName", &e.PackageFullName)
			delete(rawMsg, key)
		case "packageName":
			err = unpopulate(val, "PackageName", &e.PackageName)
			delete(rawMsg, key)
		case "packageRelativePath":
			err = unpopulate(val, "PackageRelativePath", &e.PackageRelativePath)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &e.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type HostPool.
func (h HostPool) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", h.Etag)
	populate(objectMap, "id", h.ID)
	populate(objectMap, "identity", h.Identity)
	populate(objectMap, "kind", h.Kind)
	populate(objectMap, "location", h.Location)
	populate(objectMap, "managedBy", h.ManagedBy)
	populate(objectMap, "name", h.Name)
	populate(objectMap, "plan", h.Plan)
	populate(objectMap, "properties", h.Properties)
	populate(objectMap, "sku", h.SKU)
	populate(objectMap, "systemData", h.SystemData)
	populate(objectMap, "tags", h.Tags)
	populate(objectMap, "type", h.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HostPoolPatch.
func (h HostPoolPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", h.ID)
	populate(objectMap, "name", h.Name)
	populate(objectMap, "properties", h.Properties)
	populate(objectMap, "tags", h.Tags)
	populate(objectMap, "type", h.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HostPoolProperties.
func (h HostPoolProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "agentUpdate", h.AgentUpdate)
	populate(objectMap, "applicationGroupReferences", h.ApplicationGroupReferences)
	populate(objectMap, "cloudPcResource", h.CloudPcResource)
	populate(objectMap, "customRdpProperty", h.CustomRdpProperty)
	populate(objectMap, "description", h.Description)
	populate(objectMap, "friendlyName", h.FriendlyName)
	populate(objectMap, "hostPoolType", h.HostPoolType)
	populate(objectMap, "loadBalancerType", h.LoadBalancerType)
	populate(objectMap, "maxSessionLimit", h.MaxSessionLimit)
	populate(objectMap, "migrationRequest", h.MigrationRequest)
	populate(objectMap, "objectId", h.ObjectID)
	populate(objectMap, "personalDesktopAssignmentType", h.PersonalDesktopAssignmentType)
	populate(objectMap, "preferredAppGroupType", h.PreferredAppGroupType)
	populate(objectMap, "privateEndpointConnections", h.PrivateEndpointConnections)
	populate(objectMap, "publicNetworkAccess", h.PublicNetworkAccess)
	populate(objectMap, "registrationInfo", h.RegistrationInfo)
	populate(objectMap, "ring", h.Ring)
	populate(objectMap, "ssoClientId", h.SsoClientID)
	populate(objectMap, "ssoClientSecretKeyVaultPath", h.SsoClientSecretKeyVaultPath)
	populate(objectMap, "ssoSecretType", h.SsoSecretType)
	populate(objectMap, "ssoadfsAuthority", h.SsoadfsAuthority)
	populate(objectMap, "startVMOnConnect", h.StartVMOnConnect)
	populate(objectMap, "vmTemplate", h.VMTemplate)
	populate(objectMap, "validationEnvironment", h.ValidationEnvironment)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MSIXPackagePatch.
func (m MSIXPackagePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", m.ID)
	populate(objectMap, "name", m.Name)
	populate(objectMap, "properties", m.Properties)
	populate(objectMap, "type", m.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MSIXPackageProperties.
func (m MSIXPackageProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "displayName", m.DisplayName)
	populate(objectMap, "imagePath", m.ImagePath)
	populate(objectMap, "isActive", m.IsActive)
	populate(objectMap, "isRegularRegistration", m.IsRegularRegistration)
	populateTimeRFC3339(objectMap, "lastUpdated", m.LastUpdated)
	populate(objectMap, "packageApplications", m.PackageApplications)
	populate(objectMap, "packageDependencies", m.PackageDependencies)
	populate(objectMap, "packageFamilyName", m.PackageFamilyName)
	populate(objectMap, "packageName", m.PackageName)
	populate(objectMap, "packageRelativePath", m.PackageRelativePath)
	populate(objectMap, "version", m.Version)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MSIXPackageProperties.
func (m *MSIXPackageProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "displayName":
			err = unpopulate(val, "DisplayName", &m.DisplayName)
			delete(rawMsg, key)
		case "imagePath":
			err = unpopulate(val, "ImagePath", &m.ImagePath)
			delete(rawMsg, key)
		case "isActive":
			err = unpopulate(val, "IsActive", &m.IsActive)
			delete(rawMsg, key)
		case "isRegularRegistration":
			err = unpopulate(val, "IsRegularRegistration", &m.IsRegularRegistration)
			delete(rawMsg, key)
		case "lastUpdated":
			err = unpopulateTimeRFC3339(val, "LastUpdated", &m.LastUpdated)
			delete(rawMsg, key)
		case "packageApplications":
			err = unpopulate(val, "PackageApplications", &m.PackageApplications)
			delete(rawMsg, key)
		case "packageDependencies":
			err = unpopulate(val, "PackageDependencies", &m.PackageDependencies)
			delete(rawMsg, key)
		case "packageFamilyName":
			err = unpopulate(val, "PackageFamilyName", &m.PackageFamilyName)
			delete(rawMsg, key)
		case "packageName":
			err = unpopulate(val, "PackageName", &m.PackageName)
			delete(rawMsg, key)
		case "packageRelativePath":
			err = unpopulate(val, "PackageRelativePath", &m.PackageRelativePath)
			delete(rawMsg, key)
		case "version":
			err = unpopulate(val, "Version", &m.Version)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type MsixPackageApplications.
func (m MsixPackageApplications) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "appId", m.AppID)
	populate(objectMap, "appUserModelID", m.AppUserModelID)
	populate(objectMap, "description", m.Description)
	populate(objectMap, "friendlyName", m.FriendlyName)
	populate(objectMap, "iconImageName", m.IconImageName)
	populateByteArray(objectMap, "rawIcon", m.RawIcon, runtime.Base64StdFormat)
	populateByteArray(objectMap, "rawPng", m.RawPNG, runtime.Base64StdFormat)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type MsixPackageApplications.
func (m *MsixPackageApplications) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", m, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "appId":
			err = unpopulate(val, "AppID", &m.AppID)
			delete(rawMsg, key)
		case "appUserModelID":
			err = unpopulate(val, "AppUserModelID", &m.AppUserModelID)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &m.Description)
			delete(rawMsg, key)
		case "friendlyName":
			err = unpopulate(val, "FriendlyName", &m.FriendlyName)
			delete(rawMsg, key)
		case "iconImageName":
			err = unpopulate(val, "IconImageName", &m.IconImageName)
			delete(rawMsg, key)
		case "rawIcon":
			err = runtime.DecodeByteArray(string(val), &m.RawIcon, runtime.Base64StdFormat)
			delete(rawMsg, key)
		case "rawPng":
			err = runtime.DecodeByteArray(string(val), &m.RawPNG, runtime.Base64StdFormat)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", m, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type PrivateLinkResourceProperties.
func (p PrivateLinkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "groupId", p.GroupID)
	populate(objectMap, "requiredMembers", p.RequiredMembers)
	populate(objectMap, "requiredZoneNames", p.RequiredZoneNames)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationInfo.
func (r RegistrationInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "expirationTime", r.ExpirationTime)
	populate(objectMap, "registrationTokenOperation", r.RegistrationTokenOperation)
	populate(objectMap, "token", r.Token)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationInfo.
func (r *RegistrationInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expirationTime":
			err = unpopulateTimeRFC3339(val, "ExpirationTime", &r.ExpirationTime)
			delete(rawMsg, key)
		case "registrationTokenOperation":
			err = unpopulate(val, "RegistrationTokenOperation", &r.RegistrationTokenOperation)
			delete(rawMsg, key)
		case "token":
			err = unpopulate(val, "Token", &r.Token)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type RegistrationInfoPatch.
func (r RegistrationInfoPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "expirationTime", r.ExpirationTime)
	populate(objectMap, "registrationTokenOperation", r.RegistrationTokenOperation)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type RegistrationInfoPatch.
func (r *RegistrationInfoPatch) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", r, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "expirationTime":
			err = unpopulateTimeRFC3339(val, "ExpirationTime", &r.ExpirationTime)
			delete(rawMsg, key)
		case "registrationTokenOperation":
			err = unpopulate(val, "RegistrationTokenOperation", &r.RegistrationTokenOperation)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", r, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ResourceModelWithAllowedPropertySet.
func (r ResourceModelWithAllowedPropertySet) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", r.Etag)
	populate(objectMap, "id", r.ID)
	populate(objectMap, "identity", r.Identity)
	populate(objectMap, "kind", r.Kind)
	populate(objectMap, "location", r.Location)
	populate(objectMap, "managedBy", r.ManagedBy)
	populate(objectMap, "name", r.Name)
	populate(objectMap, "plan", r.Plan)
	populate(objectMap, "sku", r.SKU)
	populate(objectMap, "tags", r.Tags)
	populate(objectMap, "type", r.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ScalingPlan.
func (s ScalingPlan) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", s.Etag)
	populate(objectMap, "id", s.ID)
	populate(objectMap, "identity", s.Identity)
	populate(objectMap, "kind", s.Kind)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "managedBy", s.ManagedBy)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "plan", s.Plan)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "sku", s.SKU)
	populate(objectMap, "systemData", s.SystemData)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ScalingPlanPatch.
func (s ScalingPlanPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ScalingPlanPatchProperties.
func (s ScalingPlanPatchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", s.Description)
	populate(objectMap, "exclusionTag", s.ExclusionTag)
	populate(objectMap, "friendlyName", s.FriendlyName)
	populate(objectMap, "hostPoolReferences", s.HostPoolReferences)
	populate(objectMap, "schedules", s.Schedules)
	populate(objectMap, "timeZone", s.TimeZone)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ScalingPlanProperties.
func (s ScalingPlanProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", s.Description)
	populate(objectMap, "exclusionTag", s.ExclusionTag)
	populate(objectMap, "friendlyName", s.FriendlyName)
	populate(objectMap, "hostPoolReferences", s.HostPoolReferences)
	populate(objectMap, "hostPoolType", s.HostPoolType)
	populate(objectMap, "objectId", s.ObjectID)
	populate(objectMap, "schedules", s.Schedules)
	populate(objectMap, "timeZone", s.TimeZone)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ScalingSchedule.
func (s ScalingSchedule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "daysOfWeek", s.DaysOfWeek)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "offPeakLoadBalancingAlgorithm", s.OffPeakLoadBalancingAlgorithm)
	populate(objectMap, "offPeakStartTime", s.OffPeakStartTime)
	populate(objectMap, "peakLoadBalancingAlgorithm", s.PeakLoadBalancingAlgorithm)
	populate(objectMap, "peakStartTime", s.PeakStartTime)
	populate(objectMap, "rampDownCapacityThresholdPct", s.RampDownCapacityThresholdPct)
	populate(objectMap, "rampDownForceLogoffUsers", s.RampDownForceLogoffUsers)
	populate(objectMap, "rampDownLoadBalancingAlgorithm", s.RampDownLoadBalancingAlgorithm)
	populate(objectMap, "rampDownMinimumHostsPct", s.RampDownMinimumHostsPct)
	populate(objectMap, "rampDownNotificationMessage", s.RampDownNotificationMessage)
	populate(objectMap, "rampDownStartTime", s.RampDownStartTime)
	populate(objectMap, "rampDownStopHostsWhen", s.RampDownStopHostsWhen)
	populate(objectMap, "rampDownWaitTimeMinutes", s.RampDownWaitTimeMinutes)
	populate(objectMap, "rampUpCapacityThresholdPct", s.RampUpCapacityThresholdPct)
	populate(objectMap, "rampUpLoadBalancingAlgorithm", s.RampUpLoadBalancingAlgorithm)
	populate(objectMap, "rampUpMinimumHostsPct", s.RampUpMinimumHostsPct)
	populate(objectMap, "rampUpStartTime", s.RampUpStartTime)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SessionHostHealthCheckFailureDetails.
func (s SessionHostHealthCheckFailureDetails) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "errorCode", s.ErrorCode)
	populateTimeRFC3339(objectMap, "lastHealthCheckDateTime", s.LastHealthCheckDateTime)
	populate(objectMap, "message", s.Message)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SessionHostHealthCheckFailureDetails.
func (s *SessionHostHealthCheckFailureDetails) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "errorCode":
			err = unpopulate(val, "ErrorCode", &s.ErrorCode)
			delete(rawMsg, key)
		case "lastHealthCheckDateTime":
			err = unpopulateTimeRFC3339(val, "LastHealthCheckDateTime", &s.LastHealthCheckDateTime)
			delete(rawMsg, key)
		case "message":
			err = unpopulate(val, "Message", &s.Message)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SessionHostPatch.
func (s SessionHostPatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SessionHostProperties.
func (s SessionHostProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "agentVersion", s.AgentVersion)
	populate(objectMap, "allowNewSession", s.AllowNewSession)
	populate(objectMap, "assignedUser", s.AssignedUser)
	populate(objectMap, "friendlyName", s.FriendlyName)
	populateTimeRFC3339(objectMap, "lastHeartBeat", s.LastHeartBeat)
	populateTimeRFC3339(objectMap, "lastUpdateTime", s.LastUpdateTime)
	populate(objectMap, "osVersion", s.OSVersion)
	populate(objectMap, "objectId", s.ObjectID)
	populate(objectMap, "resourceId", s.ResourceID)
	populate(objectMap, "sessionHostHealthCheckResults", s.SessionHostHealthCheckResults)
	populate(objectMap, "sessions", s.Sessions)
	populate(objectMap, "status", s.Status)
	populateTimeRFC3339(objectMap, "statusTimestamp", s.StatusTimestamp)
	populate(objectMap, "sxSStackVersion", s.SxSStackVersion)
	populate(objectMap, "updateErrorMessage", s.UpdateErrorMessage)
	populate(objectMap, "updateState", s.UpdateState)
	populate(objectMap, "virtualMachineId", s.VirtualMachineID)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SessionHostProperties.
func (s *SessionHostProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "agentVersion":
			err = unpopulate(val, "AgentVersion", &s.AgentVersion)
			delete(rawMsg, key)
		case "allowNewSession":
			err = unpopulate(val, "AllowNewSession", &s.AllowNewSession)
			delete(rawMsg, key)
		case "assignedUser":
			err = unpopulate(val, "AssignedUser", &s.AssignedUser)
			delete(rawMsg, key)
		case "friendlyName":
			err = unpopulate(val, "FriendlyName", &s.FriendlyName)
			delete(rawMsg, key)
		case "lastHeartBeat":
			err = unpopulateTimeRFC3339(val, "LastHeartBeat", &s.LastHeartBeat)
			delete(rawMsg, key)
		case "lastUpdateTime":
			err = unpopulateTimeRFC3339(val, "LastUpdateTime", &s.LastUpdateTime)
			delete(rawMsg, key)
		case "osVersion":
			err = unpopulate(val, "OSVersion", &s.OSVersion)
			delete(rawMsg, key)
		case "objectId":
			err = unpopulate(val, "ObjectID", &s.ObjectID)
			delete(rawMsg, key)
		case "resourceId":
			err = unpopulate(val, "ResourceID", &s.ResourceID)
			delete(rawMsg, key)
		case "sessionHostHealthCheckResults":
			err = unpopulate(val, "SessionHostHealthCheckResults", &s.SessionHostHealthCheckResults)
			delete(rawMsg, key)
		case "sessions":
			err = unpopulate(val, "Sessions", &s.Sessions)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "statusTimestamp":
			err = unpopulateTimeRFC3339(val, "StatusTimestamp", &s.StatusTimestamp)
			delete(rawMsg, key)
		case "sxSStackVersion":
			err = unpopulate(val, "SxSStackVersion", &s.SxSStackVersion)
			delete(rawMsg, key)
		case "updateErrorMessage":
			err = unpopulate(val, "UpdateErrorMessage", &s.UpdateErrorMessage)
			delete(rawMsg, key)
		case "updateState":
			err = unpopulate(val, "UpdateState", &s.UpdateState)
			delete(rawMsg, key)
		case "virtualMachineId":
			err = unpopulate(val, "VirtualMachineID", &s.VirtualMachineID)
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

// MarshalJSON implements the json.Marshaller interface for type UserSessionProperties.
func (u UserSessionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "activeDirectoryUserName", u.ActiveDirectoryUserName)
	populate(objectMap, "applicationType", u.ApplicationType)
	populateTimeRFC3339(objectMap, "createTime", u.CreateTime)
	populate(objectMap, "objectId", u.ObjectID)
	populate(objectMap, "sessionState", u.SessionState)
	populate(objectMap, "userPrincipalName", u.UserPrincipalName)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type UserSessionProperties.
func (u *UserSessionProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", u, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "activeDirectoryUserName":
			err = unpopulate(val, "ActiveDirectoryUserName", &u.ActiveDirectoryUserName)
			delete(rawMsg, key)
		case "applicationType":
			err = unpopulate(val, "ApplicationType", &u.ApplicationType)
			delete(rawMsg, key)
		case "createTime":
			err = unpopulateTimeRFC3339(val, "CreateTime", &u.CreateTime)
			delete(rawMsg, key)
		case "objectId":
			err = unpopulate(val, "ObjectID", &u.ObjectID)
			delete(rawMsg, key)
		case "sessionState":
			err = unpopulate(val, "SessionState", &u.SessionState)
			delete(rawMsg, key)
		case "userPrincipalName":
			err = unpopulate(val, "UserPrincipalName", &u.UserPrincipalName)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", u, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type Workspace.
func (w Workspace) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "etag", w.Etag)
	populate(objectMap, "id", w.ID)
	populate(objectMap, "identity", w.Identity)
	populate(objectMap, "kind", w.Kind)
	populate(objectMap, "location", w.Location)
	populate(objectMap, "managedBy", w.ManagedBy)
	populate(objectMap, "name", w.Name)
	populate(objectMap, "plan", w.Plan)
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "sku", w.SKU)
	populate(objectMap, "systemData", w.SystemData)
	populate(objectMap, "tags", w.Tags)
	populate(objectMap, "type", w.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkspacePatch.
func (w WorkspacePatch) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "properties", w.Properties)
	populate(objectMap, "tags", w.Tags)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkspacePatchProperties.
func (w WorkspacePatchProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicationGroupReferences", w.ApplicationGroupReferences)
	populate(objectMap, "description", w.Description)
	populate(objectMap, "friendlyName", w.FriendlyName)
	populate(objectMap, "publicNetworkAccess", w.PublicNetworkAccess)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type WorkspaceProperties.
func (w WorkspaceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "applicationGroupReferences", w.ApplicationGroupReferences)
	populate(objectMap, "cloudPcResource", w.CloudPcResource)
	populate(objectMap, "description", w.Description)
	populate(objectMap, "friendlyName", w.FriendlyName)
	populate(objectMap, "objectId", w.ObjectID)
	populate(objectMap, "privateEndpointConnections", w.PrivateEndpointConnections)
	populate(objectMap, "publicNetworkAccess", w.PublicNetworkAccess)
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
