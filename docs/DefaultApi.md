# \DefaultApi

All URIs are relative to *http://localhost:31453*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClusterHealth**](DefaultApi.md#ClusterHealth) | **Get** /2/clusters/{clusterUUID}/health | Get health of cluster
[**CreateCluster**](DefaultApi.md#CreateCluster) | **Post** /2/clusters | Create a cluster with the given specification
[**CreateClusterAuthz**](DefaultApi.md#CreateClusterAuthz) | **Post** /2/clusters/{clusterID}/authz | Add authorization for a cluster
[**CreateHelmChart**](DefaultApi.md#CreateHelmChart) | **Post** /2/clusters/{clusterUUID}/helmcharts | Create a helmChart for cluster with the given specification
[**DeleteCluster**](DefaultApi.md#DeleteCluster) | **Delete** /2/clusters/{clusterUUID} | Delete a cluster
[**DeleteClusterAuthz**](DefaultApi.md#DeleteClusterAuthz) | **Delete** /2/clusters/{clusterID}/authz/{authID} | Delete authorization for a cluster
[**DeleteHelmChart**](DefaultApi.md#DeleteHelmChart) | **Delete** /2/clusters/{clusterUUID}/helmcharts/{HelmChartUUID} | Delete helm chart for cluster
[**GetClusterByName**](DefaultApi.md#GetClusterByName) | **Get** /2/clusters/{clusterName} | Get a cluster by name
[**GetClusterEnv**](DefaultApi.md#GetClusterEnv) | **Get** /2/clusters/{clusterUUID}/env | Get cluster environment
[**GetClusters**](DefaultApi.md#GetClusters) | **Get** /2/clusters | Get all clusters
[**GetDashboard**](DefaultApi.md#GetDashboard) | **Get** /2/clusters/{clusterUUID}/dashboard | Get dashboard
[**GetHelmChartsByClusterUUID**](DefaultApi.md#GetHelmChartsByClusterUUID) | **Get** /2/clusters/{clusterUUID}/helmcharts | Get HelmCharts object for a given cluster
[**ListClusterAuthz**](DefaultApi.md#ListClusterAuthz) | **Get** /2/clusters/{clusterID}/authz | List authorizations for a cluster
[**PatchCluster**](DefaultApi.md#PatchCluster) | **Patch** /2/clusters/{clusterUUID} | Patch a cluster
[**UpdateCluster**](DefaultApi.md#UpdateCluster) | **Put** /2/clusters/{clusterUUID} | Update a cluster
[**UpgradeCluster**](DefaultApi.md#UpgradeCluster) | **Patch** /2/clusters/{clusterUUID}/upgrade | Upgrade a cluster


# **ClusterHealth**
> TypesClusterHealth ClusterHealth(ctx, clusterUUID)
Get health of cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterUUID** | **string**| Cluster UUID | 

### Return type

[**TypesClusterHealth**](types.ClusterHealth.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateCluster**
> CreateCluster(ctx, body)
Create a cluster with the given specification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesCluster**](TypesCluster.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateClusterAuthz**
> CreateClusterAuthz(ctx, clusterID, body)
Add authorization for a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**| Cluster UUID | 
  **body** | [**ApiAddAuthorization**](ApiAddAuthorization.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateHelmChart**
> CreateHelmChart(ctx, clusterUUID, body)
Create a helmChart for cluster with the given specification

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterUUID** | **string**| Cluster UUID | 
  **body** | [**TypesHelmChart**](TypesHelmChart.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteCluster**
> DeleteCluster(ctx, clusterUUID)
Delete a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterUUID** | **string**| Cluster UUID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteClusterAuthz**
> DeleteClusterAuthz(ctx, clusterID, authID)
Delete authorization for a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**| Cluster UUID | 
  **authID** | **string**| Authorization UUID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteHelmChart**
> DeleteHelmChart(ctx, helmChartUUID, clusterUUID)
Delete helm chart for cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **helmChartUUID** | **string**| HelmChartUUID | 
  **clusterUUID** | **string**| Cluster UUID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterByName**
> TypesCluster GetClusterByName(ctx, clusterName)
Get a cluster by name

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterName** | **string**| Cluster Name | 

### Return type

[**TypesCluster**](types.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusterEnv**
> GetClusterEnv(ctx, clusterUUID)
Get cluster environment

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterUUID** | **string**| Cluster UUID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetClusters**
> TypesCluster GetClusters(ctx, )
Get all clusters

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TypesCluster**](types.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboard**
> GetDashboard(ctx, clusterUUID)
Get dashboard

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterUUID** | **string**| Cluster UUID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHelmChartsByClusterUUID**
> TypesHelmChart GetHelmChartsByClusterUUID(ctx, clusterUUD)
Get HelmCharts object for a given cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterUUD** | **string**| Cluster UUID | 

### Return type

[**TypesHelmChart**](types.HelmChart.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListClusterAuthz**
> ListClusterAuthz(ctx, clusterID)
List authorizations for a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **clusterID** | **string**| Cluster UUID | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchCluster**
> TypesCluster PatchCluster(ctx, body)
Patch a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesCluster**](TypesCluster.md)|  | 

### Return type

[**TypesCluster**](types.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCluster**
> TypesCluster UpdateCluster(ctx, body)
Update a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesCluster**](TypesCluster.md)|  | 

### Return type

[**TypesCluster**](types.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpgradeCluster**
> TypesCluster UpgradeCluster(ctx, body)
Upgrade a cluster

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TypesCluster**](TypesCluster.md)|  | 

### Return type

[**TypesCluster**](types.Cluster.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

