//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmesh

import (
	"encoding/json"
	"fmt"
	"github.com/c-brooks/azure-sdk-for-go/sdk/azcore"
	"reflect"
)

// MarshalJSON implements the json.Marshaller interface for type AddRemoveReplicaScalingMechanism.
func (a AddRemoveReplicaScalingMechanism) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["kind"] = AutoScalingMechanismKindAddRemoveReplica
	populate(objectMap, "maxCount", a.MaxCount)
	populate(objectMap, "minCount", a.MinCount)
	populate(objectMap, "scaleIncrement", a.ScaleIncrement)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AddRemoveReplicaScalingMechanism.
func (a *AddRemoveReplicaScalingMechanism) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "kind":
			err = unpopulate(val, "Kind", &a.Kind)
			delete(rawMsg, key)
		case "maxCount":
			err = unpopulate(val, "MaxCount", &a.MaxCount)
			delete(rawMsg, key)
		case "minCount":
			err = unpopulate(val, "MinCount", &a.MinCount)
			delete(rawMsg, key)
		case "scaleIncrement":
			err = unpopulate(val, "ScaleIncrement", &a.ScaleIncrement)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationProperties.
func (a ApplicationProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "debugParams", a.DebugParams)
	populate(objectMap, "description", a.Description)
	populate(objectMap, "diagnostics", a.Diagnostics)
	populate(objectMap, "healthState", a.HealthState)
	populate(objectMap, "serviceNames", a.ServiceNames)
	populate(objectMap, "services", a.Services)
	populate(objectMap, "status", a.Status)
	populate(objectMap, "statusDetails", a.StatusDetails)
	populate(objectMap, "unhealthyEvaluation", a.UnhealthyEvaluation)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationResourceDescription.
func (a ApplicationResourceDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", a.ID)
	populate(objectMap, "location", a.Location)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "properties", a.Properties)
	populate(objectMap, "tags", a.Tags)
	populate(objectMap, "type", a.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationResourceProperties.
func (a ApplicationResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "debugParams", a.DebugParams)
	populate(objectMap, "description", a.Description)
	populate(objectMap, "diagnostics", a.Diagnostics)
	populate(objectMap, "healthState", a.HealthState)
	populate(objectMap, "provisioningState", a.ProvisioningState)
	populate(objectMap, "serviceNames", a.ServiceNames)
	populate(objectMap, "services", a.Services)
	populate(objectMap, "status", a.Status)
	populate(objectMap, "statusDetails", a.StatusDetails)
	populate(objectMap, "unhealthyEvaluation", a.UnhealthyEvaluation)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationScopedVolume.
func (a ApplicationScopedVolume) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "creationParameters", a.CreationParameters)
	populate(objectMap, "destinationPath", a.DestinationPath)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "readOnly", a.ReadOnly)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ApplicationScopedVolume.
func (a *ApplicationScopedVolume) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "creationParameters":
			a.CreationParameters, err = unmarshalApplicationScopedVolumeCreationParametersClassification(val)
			delete(rawMsg, key)
		case "destinationPath":
			err = unpopulate(val, "DestinationPath", &a.DestinationPath)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "readOnly":
			err = unpopulate(val, "ReadOnly", &a.ReadOnly)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk.
func (a ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", a.Description)
	objectMap["kind"] = ApplicationScopedVolumeKindServiceFabricVolumeDisk
	populate(objectMap, "sizeDisk", a.SizeDisk)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk.
func (a *ApplicationScopedVolumeCreationParametersServiceFabricVolumeDisk) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &a.Description)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &a.Kind)
			delete(rawMsg, key)
		case "sizeDisk":
			err = unpopulate(val, "SizeDisk", &a.SizeDisk)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AutoScalingPolicy.
