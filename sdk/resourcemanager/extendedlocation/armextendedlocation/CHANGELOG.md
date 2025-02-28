# Release History

## 1.1.0-beta.1 (2022-06-28)
### Features Added

- New function `*ResourceSyncRulesClient.BeginCreateOrUpdate(context.Context, string, string, string, ResourceSyncRule, *ResourceSyncRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ResourceSyncRulesClientCreateOrUpdateResponse], error)`
- New function `NewResourceSyncRulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ResourceSyncRulesClient, error)`
- New function `*ResourceSyncRulesClient.BeginUpdate(context.Context, string, string, string, PatchableResourceSyncRule, *ResourceSyncRulesClientBeginUpdateOptions) (*runtime.Poller[ResourceSyncRulesClientUpdateResponse], error)`
- New function `*ResourceSyncRulesClient.Get(context.Context, string, string, string, *ResourceSyncRulesClientGetOptions) (ResourceSyncRulesClientGetResponse, error)`
- New function `*ResourceSyncRulesClient.Delete(context.Context, string, string, string, *ResourceSyncRulesClientDeleteOptions) (ResourceSyncRulesClientDeleteResponse, error)`
- New function `*CustomLocationsClient.FindTargetResourceGroup(context.Context, string, string, CustomLocationFindTargetResourceGroupProperties, *CustomLocationsClientFindTargetResourceGroupOptions) (CustomLocationsClientFindTargetResourceGroupResponse, error)`
- New function `*ResourceSyncRulesClient.NewListByCustomLocationIDPager(string, string, *ResourceSyncRulesClientListByCustomLocationIDOptions) *runtime.Pager[ResourceSyncRulesClientListByCustomLocationIDResponse]`
- New struct `CustomLocationFindTargetResourceGroupProperties`
- New struct `CustomLocationFindTargetResourceGroupResult`
- New struct `CustomLocationsClientFindTargetResourceGroupOptions`
- New struct `CustomLocationsClientFindTargetResourceGroupResponse`
- New struct `MatchExpressionsProperties`
- New struct `PatchableResourceSyncRule`
- New struct `ResourceSyncRule`
- New struct `ResourceSyncRuleListResult`
- New struct `ResourceSyncRuleProperties`
- New struct `ResourceSyncRulePropertiesSelector`
- New struct `ResourceSyncRulesClient`
- New struct `ResourceSyncRulesClientBeginCreateOrUpdateOptions`
- New struct `ResourceSyncRulesClientBeginUpdateOptions`
- New struct `ResourceSyncRulesClientCreateOrUpdateResponse`
- New struct `ResourceSyncRulesClientDeleteOptions`
- New struct `ResourceSyncRulesClientDeleteResponse`
- New struct `ResourceSyncRulesClientGetOptions`
- New struct `ResourceSyncRulesClientGetResponse`
- New struct `ResourceSyncRulesClientListByCustomLocationIDOptions`
- New struct `ResourceSyncRulesClientListByCustomLocationIDResponse`
- New struct `ResourceSyncRulesClientUpdateResponse`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).