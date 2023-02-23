//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

import (
	"encoding/json"
	"fmt"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type CanaryTrafficRegionRolloutConfiguration.
func (c CanaryTrafficRegionRolloutConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", c.Regions)
	populate(objectMap, "skipRegions", c.SkipRegions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CheckNameAvailabilitySpecifications.
func (c CheckNameAvailabilitySpecifications) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "enableDefaultValidation", c.EnableDefaultValidation)
	populate(objectMap, "resourceTypesWithCustomValidation", c.ResourceTypesWithCustomValidation)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CustomRolloutPropertiesSpecification.
func (c CustomRolloutPropertiesSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "canary", c.Canary)
	populate(objectMap, "providerRegistration", c.ProviderRegistration)
	populate(objectMap, "resourceTypeRegistrations", c.ResourceTypeRegistrations)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CustomRolloutPropertiesStatus.
func (c CustomRolloutPropertiesStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "completedRegions", c.CompletedRegions)
	populate(objectMap, "failedOrSkippedRegions", c.FailedOrSkippedRegions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CustomRolloutSpecification.
func (c CustomRolloutSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "canary", c.Canary)
	populate(objectMap, "providerRegistration", c.ProviderRegistration)
	populate(objectMap, "resourceTypeRegistrations", c.ResourceTypeRegistrations)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CustomRolloutSpecificationCanary.
func (c CustomRolloutSpecificationCanary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", c.Regions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type CustomRolloutStatus.
func (c CustomRolloutStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "completedRegions", c.CompletedRegions)
	populate(objectMap, "failedOrSkippedRegions", c.FailedOrSkippedRegions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutPropertiesSpecification.
func (d DefaultRolloutPropertiesSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "canary", d.Canary)
	populate(objectMap, "highTraffic", d.HighTraffic)
	populate(objectMap, "lowTraffic", d.LowTraffic)
	populate(objectMap, "mediumTraffic", d.MediumTraffic)
	populate(objectMap, "providerRegistration", d.ProviderRegistration)
	populate(objectMap, "resourceTypeRegistrations", d.ResourceTypeRegistrations)
	populate(objectMap, "restOfTheWorldGroupOne", d.RestOfTheWorldGroupOne)
	populate(objectMap, "restOfTheWorldGroupTwo", d.RestOfTheWorldGroupTwo)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutPropertiesStatus.
func (d DefaultRolloutPropertiesStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "completedRegions", d.CompletedRegions)
	populate(objectMap, "failedOrSkippedRegions", d.FailedOrSkippedRegions)
	populate(objectMap, "nextTrafficRegion", d.NextTrafficRegion)
	populateTimeRFC3339(objectMap, "nextTrafficRegionScheduledTime", d.NextTrafficRegionScheduledTime)
	populate(objectMap, "subscriptionReregistrationResult", d.SubscriptionReregistrationResult)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DefaultRolloutPropertiesStatus.
func (d *DefaultRolloutPropertiesStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "completedRegions":
			err = unpopulate(val, "CompletedRegions", &d.CompletedRegions)
			delete(rawMsg, key)
		case "failedOrSkippedRegions":
			err = unpopulate(val, "FailedOrSkippedRegions", &d.FailedOrSkippedRegions)
			delete(rawMsg, key)
		case "nextTrafficRegion":
			err = unpopulate(val, "NextTrafficRegion", &d.NextTrafficRegion)
			delete(rawMsg, key)
		case "nextTrafficRegionScheduledTime":
			err = unpopulateTimeRFC3339(val, "NextTrafficRegionScheduledTime", &d.NextTrafficRegionScheduledTime)
			delete(rawMsg, key)
		case "subscriptionReregistrationResult":
			err = unpopulate(val, "SubscriptionReregistrationResult", &d.SubscriptionReregistrationResult)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutSpecification.
func (d DefaultRolloutSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "canary", d.Canary)
	populate(objectMap, "highTraffic", d.HighTraffic)
	populate(objectMap, "lowTraffic", d.LowTraffic)
	populate(objectMap, "mediumTraffic", d.MediumTraffic)
	populate(objectMap, "providerRegistration", d.ProviderRegistration)
	populate(objectMap, "resourceTypeRegistrations", d.ResourceTypeRegistrations)
	populate(objectMap, "restOfTheWorldGroupOne", d.RestOfTheWorldGroupOne)
	populate(objectMap, "restOfTheWorldGroupTwo", d.RestOfTheWorldGroupTwo)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutSpecificationCanary.
func (d DefaultRolloutSpecificationCanary) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", d.Regions)
	populate(objectMap, "skipRegions", d.SkipRegions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutSpecificationHighTraffic.
func (d DefaultRolloutSpecificationHighTraffic) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", d.Regions)
	populate(objectMap, "waitDuration", d.WaitDuration)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutSpecificationLowTraffic.
func (d DefaultRolloutSpecificationLowTraffic) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", d.Regions)
	populate(objectMap, "waitDuration", d.WaitDuration)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutSpecificationMediumTraffic.
func (d DefaultRolloutSpecificationMediumTraffic) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", d.Regions)
	populate(objectMap, "waitDuration", d.WaitDuration)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutSpecificationRestOfTheWorldGroupOne.
func (d DefaultRolloutSpecificationRestOfTheWorldGroupOne) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", d.Regions)
	populate(objectMap, "waitDuration", d.WaitDuration)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutSpecificationRestOfTheWorldGroupTwo.
func (d DefaultRolloutSpecificationRestOfTheWorldGroupTwo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", d.Regions)
	populate(objectMap, "waitDuration", d.WaitDuration)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type DefaultRolloutStatus.
func (d DefaultRolloutStatus) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "completedRegions", d.CompletedRegions)
	populate(objectMap, "failedOrSkippedRegions", d.FailedOrSkippedRegions)
	populate(objectMap, "nextTrafficRegion", d.NextTrafficRegion)
	populateTimeRFC3339(objectMap, "nextTrafficRegionScheduledTime", d.NextTrafficRegionScheduledTime)
	populate(objectMap, "subscriptionReregistrationResult", d.SubscriptionReregistrationResult)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DefaultRolloutStatus.
func (d *DefaultRolloutStatus) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "completedRegions":
			err = unpopulate(val, "CompletedRegions", &d.CompletedRegions)
			delete(rawMsg, key)
		case "failedOrSkippedRegions":
			err = unpopulate(val, "FailedOrSkippedRegions", &d.FailedOrSkippedRegions)
			delete(rawMsg, key)
		case "nextTrafficRegion":
			err = unpopulate(val, "NextTrafficRegion", &d.NextTrafficRegion)
			delete(rawMsg, key)
		case "nextTrafficRegionScheduledTime":
			err = unpopulateTimeRFC3339(val, "NextTrafficRegionScheduledTime", &d.NextTrafficRegionScheduledTime)
			delete(rawMsg, key)
		case "subscriptionReregistrationResult":
			err = unpopulate(val, "SubscriptionReregistrationResult", &d.SubscriptionReregistrationResult)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ErrorInnerError.
func (e *ErrorInnerError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", e, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &e.Code)
			delete(rawMsg, key)
		case "innerError":
			err = unpopulate(val, "InnerError", &e.InnerError)
			delete(rawMsg, key)
		default:
			if e.AdditionalProperties == nil {
				e.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				e.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", e, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ExtendedErrorInfo.
func (e ExtendedErrorInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ExtensionOptions.
func (e ExtensionOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "request", e.Request)
	populate(objectMap, "response", e.Response)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InnerError.
func (i *InnerError) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "code":
			err = unpopulate(val, "Code", &i.Code)
			delete(rawMsg, key)
		case "innerError":
			err = unpopulate(val, "InnerError", &i.InnerError)
			delete(rawMsg, key)
		default:
			if i.AdditionalProperties == nil {
				i.AdditionalProperties = map[string]interface{}{}
			}
			if val != nil {
				var aux interface{}
				err = json.Unmarshal(val, &aux)
				i.AdditionalProperties[key] = aux
			}
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LoggingHiddenPropertyPath.
func (l LoggingHiddenPropertyPath) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hiddenPathsOnRequest", l.HiddenPathsOnRequest)
	populate(objectMap, "hiddenPathsOnResponse", l.HiddenPathsOnResponse)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type LoggingRuleHiddenPropertyPaths.
func (l LoggingRuleHiddenPropertyPaths) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hiddenPathsOnRequest", l.HiddenPathsOnRequest)
	populate(objectMap, "hiddenPathsOnResponse", l.HiddenPathsOnResponse)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type Metadata.
func (m Metadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "providerAuthentication", m.ProviderAuthentication)
	populate(objectMap, "providerAuthorizations", m.ProviderAuthorizations)
	populate(objectMap, "thirdPartyProviderAuthorization", m.ThirdPartyProviderAuthorization)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MetadataProviderAuthentication.
func (m MetadataProviderAuthentication) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedAudiences", m.AllowedAudiences)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type MetadataThirdPartyProviderAuthorization.
func (m MetadataThirdPartyProviderAuthorization) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizations", m.Authorizations)
	populate(objectMap, "managedByTenantId", m.ManagedByTenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NotificationEndpoint.
func (n NotificationEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "locations", n.Locations)
	populate(objectMap, "notificationDestination", n.NotificationDestination)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NotificationRegistrationProperties.
func (n NotificationRegistrationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "includedEvents", n.IncludedEvents)
	populate(objectMap, "messageScope", n.MessageScope)
	populate(objectMap, "notificationEndpoints", n.NotificationEndpoints)
	populate(objectMap, "notificationMode", n.NotificationMode)
	populate(objectMap, "provisioningState", n.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NotificationRegistrationPropertiesAutoGenerated.
func (n NotificationRegistrationPropertiesAutoGenerated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "includedEvents", n.IncludedEvents)
	populate(objectMap, "messageScope", n.MessageScope)
	populate(objectMap, "notificationEndpoints", n.NotificationEndpoints)
	populate(objectMap, "notificationMode", n.NotificationMode)
	populate(objectMap, "provisioningState", n.ProvisioningState)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type OperationsPutContent.
func (o OperationsPutContent) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contents", o.Contents)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProviderRegistrationProperties.
func (p ProviderRegistrationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capabilities", p.Capabilities)
	populate(objectMap, "featuresRule", p.FeaturesRule)
	populate(objectMap, "management", p.Management)
	populate(objectMap, "metadata", &p.Metadata)
	populate(objectMap, "namespace", p.Namespace)
	populate(objectMap, "providerAuthentication", p.ProviderAuthentication)
	populate(objectMap, "providerAuthorizations", p.ProviderAuthorizations)
	populate(objectMap, "providerHubMetadata", p.ProviderHubMetadata)
	populate(objectMap, "providerType", p.ProviderType)
	populate(objectMap, "providerVersion", p.ProviderVersion)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "requestHeaderOptions", p.RequestHeaderOptions)
	populate(objectMap, "requiredFeatures", p.RequiredFeatures)
	populate(objectMap, "subscriptionLifecycleNotificationSpecifications", p.SubscriptionLifecycleNotificationSpecifications)
	populate(objectMap, "templateDeploymentOptions", p.TemplateDeploymentOptions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProviderRegistrationPropertiesAutoGenerated.
func (p ProviderRegistrationPropertiesAutoGenerated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capabilities", p.Capabilities)
	populate(objectMap, "featuresRule", p.FeaturesRule)
	populate(objectMap, "management", p.Management)
	populate(objectMap, "metadata", &p.Metadata)
	populate(objectMap, "namespace", p.Namespace)
	populate(objectMap, "providerAuthentication", p.ProviderAuthentication)
	populate(objectMap, "providerAuthorizations", p.ProviderAuthorizations)
	populate(objectMap, "providerHubMetadata", p.ProviderHubMetadata)
	populate(objectMap, "providerType", p.ProviderType)
	populate(objectMap, "providerVersion", p.ProviderVersion)
	populate(objectMap, "provisioningState", p.ProvisioningState)
	populate(objectMap, "requestHeaderOptions", p.RequestHeaderOptions)
	populate(objectMap, "requiredFeatures", p.RequiredFeatures)
	populate(objectMap, "subscriptionLifecycleNotificationSpecifications", p.SubscriptionLifecycleNotificationSpecifications)
	populate(objectMap, "templateDeploymentOptions", p.TemplateDeploymentOptions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProviderRegistrationPropertiesProviderHubMetadata.
func (p ProviderRegistrationPropertiesProviderHubMetadata) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "providerAuthentication", p.ProviderAuthentication)
	populate(objectMap, "providerAuthorizations", p.ProviderAuthorizations)
	populate(objectMap, "thirdPartyProviderAuthorization", p.ThirdPartyProviderAuthorization)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ProviderRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications.
func (p ProviderRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "softDeleteTTL", p.SoftDeleteTTL)
	populate(objectMap, "subscriptionStateOverrideActions", p.SubscriptionStateOverrideActions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderAuthentication.
func (r ResourceProviderAuthentication) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedAudiences", r.AllowedAudiences)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderCapabilities.
func (r ResourceProviderCapabilities) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "effect", r.Effect)
	populate(objectMap, "quotaId", r.QuotaID)
	populate(objectMap, "requiredFeatures", r.RequiredFeatures)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderManagement.
func (r ResourceProviderManagement) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "incidentContactEmail", r.IncidentContactEmail)
	populate(objectMap, "incidentRoutingService", r.IncidentRoutingService)
	populate(objectMap, "incidentRoutingTeam", r.IncidentRoutingTeam)
	populate(objectMap, "manifestOwners", r.ManifestOwners)
	populate(objectMap, "resourceAccessPolicy", r.ResourceAccessPolicy)
	populate(objectMap, "resourceAccessRoles", r.ResourceAccessRoles)
	populate(objectMap, "schemaOwners", r.SchemaOwners)
	populate(objectMap, "serviceTreeInfos", r.ServiceTreeInfos)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderManifestManagement.
func (r ResourceProviderManifestManagement) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "incidentContactEmail", r.IncidentContactEmail)
	populate(objectMap, "incidentRoutingService", r.IncidentRoutingService)
	populate(objectMap, "incidentRoutingTeam", r.IncidentRoutingTeam)
	populate(objectMap, "manifestOwners", r.ManifestOwners)
	populate(objectMap, "resourceAccessPolicy", r.ResourceAccessPolicy)
	populate(objectMap, "resourceAccessRoles", r.ResourceAccessRoles)
	populate(objectMap, "schemaOwners", r.SchemaOwners)
	populate(objectMap, "serviceTreeInfos", r.ServiceTreeInfos)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderManifestProperties.
func (r ResourceProviderManifestProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capabilities", r.Capabilities)
	populate(objectMap, "featuresRule", r.FeaturesRule)
	populate(objectMap, "management", r.Management)
	populate(objectMap, "metadata", &r.Metadata)
	populate(objectMap, "namespace", r.Namespace)
	populate(objectMap, "providerAuthentication", r.ProviderAuthentication)
	populate(objectMap, "providerAuthorizations", r.ProviderAuthorizations)
	populate(objectMap, "providerType", r.ProviderType)
	populate(objectMap, "providerVersion", r.ProviderVersion)
	populate(objectMap, "requestHeaderOptions", r.RequestHeaderOptions)
	populate(objectMap, "requiredFeatures", r.RequiredFeatures)
	populate(objectMap, "templateDeploymentOptions", r.TemplateDeploymentOptions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderManifestPropertiesManagement.
func (r ResourceProviderManifestPropertiesManagement) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "incidentContactEmail", r.IncidentContactEmail)
	populate(objectMap, "incidentRoutingService", r.IncidentRoutingService)
	populate(objectMap, "incidentRoutingTeam", r.IncidentRoutingTeam)
	populate(objectMap, "manifestOwners", r.ManifestOwners)
	populate(objectMap, "resourceAccessPolicy", r.ResourceAccessPolicy)
	populate(objectMap, "resourceAccessRoles", r.ResourceAccessRoles)
	populate(objectMap, "schemaOwners", r.SchemaOwners)
	populate(objectMap, "serviceTreeInfos", r.ServiceTreeInfos)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderManifestPropertiesProviderAuthentication.
func (r ResourceProviderManifestPropertiesProviderAuthentication) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedAudiences", r.AllowedAudiences)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderManifestPropertiesTemplateDeploymentOptions.
func (r ResourceProviderManifestPropertiesTemplateDeploymentOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "preflightOptions", r.PreflightOptions)
	populate(objectMap, "preflightSupported", r.PreflightSupported)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceProviderManifestProviderAuthentication.
func (r ResourceProviderManifestProviderAuthentication) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedAudiences", r.AllowedAudiences)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeEndpoint.
func (r ResourceTypeEndpoint) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "apiVersions", r.APIVersions)
	populate(objectMap, "enabled", r.Enabled)
	populate(objectMap, "extensions", r.Extensions)
	populate(objectMap, "featuresRule", r.FeaturesRule)
	populate(objectMap, "locations", r.Locations)
	populate(objectMap, "requiredFeatures", r.RequiredFeatures)
	populate(objectMap, "timeout", r.Timeout)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeExtension.
func (r ResourceTypeExtension) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endpointUri", r.EndpointURI)
	populate(objectMap, "extensionCategories", r.ExtensionCategories)
	populate(objectMap, "timeout", r.Timeout)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeExtensionOptionsResourceCreationBegin.
func (r ResourceTypeExtensionOptionsResourceCreationBegin) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "request", r.Request)
	populate(objectMap, "response", r.Response)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeRegistrationProperties.
func (r ResourceTypeRegistrationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedUnauthorizedActions", r.AllowedUnauthorizedActions)
	populate(objectMap, "authorizationActionMappings", r.AuthorizationActionMappings)
	populate(objectMap, "checkNameAvailabilitySpecifications", r.CheckNameAvailabilitySpecifications)
	populate(objectMap, "defaultApiVersion", r.DefaultAPIVersion)
	populate(objectMap, "disallowedActionVerbs", r.DisallowedActionVerbs)
	populate(objectMap, "enableAsyncOperation", r.EnableAsyncOperation)
	populate(objectMap, "enableThirdPartyS2S", r.EnableThirdPartyS2S)
	populate(objectMap, "endpoints", r.Endpoints)
	populate(objectMap, "extendedLocations", r.ExtendedLocations)
	populate(objectMap, "extensionOptions", r.ExtensionOptions)
	populate(objectMap, "featuresRule", r.FeaturesRule)
	populate(objectMap, "identityManagement", r.IdentityManagement)
	populate(objectMap, "isPureProxy", r.IsPureProxy)
	populate(objectMap, "linkedAccessChecks", r.LinkedAccessChecks)
	populate(objectMap, "loggingRules", r.LoggingRules)
	populate(objectMap, "marketplaceType", r.MarketplaceType)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "regionality", r.Regionality)
	populate(objectMap, "requestHeaderOptions", r.RequestHeaderOptions)
	populate(objectMap, "requiredFeatures", r.RequiredFeatures)
	populate(objectMap, "resourceDeletionPolicy", r.ResourceDeletionPolicy)
	populate(objectMap, "resourceMovePolicy", r.ResourceMovePolicy)
	populate(objectMap, "routingType", r.RoutingType)
	populate(objectMap, "serviceTreeInfos", r.ServiceTreeInfos)
	populate(objectMap, "subscriptionLifecycleNotificationSpecifications", r.SubscriptionLifecycleNotificationSpecifications)
	populate(objectMap, "subscriptionStateRules", r.SubscriptionStateRules)
	populate(objectMap, "swaggerSpecifications", r.SwaggerSpecifications)
	populate(objectMap, "templateDeploymentOptions", r.TemplateDeploymentOptions)
	populate(objectMap, "throttlingRules", r.ThrottlingRules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeRegistrationPropertiesAutoGenerated.
func (r ResourceTypeRegistrationPropertiesAutoGenerated) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedUnauthorizedActions", r.AllowedUnauthorizedActions)
	populate(objectMap, "authorizationActionMappings", r.AuthorizationActionMappings)
	populate(objectMap, "checkNameAvailabilitySpecifications", r.CheckNameAvailabilitySpecifications)
	populate(objectMap, "defaultApiVersion", r.DefaultAPIVersion)
	populate(objectMap, "disallowedActionVerbs", r.DisallowedActionVerbs)
	populate(objectMap, "enableAsyncOperation", r.EnableAsyncOperation)
	populate(objectMap, "enableThirdPartyS2S", r.EnableThirdPartyS2S)
	populate(objectMap, "endpoints", r.Endpoints)
	populate(objectMap, "extendedLocations", r.ExtendedLocations)
	populate(objectMap, "extensionOptions", r.ExtensionOptions)
	populate(objectMap, "featuresRule", r.FeaturesRule)
	populate(objectMap, "identityManagement", r.IdentityManagement)
	populate(objectMap, "isPureProxy", r.IsPureProxy)
	populate(objectMap, "linkedAccessChecks", r.LinkedAccessChecks)
	populate(objectMap, "loggingRules", r.LoggingRules)
	populate(objectMap, "marketplaceType", r.MarketplaceType)
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "regionality", r.Regionality)
	populate(objectMap, "requestHeaderOptions", r.RequestHeaderOptions)
	populate(objectMap, "requiredFeatures", r.RequiredFeatures)
	populate(objectMap, "resourceDeletionPolicy", r.ResourceDeletionPolicy)
	populate(objectMap, "resourceMovePolicy", r.ResourceMovePolicy)
	populate(objectMap, "routingType", r.RoutingType)
	populate(objectMap, "serviceTreeInfos", r.ServiceTreeInfos)
	populate(objectMap, "subscriptionLifecycleNotificationSpecifications", r.SubscriptionLifecycleNotificationSpecifications)
	populate(objectMap, "subscriptionStateRules", r.SubscriptionStateRules)
	populate(objectMap, "swaggerSpecifications", r.SwaggerSpecifications)
	populate(objectMap, "templateDeploymentOptions", r.TemplateDeploymentOptions)
	populate(objectMap, "throttlingRules", r.ThrottlingRules)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecifications.
func (r ResourceTypeRegistrationPropertiesCheckNameAvailabilitySpecifications) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "enableDefaultValidation", r.EnableDefaultValidation)
	populate(objectMap, "resourceTypesWithCustomValidation", r.ResourceTypesWithCustomValidation)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications.
func (r ResourceTypeRegistrationPropertiesSubscriptionLifecycleNotificationSpecifications) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "softDeleteTTL", r.SoftDeleteTTL)
	populate(objectMap, "subscriptionStateOverrideActions", r.SubscriptionStateOverrideActions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeRegistrationPropertiesTemplateDeploymentOptions.
func (r ResourceTypeRegistrationPropertiesTemplateDeploymentOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "preflightOptions", r.PreflightOptions)
	populate(objectMap, "preflightSupported", r.PreflightSupported)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ResourceTypeSKU.
func (r ResourceTypeSKU) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "provisioningState", r.ProvisioningState)
	populate(objectMap, "skuSettings", r.SKUSettings)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type RolloutStatusBase.
func (r RolloutStatusBase) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "completedRegions", r.CompletedRegions)
	populate(objectMap, "failedOrSkippedRegions", r.FailedOrSkippedRegions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKULocationInfo.
func (s SKULocationInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "extendedLocations", s.ExtendedLocations)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "type", s.Type)
	populate(objectMap, "zoneDetails", s.ZoneDetails)
	populate(objectMap, "zones", s.Zones)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKUResourceProperties.
func (s SKUResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "skuSettings", s.SKUSettings)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKUSetting.
func (s SKUSetting) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capabilities", s.Capabilities)
	populate(objectMap, "capacity", s.Capacity)
	populate(objectMap, "costs", s.Costs)
	populate(objectMap, "family", s.Family)
	populate(objectMap, "kind", s.Kind)
	populate(objectMap, "locationInfo", s.LocationInfo)
	populate(objectMap, "locations", s.Locations)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "requiredFeatures", s.RequiredFeatures)
	populate(objectMap, "requiredQuotaIds", s.RequiredQuotaIDs)
	populate(objectMap, "size", s.Size)
	populate(objectMap, "tier", s.Tier)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SKUZoneDetail.
func (s SKUZoneDetail) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "capabilities", s.Capabilities)
	populate(objectMap, "name", s.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubscriptionLifecycleNotificationSpecifications.
func (s SubscriptionLifecycleNotificationSpecifications) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "softDeleteTTL", s.SoftDeleteTTL)
	populate(objectMap, "subscriptionStateOverrideActions", s.SubscriptionStateOverrideActions)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SubscriptionStateRule.
func (s SubscriptionStateRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "allowedActions", s.AllowedActions)
	populate(objectMap, "state", s.State)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type SwaggerSpecification.
func (s SwaggerSpecification) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "apiVersions", s.APIVersions)
	populate(objectMap, "swaggerSpecFolderUri", s.SwaggerSpecFolderURI)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TemplateDeploymentOptions.
func (t TemplateDeploymentOptions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "preflightOptions", t.PreflightOptions)
	populate(objectMap, "preflightSupported", t.PreflightSupported)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ThirdPartyProviderAuthorization.
func (t ThirdPartyProviderAuthorization) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "authorizations", t.Authorizations)
	populate(objectMap, "managedByTenantId", t.ManagedByTenantID)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ThrottlingRule.
func (t ThrottlingRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "action", t.Action)
	populate(objectMap, "metrics", t.Metrics)
	populate(objectMap, "requiredFeatures", t.RequiredFeatures)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TrafficRegionRolloutConfiguration.
func (t TrafficRegionRolloutConfiguration) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", t.Regions)
	populate(objectMap, "waitDuration", t.WaitDuration)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type TrafficRegions.
func (t TrafficRegions) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "regions", t.Regions)
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