func (a AutoScalingPolicy) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "mechanism", a.Mechanism)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "trigger", a.Trigger)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AutoScalingPolicy.
func (a *AutoScalingPolicy) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "mechanism":
			a.Mechanism, err = unmarshalAutoScalingMechanismClassification(val)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "trigger":
			a.Trigger, err = unmarshalAutoScalingTriggerClassification(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AutoScalingResourceMetric.
func (a AutoScalingResourceMetric) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["kind"] = AutoScalingMetricKindResource
	populate(objectMap, "name", a.Name)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AutoScalingResourceMetric.
func (a *AutoScalingResourceMetric) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "kind":
			err = unpopulate(val, "Kind", &a.Kind)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AverageLoadScalingTrigger.
func (a AverageLoadScalingTrigger) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	objectMap["kind"] = AutoScalingTriggerKindAverageLoad
	populate(objectMap, "lowerLoadThreshold", a.LowerLoadThreshold)
	populate(objectMap, "metric", a.Metric)
	populate(objectMap, "scaleIntervalInSeconds", a.ScaleIntervalInSeconds)
	populate(objectMap, "upperLoadThreshold", a.UpperLoadThreshold)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AverageLoadScalingTrigger.
func (a *AverageLoadScalingTrigger) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "kind":
			err = unpopulate(val, "Kind", &a.Kind)
			delete(rawMsg, key)
		case "lowerLoadThreshold":
			err = unpopulate(val, "LowerLoadThreshold", &a.LowerLoadThreshold)
			delete(rawMsg, key)
		case "metric":
			a.Metric, err = unmarshalAutoScalingMetricClassification(val)
			delete(rawMsg, key)
		case "scaleIntervalInSeconds":
			err = unpopulate(val, "ScaleIntervalInSeconds", &a.ScaleIntervalInSeconds)
			delete(rawMsg, key)
		case "upperLoadThreshold":
			err = unpopulate(val, "UpperLoadThreshold", &a.UpperLoadThreshold)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type AzureInternalMonitoringPipelineSinkDescription.
func (a AzureInternalMonitoringPipelineSinkDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "accountName", a.AccountName)
	populate(objectMap, "autoKeyConfigUrl", a.AutoKeyConfigURL)
	populate(objectMap, "description", a.Description)
	populate(objectMap, "fluentdConfigUrl", &a.FluentdConfigURL)
	objectMap["kind"] = DiagnosticsSinkKindAzureInternalMonitoringPipeline
	populate(objectMap, "maConfigUrl", a.MaConfigURL)
	populate(objectMap, "name", a.Name)
	populate(objectMap, "namespace", a.Namespace)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type AzureInternalMonitoringPipelineSinkDescription.
func (a *AzureInternalMonitoringPipelineSinkDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", a, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "accountName":
			err = unpopulate(val, "AccountName", &a.AccountName)
			delete(rawMsg, key)
		case "autoKeyConfigUrl":
			err = unpopulate(val, "AutoKeyConfigURL", &a.AutoKeyConfigURL)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &a.Description)
			delete(rawMsg, key)
		case "fluentdConfigUrl":
			err = unpopulate(val, "FluentdConfigURL", &a.FluentdConfigURL)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &a.Kind)
			delete(rawMsg, key)
		case "maConfigUrl":
			err = unpopulate(val, "MaConfigURL", &a.MaConfigURL)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &a.Name)
			delete(rawMsg, key)
		case "namespace":
			err = unpopulate(val, "Namespace", &a.Namespace)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", a, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type ContainerCodePackageProperties.
func (c ContainerCodePackageProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "commands", c.Commands)
	populate(objectMap, "diagnostics", c.Diagnostics)
	populate(objectMap, "endpoints", c.Endpoints)
	populate(objectMap, "entrypoint", c.Entrypoint)
	populate(objectMap, "environmentVariables", c.EnvironmentVariables)
	populate(objectMap, "image", c.Image)
	populate(objectMap, "imageRegistryCredential", c.ImageRegistryCredential)
	populate(objectMap, "instanceView", c.InstanceView)
	populate(objectMap, "labels", c.Labels)
	populate(objectMap, "name", c.Name)
	populate(objectMap, "reliableCollectionsRefs", c.ReliableCollectionsRefs)
	populate(objectMap, "resources", c.Resources)
	populate(objectMap, "settings", c.Settings)
	populate(objectMap, "volumeRefs", c.VolumeRefs)
	populate(objectMap, "volumes", c.Volumes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContainerInstanceView.
func (c ContainerInstanceView) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "currentState", c.CurrentState)
	populate(objectMap, "events", c.Events)
	populate(objectMap, "previousState", c.PreviousState)
	populate(objectMap, "restartCount", c.RestartCount)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ContainerState.
func (c ContainerState) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "detailStatus", c.DetailStatus)
	populate(objectMap, "exitCode", c.ExitCode)
	populateTimeRFC3339(objectMap, "finishTime", c.FinishTime)
	populateTimeRFC3339(objectMap, "startTime", c.StartTime)
	populate(objectMap, "state", c.State)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ContainerState.
func (c *ContainerState) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", c, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "detailStatus":
			err = unpopulate(val, "DetailStatus", &c.DetailStatus)
			delete(rawMsg, key)
		case "exitCode":
			err = unpopulate(val, "ExitCode", &c.ExitCode)
			delete(rawMsg, key)
		case "finishTime":
			err = unpopulateTimeRFC3339(val, "FinishTime", &c.FinishTime)
			delete(rawMsg, key)
		case "startTime":
			err = unpopulateTimeRFC3339(val, "StartTime", &c.StartTime)
			delete(rawMsg, key)
		case "state":
			err = unpopulate(val, "State", &c.State)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", c, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DiagnosticsDescription.
