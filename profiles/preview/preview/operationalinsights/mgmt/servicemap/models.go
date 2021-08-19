//go:build go1.9
// +build go1.9

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package servicemap

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/preview/operationalinsights/mgmt/2015-11-01-preview/servicemap"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type Accuracy = original.Accuracy

const (
	Actual    Accuracy = original.Actual
	Estimated Accuracy = original.Estimated
)

type AzureCloudServiceRoleType = original.AzureCloudServiceRoleType

const (
	Unknown AzureCloudServiceRoleType = original.Unknown
	Web     AzureCloudServiceRoleType = original.Web
	Worker  AzureCloudServiceRoleType = original.Worker
)

type Bitness = original.Bitness

const (
	SixFourbit  Bitness = original.SixFourbit
	ThreeTwobit Bitness = original.ThreeTwobit
)

type ConnectionFailureState = original.ConnectionFailureState

const (
	Failed ConnectionFailureState = original.Failed
	Mixed  ConnectionFailureState = original.Mixed
	Ok     ConnectionFailureState = original.Ok
)

type HypervisorType = original.HypervisorType

const (
	HypervisorTypeHyperv  HypervisorType = original.HypervisorTypeHyperv
	HypervisorTypeUnknown HypervisorType = original.HypervisorTypeUnknown
)

type Kind = original.Kind

const (
	KindRefclientgroup      Kind = original.KindRefclientgroup
	KindRefmachine          Kind = original.KindRefmachine
	KindRefmachinewithhints Kind = original.KindRefmachinewithhints
	KindRefport             Kind = original.KindRefport
	KindRefprocess          Kind = original.KindRefprocess
	KindResourceReference   Kind = original.KindResourceReference
)

type KindBasicCoreResource = original.KindBasicCoreResource

const (
	KindClientGroup  KindBasicCoreResource = original.KindClientGroup
	KindCoreResource KindBasicCoreResource = original.KindCoreResource
	KindMachine      KindBasicCoreResource = original.KindMachine
	KindMachineGroup KindBasicCoreResource = original.KindMachineGroup
	KindPort         KindBasicCoreResource = original.KindPort
	KindProcess      KindBasicCoreResource = original.KindProcess
)

type KindBasicHostingConfiguration = original.KindBasicHostingConfiguration

const (
	KindHostingConfiguration KindBasicHostingConfiguration = original.KindHostingConfiguration
	KindProviderazure        KindBasicHostingConfiguration = original.KindProviderazure
)

type KindBasicMapRequest = original.KindBasicMapRequest

const (
	KindMapmachineGroupDependency  KindBasicMapRequest = original.KindMapmachineGroupDependency
	KindMapmachineListDependency   KindBasicMapRequest = original.KindMapmachineListDependency
	KindMapRequest                 KindBasicMapRequest = original.KindMapRequest
	KindMapsingleMachineDependency KindBasicMapRequest = original.KindMapsingleMachineDependency
	KindMultipleMachinesMapRequest KindBasicMapRequest = original.KindMultipleMachinesMapRequest
)

type KindBasicProcessHostingConfiguration = original.KindBasicProcessHostingConfiguration

const (
	KindBasicProcessHostingConfigurationKindProcessHostingConfiguration KindBasicProcessHostingConfiguration = original.KindBasicProcessHostingConfigurationKindProcessHostingConfiguration
	KindBasicProcessHostingConfigurationKindProviderazure               KindBasicProcessHostingConfiguration = original.KindBasicProcessHostingConfigurationKindProviderazure
)

type KindBasicRelationship = original.KindBasicRelationship

const (
	KindRelacceptor   KindBasicRelationship = original.KindRelacceptor
	KindRelationship  KindBasicRelationship = original.KindRelationship
	KindRelconnection KindBasicRelationship = original.KindRelconnection
)

type MachineGroupType = original.MachineGroupType

const (
	MachineGroupTypeAzureCs    MachineGroupType = original.MachineGroupTypeAzureCs
	MachineGroupTypeAzureSf    MachineGroupType = original.MachineGroupTypeAzureSf
	MachineGroupTypeAzureVmss  MachineGroupType = original.MachineGroupTypeAzureVmss
	MachineGroupTypeUnknown    MachineGroupType = original.MachineGroupTypeUnknown
	MachineGroupTypeUserStatic MachineGroupType = original.MachineGroupTypeUserStatic
)

type MachineRebootStatus = original.MachineRebootStatus

const (
	MachineRebootStatusNotRebooted MachineRebootStatus = original.MachineRebootStatusNotRebooted
	MachineRebootStatusRebooted    MachineRebootStatus = original.MachineRebootStatusRebooted
	MachineRebootStatusUnknown     MachineRebootStatus = original.MachineRebootStatusUnknown
)

type MonitoringState = original.MonitoringState

const (
	Discovered MonitoringState = original.Discovered
	Monitored  MonitoringState = original.Monitored
)

type OperatingSystemFamily = original.OperatingSystemFamily

const (
	OperatingSystemFamilyAix     OperatingSystemFamily = original.OperatingSystemFamilyAix
	OperatingSystemFamilyLinux   OperatingSystemFamily = original.OperatingSystemFamilyLinux
	OperatingSystemFamilySolaris OperatingSystemFamily = original.OperatingSystemFamilySolaris
	OperatingSystemFamilyUnknown OperatingSystemFamily = original.OperatingSystemFamilyUnknown
	OperatingSystemFamilyWindows OperatingSystemFamily = original.OperatingSystemFamilyWindows
)

type ProcessRole = original.ProcessRole

const (
	AppServer      ProcessRole = original.AppServer
	DatabaseServer ProcessRole = original.DatabaseServer
	LdapServer     ProcessRole = original.LdapServer
	SmbServer      ProcessRole = original.SmbServer
	WebServer      ProcessRole = original.WebServer
)

type Provider = original.Provider

const (
	Azure Provider = original.Azure
)

type Provider1 = original.Provider1

const (
	Provider1Azure Provider1 = original.Provider1Azure
)

type VirtualMachineType = original.VirtualMachineType

const (
	VirtualMachineTypeHyperv    VirtualMachineType = original.VirtualMachineTypeHyperv
	VirtualMachineTypeLdom      VirtualMachineType = original.VirtualMachineTypeLdom
	VirtualMachineTypeLpar      VirtualMachineType = original.VirtualMachineTypeLpar
	VirtualMachineTypeUnknown   VirtualMachineType = original.VirtualMachineTypeUnknown
	VirtualMachineTypeVirtualPc VirtualMachineType = original.VirtualMachineTypeVirtualPc
	VirtualMachineTypeVmware    VirtualMachineType = original.VirtualMachineTypeVmware
	VirtualMachineTypeXen       VirtualMachineType = original.VirtualMachineTypeXen
)

type VirtualizationState = original.VirtualizationState

const (
	VirtualizationStateHypervisor VirtualizationState = original.VirtualizationStateHypervisor
	VirtualizationStatePhysical   VirtualizationState = original.VirtualizationStatePhysical
	VirtualizationStateUnknown    VirtualizationState = original.VirtualizationStateUnknown
	VirtualizationStateVirtual    VirtualizationState = original.VirtualizationStateVirtual
)