func (d DiagnosticsDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "defaultSinkRefs", d.DefaultSinkRefs)
	populate(objectMap, "enabled", d.Enabled)
	populate(objectMap, "sinks", d.Sinks)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type DiagnosticsDescription.
func (d *DiagnosticsDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", d, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "defaultSinkRefs":
			err = unpopulate(val, "DefaultSinkRefs", &d.DefaultSinkRefs)
			delete(rawMsg, key)
		case "enabled":
			err = unpopulate(val, "Enabled", &d.Enabled)
			delete(rawMsg, key)
		case "sinks":
			d.Sinks, err = unmarshalDiagnosticsSinkPropertiesClassificationArray(val)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", d, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type DiagnosticsRef.
func (d DiagnosticsRef) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "enabled", d.Enabled)
	populate(objectMap, "sinkRefs", d.SinkRefs)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GatewayProperties.
func (g GatewayProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", g.Description)
	populate(objectMap, "destinationNetwork", g.DestinationNetwork)
	populate(objectMap, "http", g.HTTP)
	populate(objectMap, "ipAddress", g.IPAddress)
	populate(objectMap, "sourceNetwork", g.SourceNetwork)
	populate(objectMap, "status", g.Status)
	populate(objectMap, "statusDetails", g.StatusDetails)
	populate(objectMap, "tcp", g.TCP)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GatewayResourceDescription.
func (g GatewayResourceDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", g.ID)
	populate(objectMap, "location", g.Location)
	populate(objectMap, "name", g.Name)
	populate(objectMap, "properties", g.Properties)
	populate(objectMap, "tags", g.Tags)
	populate(objectMap, "type", g.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type GatewayResourceProperties.
func (g GatewayResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", g.Description)
	populate(objectMap, "destinationNetwork", g.DestinationNetwork)
	populate(objectMap, "http", g.HTTP)
	populate(objectMap, "ipAddress", g.IPAddress)
	populate(objectMap, "provisioningState", g.ProvisioningState)
	populate(objectMap, "sourceNetwork", g.SourceNetwork)
	populate(objectMap, "status", g.Status)
	populate(objectMap, "statusDetails", g.StatusDetails)
	populate(objectMap, "tcp", g.TCP)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HTTPConfig.
func (h HTTPConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "hosts", h.Hosts)
	populate(objectMap, "name", h.Name)
	populate(objectMap, "port", h.Port)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HTTPHostConfig.
func (h HTTPHostConfig) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "name", h.Name)
	populate(objectMap, "routes", h.Routes)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type HTTPRouteMatchRule.
func (h HTTPRouteMatchRule) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "headers", h.Headers)
	populate(objectMap, "path", h.Path)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type InlinedValueSecretResourceProperties.
func (i InlinedValueSecretResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contentType", i.ContentType)
	populate(objectMap, "description", i.Description)
	objectMap["kind"] = SecretKindInlinedValue
	populate(objectMap, "provisioningState", i.ProvisioningState)
	populate(objectMap, "status", i.Status)
	populate(objectMap, "statusDetails", i.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type InlinedValueSecretResourceProperties.
func (i *InlinedValueSecretResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", i, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentType":
			err = unpopulate(val, "ContentType", &i.ContentType)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &i.Description)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &i.Kind)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &i.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &i.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &i.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", i, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type LocalNetworkResourceProperties.
func (l LocalNetworkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", l.Description)
	objectMap["kind"] = NetworkKindLocal
	populate(objectMap, "networkAddressPrefix", l.NetworkAddressPrefix)
	populate(objectMap, "provisioningState", l.ProvisioningState)
	populate(objectMap, "status", l.Status)
	populate(objectMap, "statusDetails", l.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type LocalNetworkResourceProperties.
func (l *LocalNetworkResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", l, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &l.Description)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &l.Kind)
			delete(rawMsg, key)
		case "networkAddressPrefix":
			err = unpopulate(val, "NetworkAddressPrefix", &l.NetworkAddressPrefix)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &l.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &l.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &l.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", l, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NetworkRef.
func (n NetworkRef) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "endpointRefs", n.EndpointRefs)
	populate(objectMap, "name", n.Name)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type NetworkResourceDescription.
func (n NetworkResourceDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", n.ID)
	populate(objectMap, "location", n.Location)
	populate(objectMap, "name", n.Name)
	populate(objectMap, "properties", n.Properties)
	populate(objectMap, "tags", n.Tags)
	populate(objectMap, "type", n.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkResourceDescription.
func (n *NetworkResourceDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &n.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &n.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &n.Name)
			delete(rawMsg, key)
		case "properties":
			n.Properties, err = unmarshalNetworkResourcePropertiesClassification(val)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &n.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &n.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type NetworkResourceProperties.
func (n NetworkResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", n.Description)
	objectMap["kind"] = "NetworkResourceProperties"
	populate(objectMap, "provisioningState", n.ProvisioningState)
	populate(objectMap, "status", n.Status)
	populate(objectMap, "statusDetails", n.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type NetworkResourceProperties.
func (n *NetworkResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", n, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, "Description", &n.Description)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &n.Kind)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &n.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &n.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &n.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", n, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SecretResourceDescription.
func (s SecretResourceDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecretResourceDescription.
func (s *SecretResourceDescription) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "id":
			err = unpopulate(val, "ID", &s.ID)
			delete(rawMsg, key)
		case "location":
			err = unpopulate(val, "Location", &s.Location)
			delete(rawMsg, key)
		case "name":
			err = unpopulate(val, "Name", &s.Name)
			delete(rawMsg, key)
		case "properties":
			s.Properties, err = unmarshalSecretResourcePropertiesClassification(val)
			delete(rawMsg, key)
		case "tags":
			err = unpopulate(val, "Tags", &s.Tags)
			delete(rawMsg, key)
		case "type":
			err = unpopulate(val, "Type", &s.Type)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SecretResourceProperties.
func (s SecretResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "contentType", s.ContentType)
	populate(objectMap, "description", s.Description)
	objectMap["kind"] = "SecretResourceProperties"
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "statusDetails", s.StatusDetails)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SecretResourceProperties.
func (s *SecretResourceProperties) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return fmt.Errorf("unmarshalling type %T: %v", s, err)
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "contentType":
			err = unpopulate(val, "ContentType", &s.ContentType)
			delete(rawMsg, key)
		case "description":
			err = unpopulate(val, "Description", &s.Description)
			delete(rawMsg, key)
		case "kind":
			err = unpopulate(val, "Kind", &s.Kind)
			delete(rawMsg, key)
		case "provisioningState":
			err = unpopulate(val, "ProvisioningState", &s.ProvisioningState)
			delete(rawMsg, key)
		case "status":
			err = unpopulate(val, "Status", &s.Status)
			delete(rawMsg, key)
		case "statusDetails":
			err = unpopulate(val, "StatusDetails", &s.StatusDetails)
			delete(rawMsg, key)
		}
		if err != nil {
			return fmt.Errorf("unmarshalling type %T: %v", s, err)
		}
	}
	return nil
}

// MarshalJSON implements the json.Marshaller interface for type SecretValueResourceDescription.
func (s SecretValueResourceDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", s.ID)
	populate(objectMap, "location", s.Location)
	populate(objectMap, "name", s.Name)
	populate(objectMap, "properties", s.Properties)
	populate(objectMap, "tags", s.Tags)
	populate(objectMap, "type", s.Type)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceProperties.
func (s ServiceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoScalingPolicies", s.AutoScalingPolicies)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "healthState", s.HealthState)
	populate(objectMap, "replicaCount", s.ReplicaCount)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "statusDetails", s.StatusDetails)
	populate(objectMap, "unhealthyEvaluation", s.UnhealthyEvaluation)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceReplicaDescription.
func (s ServiceReplicaDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "codePackages", s.CodePackages)
	populate(objectMap, "diagnostics", s.Diagnostics)
	populate(objectMap, "networkRefs", s.NetworkRefs)
	populate(objectMap, "osType", s.OSType)
	populate(objectMap, "replicaName", s.ReplicaName)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceReplicaProperties.
func (s ServiceReplicaProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "codePackages", s.CodePackages)
	populate(objectMap, "diagnostics", s.Diagnostics)
	populate(objectMap, "networkRefs", s.NetworkRefs)
	populate(objectMap, "osType", s.OSType)
	return json.Marshal(objectMap)
}

// MarshalJSON implements the json.Marshaller interface for type ServiceResourceProperties.
func (s ServiceResourceProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "autoScalingPolicies", s.AutoScalingPolicies)
	populate(objectMap, "codePackages", s.CodePackages)
	populate(objectMap, "description", s.Description)
	populate(objectMap, "diagnostics", s.Diagnostics)
	populate(objectMap, "healthState", s.HealthState)
	populate(objectMap, "networkRefs", s.NetworkRefs)
	populate(objectMap, "osType", s.OSType)
	populate(objectMap, "provisioningState", s.ProvisioningState)
	populate(objectMap, "replicaCount", s.ReplicaCount)
	populate(objectMap, "status", s.Status)
	populate(objectMap, "statusDetails", s.StatusDetails)
	populate(objectMap, "unhealthyEvaluation", s.UnhealthyEvaluation)
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

// MarshalJSON implements the json.Marshaller interface for type VolumeResourceDescription.
func (v VolumeResourceDescription) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", v.ID)
	populate(objectMap, "location", v.Location)
	populate(objectMap, "name", v.Name)
	populate(objectMap, "properties", v.Properties)
	populate(objectMap, "tags", v.Tags)
	populate(objectMap, "type", v.Type)
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