type Acceptor = original.Acceptor
type AcceptorProperties = original.AcceptorProperties
type AgentConfiguration = original.AgentConfiguration
type AzureCloudServiceConfiguration = original.AzureCloudServiceConfiguration
type AzureHostingConfiguration = original.AzureHostingConfiguration
type AzureProcessHostingConfiguration = original.AzureProcessHostingConfiguration
type AzureServiceFabricClusterConfiguration = original.AzureServiceFabricClusterConfiguration
type AzureVMScaleSetConfiguration = original.AzureVMScaleSetConfiguration
type BaseClient = original.BaseClient
type BasicCoreResource = original.BasicCoreResource
type BasicHostingConfiguration = original.BasicHostingConfiguration
type BasicMapRequest = original.BasicMapRequest
type BasicMultipleMachinesMapRequest = original.BasicMultipleMachinesMapRequest
type BasicProcessHostingConfiguration = original.BasicProcessHostingConfiguration
type BasicRelationship = original.BasicRelationship
type BasicResourceReference = original.BasicResourceReference
type ClientGroup = original.ClientGroup
type ClientGroupMember = original.ClientGroupMember
type ClientGroupMemberProperties = original.ClientGroupMemberProperties
type ClientGroupMembersCollection = original.ClientGroupMembersCollection
type ClientGroupMembersCollectionIterator = original.ClientGroupMembersCollectionIterator
type ClientGroupMembersCollectionPage = original.ClientGroupMembersCollectionPage
type ClientGroupMembersCount = original.ClientGroupMembersCount
type ClientGroupProperties = original.ClientGroupProperties
type ClientGroupReference = original.ClientGroupReference
type ClientGroupsClient = original.ClientGroupsClient
type Connection = original.Connection
type ConnectionCollection = original.ConnectionCollection
type ConnectionCollectionIterator = original.ConnectionCollectionIterator
type ConnectionCollectionPage = original.ConnectionCollectionPage
type ConnectionProperties = original.ConnectionProperties
type CoreResource = original.CoreResource
type Error = original.Error
type ErrorResponse = original.ErrorResponse
type HostingConfiguration = original.HostingConfiguration
type HypervisorConfiguration = original.HypervisorConfiguration
type ImageConfiguration = original.ImageConfiguration
type Ipv4NetworkInterface = original.Ipv4NetworkInterface
type Ipv6NetworkInterface = original.Ipv6NetworkInterface
type Liveness = original.Liveness
type Machine = original.Machine
type MachineCollection = original.MachineCollection
type MachineCollectionIterator = original.MachineCollectionIterator
type MachineCollectionPage = original.MachineCollectionPage
type MachineCountsByOperatingSystem = original.MachineCountsByOperatingSystem
type MachineGroup = original.MachineGroup
type MachineGroupCollection = original.MachineGroupCollection
type MachineGroupCollectionIterator = original.MachineGroupCollectionIterator
type MachineGroupCollectionPage = original.MachineGroupCollectionPage
type MachineGroupMapRequest = original.MachineGroupMapRequest
type MachineGroupProperties = original.MachineGroupProperties
type MachineGroupsClient = original.MachineGroupsClient
type MachineListMapRequest = original.MachineListMapRequest
type MachineProperties = original.MachineProperties
type MachineReference = original.MachineReference
type MachineReferenceWithHints = original.MachineReferenceWithHints
type MachineReferenceWithHintsProperties = original.MachineReferenceWithHintsProperties
type MachineResourcesConfiguration = original.MachineResourcesConfiguration
type MachinesClient = original.MachinesClient
type MachinesSummary = original.MachinesSummary
type MachinesSummaryProperties = original.MachinesSummaryProperties
type Map = original.Map
type MapEdges = original.MapEdges
type MapNodes = original.MapNodes
type MapRequest = original.MapRequest
type MapResponse = original.MapResponse
type MapsClient = original.MapsClient
type MultipleMachinesMapRequest = original.MultipleMachinesMapRequest
type NetworkConfiguration = original.NetworkConfiguration
type OperatingSystemConfiguration = original.OperatingSystemConfiguration
type Port = original.Port
type PortCollection = original.PortCollection
type PortCollectionIterator = original.PortCollectionIterator
type PortCollectionPage = original.PortCollectionPage
type PortProperties = original.PortProperties
type PortReference = original.PortReference
type PortReferenceProperties = original.PortReferenceProperties
type PortsClient = original.PortsClient
type Process = original.Process
type ProcessCollection = original.ProcessCollection
type ProcessCollectionIterator = original.ProcessCollectionIterator
type ProcessCollectionPage = original.ProcessCollectionPage
type ProcessDetails = original.ProcessDetails
type ProcessHostedService = original.ProcessHostedService
type ProcessHostingConfiguration = original.ProcessHostingConfiguration
type ProcessProperties = original.ProcessProperties
type ProcessReference = original.ProcessReference
type ProcessReferenceProperties = original.ProcessReferenceProperties
type ProcessUser = original.ProcessUser
type ProcessesClient = original.ProcessesClient
type Relationship = original.Relationship
type RelationshipProperties = original.RelationshipProperties
type Resource = original.Resource
type ResourceReference = original.ResourceReference
type SingleMachineDependencyMapRequest = original.SingleMachineDependencyMapRequest
type SummariesClient = original.SummariesClient
type Summary = original.Summary
type SummaryProperties = original.SummaryProperties
type Timezone = original.Timezone
type VirtualMachineConfiguration = original.VirtualMachineConfiguration

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewClientGroupMembersCollectionIterator(page ClientGroupMembersCollectionPage) ClientGroupMembersCollectionIterator {
	return original.NewClientGroupMembersCollectionIterator(page)
}
func NewClientGroupMembersCollectionPage(cur ClientGroupMembersCollection, getNextPage func(context.Context, ClientGroupMembersCollection) (ClientGroupMembersCollection, error)) ClientGroupMembersCollectionPage {
	return original.NewClientGroupMembersCollectionPage(cur, getNextPage)
}
func NewClientGroupsClient(subscriptionID string) ClientGroupsClient {
	return original.NewClientGroupsClient(subscriptionID)
}
func NewClientGroupsClientWithBaseURI(baseURI string, subscriptionID string) ClientGroupsClient {
	return original.NewClientGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewConnectionCollectionIterator(page ConnectionCollectionPage) ConnectionCollectionIterator {
	return original.NewConnectionCollectionIterator(page)
}
func NewConnectionCollectionPage(cur ConnectionCollection, getNextPage func(context.Context, ConnectionCollection) (ConnectionCollection, error)) ConnectionCollectionPage {
	return original.NewConnectionCollectionPage(cur, getNextPage)
}
func NewMachineCollectionIterator(page MachineCollectionPage) MachineCollectionIterator {
	return original.NewMachineCollectionIterator(page)
}
func NewMachineCollectionPage(cur MachineCollection, getNextPage func(context.Context, MachineCollection) (MachineCollection, error)) MachineCollectionPage {
	return original.NewMachineCollectionPage(cur, getNextPage)
}
func NewMachineGroupCollectionIterator(page MachineGroupCollectionPage) MachineGroupCollectionIterator {
	return original.NewMachineGroupCollectionIterator(page)
}
func NewMachineGroupCollectionPage(cur MachineGroupCollection, getNextPage func(context.Context, MachineGroupCollection) (MachineGroupCollection, error)) MachineGroupCollectionPage {
	return original.NewMachineGroupCollectionPage(cur, getNextPage)
}
func NewMachineGroupsClient(subscriptionID string) MachineGroupsClient {
	return original.NewMachineGroupsClient(subscriptionID)
}
func NewMachineGroupsClientWithBaseURI(baseURI string, subscriptionID string) MachineGroupsClient {
	return original.NewMachineGroupsClientWithBaseURI(baseURI, subscriptionID)
}
func NewMachinesClient(subscriptionID string) MachinesClient {
	return original.NewMachinesClient(subscriptionID)
}
func NewMachinesClientWithBaseURI(baseURI string, subscriptionID string) MachinesClient {
	return original.NewMachinesClientWithBaseURI(baseURI, subscriptionID)
}
func NewMapsClient(subscriptionID string) MapsClient {
	return original.NewMapsClient(subscriptionID)
}
func NewMapsClientWithBaseURI(baseURI string, subscriptionID string) MapsClient {
	return original.NewMapsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPortCollectionIterator(page PortCollectionPage) PortCollectionIterator {
	return original.NewPortCollectionIterator(page)
}
func NewPortCollectionPage(cur PortCollection, getNextPage func(context.Context, PortCollection) (PortCollection, error)) PortCollectionPage {
	return original.NewPortCollectionPage(cur, getNextPage)
}
func NewPortsClient(subscriptionID string) PortsClient {
	return original.NewPortsClient(subscriptionID)
}
func NewPortsClientWithBaseURI(baseURI string, subscriptionID string) PortsClient {
	return original.NewPortsClientWithBaseURI(baseURI, subscriptionID)
}
func NewProcessCollectionIterator(page ProcessCollectionPage) ProcessCollectionIterator {
	return original.NewProcessCollectionIterator(page)
}
func NewProcessCollectionPage(cur ProcessCollection, getNextPage func(context.Context, ProcessCollection) (ProcessCollection, error)) ProcessCollectionPage {
	return original.NewProcessCollectionPage(cur, getNextPage)
}
func NewProcessesClient(subscriptionID string) ProcessesClient {
	return original.NewProcessesClient(subscriptionID)
}
func NewProcessesClientWithBaseURI(baseURI string, subscriptionID string) ProcessesClient {
	return original.NewProcessesClientWithBaseURI(baseURI, subscriptionID)
}
func NewSummariesClient(subscriptionID string) SummariesClient {
	return original.NewSummariesClient(subscriptionID)
}
func NewSummariesClientWithBaseURI(baseURI string, subscriptionID string) SummariesClient {
	return original.NewSummariesClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleAccuracyValues() []Accuracy {
	return original.PossibleAccuracyValues()
}
func PossibleAzureCloudServiceRoleTypeValues() []AzureCloudServiceRoleType {
	return original.PossibleAzureCloudServiceRoleTypeValues()
}
func PossibleBitnessValues() []Bitness {
	return original.PossibleBitnessValues()
}
func PossibleConnectionFailureStateValues() []ConnectionFailureState {
	return original.PossibleConnectionFailureStateValues()
}
func PossibleHypervisorTypeValues() []HypervisorType {
	return original.PossibleHypervisorTypeValues()
}
func PossibleKindBasicCoreResourceValues() []KindBasicCoreResource {
	return original.PossibleKindBasicCoreResourceValues()
}
func PossibleKindBasicHostingConfigurationValues() []KindBasicHostingConfiguration {
	return original.PossibleKindBasicHostingConfigurationValues()
}
func PossibleKindBasicMapRequestValues() []KindBasicMapRequest {
	return original.PossibleKindBasicMapRequestValues()
}
func PossibleKindBasicProcessHostingConfigurationValues() []KindBasicProcessHostingConfiguration {
	return original.PossibleKindBasicProcessHostingConfigurationValues()
}
func PossibleKindBasicRelationshipValues() []KindBasicRelationship {
	return original.PossibleKindBasicRelationshipValues()
}
func PossibleKindValues() []Kind {
	return original.PossibleKindValues()
}
func PossibleMachineGroupTypeValues() []MachineGroupType {
	return original.PossibleMachineGroupTypeValues()
}
func PossibleMachineRebootStatusValues() []MachineRebootStatus {
	return original.PossibleMachineRebootStatusValues()
}
func PossibleMonitoringStateValues() []MonitoringState {
	return original.PossibleMonitoringStateValues()
}
func PossibleOperatingSystemFamilyValues() []OperatingSystemFamily {
	return original.PossibleOperatingSystemFamilyValues()
}
func PossibleProcessRoleValues() []ProcessRole {
	return original.PossibleProcessRoleValues()
}
func PossibleProvider1Values() []Provider1 {
	return original.PossibleProvider1Values()
}
func PossibleProviderValues() []Provider {
	return original.PossibleProviderValues()
}
func PossibleVirtualMachineTypeValues() []VirtualMachineType {
	return original.PossibleVirtualMachineTypeValues()
}
func PossibleVirtualizationStateValues() []VirtualizationState {
	return original.PossibleVirtualizationStateValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/preview"
}
func Version() string {
	return original.Version()
}
