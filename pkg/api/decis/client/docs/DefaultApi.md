# \DefaultApi

All URIs are relative to *https://api.openshift.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDecision**](DefaultApi.md#CreateDecision) | **Post** /api/daas-api/v1/decisions | Create a new decision Request
[**CreateServiceAccount**](DefaultApi.md#CreateServiceAccount) | **Post** /api/daas-api/v1/serviceaccounts | Create a service account
[**DeleteDecisionById**](DefaultApi.md#DeleteDecisionById) | **Delete** /api/daas-api/v1/decisions/{id} | Delete a decision request by id
[**DeleteServiceAccount**](DefaultApi.md#DeleteServiceAccount) | **Delete** /api/daas-api/v1/serviceaccounts/{id} | Delete service account
[**GetDecisionById**](DefaultApi.md#GetDecisionById) | **Get** /api/daas-api/v1/decisions/{id} | Get a decision request by id
[**GetMetricsByInstantQuery**](DefaultApi.md#GetMetricsByInstantQuery) | **Get** /api/daas-api/v1/decisions/{id}/metrics/query | Get metrics with instant query by decision id.
[**GetMetricsByRangeQuery**](DefaultApi.md#GetMetricsByRangeQuery) | **Get** /api/daas-api/v1/decisions/{id}/metrics/query_range | Get metrics with timeseries range query by decision id.
[**GetServiceAccountById**](DefaultApi.md#GetServiceAccountById) | **Get** /api/daas-api/v1/serviceaccounts/{id} | get service account by id
[**ListCloudProviderRegions**](DefaultApi.md#ListCloudProviderRegions) | **Get** /api/daas-api/v1/cloud_providers/{id}/regions | Retrieves the list of supported regions of the supported cloud provider.
[**ListCloudProviders**](DefaultApi.md#ListCloudProviders) | **Get** /api/daas-api/v1/cloud_providers | Retrieves the list of supported cloud providers.
[**ListDecisions**](DefaultApi.md#ListDecisions) | **Get** /api/daas-api/v1/decisions | Returns a list of Decision requests
[**ListServiceAccounts**](DefaultApi.md#ListServiceAccounts) | **Get** /api/daas-api/v1/serviceaccounts | List service accounts
[**ResetServiceAccountCreds**](DefaultApi.md#ResetServiceAccountCreds) | **Post** /api/daas-api/v1/serviceaccounts/{id}/reset-credentials | reset credentials for the service account
[**ServiceStatus**](DefaultApi.md#ServiceStatus) | **Get** /api/daas-api/v1/status | Retrieves the status of resources e.g whether we have reached maximum service capacity
[**VersionMetadata**](DefaultApi.md#VersionMetadata) | **Get** /api/daas-api/v1 | Retrieves the version metadata



## CreateDecision

> DecisionRequest CreateDecision(ctx).Async(async).DecisionRequestPayload(decisionRequestPayload).Execute()

Create a new decision Request

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    async := true // bool | Perform the action in an asynchronous manner
    decisionRequestPayload := *openapiclient.NewDecisionRequestPayload("Name_example") // DecisionRequestPayload | Decision data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateDecision(context.Background()).Async(async).DecisionRequestPayload(decisionRequestPayload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateDecision``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDecision`: DecisionRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateDecision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **async** | **bool** | Perform the action in an asynchronous manner | 
 **decisionRequestPayload** | [**DecisionRequestPayload**](DecisionRequestPayload.md) | Decision data | 

### Return type

[**DecisionRequest**](DecisionRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceAccount

> ServiceAccount CreateServiceAccount(ctx).ServiceAccountRequest(serviceAccountRequest).Execute()

Create a service account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    serviceAccountRequest := *openapiclient.NewServiceAccountRequest("Name_example") // ServiceAccountRequest | service account request

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.CreateServiceAccount(context.Background()).ServiceAccountRequest(serviceAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.CreateServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateServiceAccount`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.CreateServiceAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceAccountRequest** | [**ServiceAccountRequest**](ServiceAccountRequest.md) | service account request | 

### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDecisionById

> Error DeleteDecisionById(ctx, id).Async(async).Execute()

Delete a decision request by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    async := true // bool | Perform the action in an asynchronous manner

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteDecisionById(context.Background(), id).Async(async).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteDecisionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDecisionById`: Error
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteDecisionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDecisionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **async** | **bool** | Perform the action in an asynchronous manner | 

### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteServiceAccount

> Error DeleteServiceAccount(ctx, id).Execute()

Delete service account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.DeleteServiceAccount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DeleteServiceAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteServiceAccount`: Error
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DeleteServiceAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteServiceAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Error**](Error.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDecisionById

> DecisionRequest GetDecisionById(ctx, id).Execute()

Get a decision request by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetDecisionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetDecisionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDecisionById`: DecisionRequest
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetDecisionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDecisionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DecisionRequest**](DecisionRequest.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsByInstantQuery

> MetricsInstantQueryList GetMetricsByInstantQuery(ctx, id).Filters(filters).Execute()

Get metrics with instant query by decision id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    filters := []string{"Inner_example"} // []string | List of metrics to fetch. Fetch all metrics when empty. List entries are decision internal metric names. (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMetricsByInstantQuery(context.Background(), id).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMetricsByInstantQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsByInstantQuery`: MetricsInstantQueryList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMetricsByInstantQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsByInstantQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filters** | **[]string** | List of metrics to fetch. Fetch all metrics when empty. List entries are decision internal metric names. | [default to []]

### Return type

[**MetricsInstantQueryList**](MetricsInstantQueryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMetricsByRangeQuery

> MetricsRangeQueryList GetMetricsByRangeQuery(ctx, id).Duration(duration).Interval(interval).Filters(filters).Execute()

Get metrics with timeseries range query by decision id.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    duration := int64(5) // int64 | The length of time in minutes over which to return the metrics. (default to 5)
    interval := int64(30) // int64 | The interval in seconds between data points. (default to 30)
    filters := []string{"Inner_example"} // []string | List of metrics to fetch. Fetch all metrics when empty. List entries are decision internal metric names. (optional) (default to [])

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetMetricsByRangeQuery(context.Background(), id).Duration(duration).Interval(interval).Filters(filters).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetMetricsByRangeQuery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsByRangeQuery`: MetricsRangeQueryList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetMetricsByRangeQuery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsByRangeQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duration** | **int64** | The length of time in minutes over which to return the metrics. | [default to 5]
 **interval** | **int64** | The interval in seconds between data points. | [default to 30]
 **filters** | **[]string** | List of metrics to fetch. Fetch all metrics when empty. List entries are decision internal metric names. | [default to []]

### Return type

[**MetricsRangeQueryList**](MetricsRangeQueryList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetServiceAccountById

> ServiceAccount GetServiceAccountById(ctx, id).Execute()

get service account by id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.GetServiceAccountById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetServiceAccountById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetServiceAccountById`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetServiceAccountById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetServiceAccountByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviderRegions

> CloudRegionList ListCloudProviderRegions(ctx, id).Page(page).Size(size).Execute()

Retrieves the list of supported regions of the supported cloud provider.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCloudProviderRegions(context.Background(), id).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCloudProviderRegions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudProviderRegions`: CloudRegionList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCloudProviderRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProviderRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 

### Return type

[**CloudRegionList**](CloudRegionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCloudProviders

> CloudProviderList ListCloudProviders(ctx).Page(page).Size(size).Execute()

Retrieves the list of supported cloud providers.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListCloudProviders(context.Background()).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListCloudProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCloudProviders`: CloudProviderList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListCloudProviders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCloudProvidersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 

### Return type

[**CloudProviderList**](CloudProviderList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDecisions

> DecisionList ListDecisions(ctx).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()

Returns a list of Decision requests

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    page := "1" // string | Page index (optional)
    size := "100" // string | Number of items in each page (optional)
    orderBy := "name asc" // string | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the decisionRequests fields. For example, in order to retrieve all decisions ordered by their name:  ```sql name asc ```  Or in order to retrieve all decisions ordered by their name _and_ created date:  ```sql name asc, created_at asc ```  If the parameter isn't provided, or if the value is empty, then the results will be ordered by name. (optional)
    search := "name = my-decision and cloud_provider = aws" // string | Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: cloud_provider, name, owner, region and status. Allowed comparators are `<>`, `=` or `LIKE`. Allowed joins are `AND` and `OR`, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve decision request with name equal `my-decision` and region equal `aws`, the value should be:  ``` name = my-decision and cloud_provider = aws ```  To retrieve decision request with its name starting with `my`, the value should be:  ``` name like my%25 ```  If the parameter isn't provided, or if the value is empty, then all the decisions that the user has permission to see will be returned.  Note. If the query is invalid, an error will be returned  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListDecisions(context.Background()).Page(page).Size(size).OrderBy(orderBy).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDecisions`: DecisionList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListDecisions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string** | Page index | 
 **size** | **string** | Number of items in each page | 
 **orderBy** | **string** | Specifies the order by criteria. The syntax of this parameter is similar to the syntax of the _order by_ clause of an SQL statement. Each query can be ordered by any of the decisionRequests fields. For example, in order to retrieve all decisions ordered by their name:  &#x60;&#x60;&#x60;sql name asc &#x60;&#x60;&#x60;  Or in order to retrieve all decisions ordered by their name _and_ created date:  &#x60;&#x60;&#x60;sql name asc, created_at asc &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then the results will be ordered by name. | 
 **search** | **string** | Search criteria.  The syntax of this parameter is similar to the syntax of the _where_ clause of an SQL statement. Allowed fields in the search are: cloud_provider, name, owner, region and status. Allowed comparators are &#x60;&lt;&gt;&#x60;, &#x60;&#x3D;&#x60; or &#x60;LIKE&#x60;. Allowed joins are &#x60;AND&#x60; and &#x60;OR&#x60;, however there is a limit of max 10 joins in the search query.  Examples:  To retrieve decision request with name equal &#x60;my-decision&#x60; and region equal &#x60;aws&#x60;, the value should be:  &#x60;&#x60;&#x60; name &#x3D; my-decision and cloud_provider &#x3D; aws &#x60;&#x60;&#x60;  To retrieve decision request with its name starting with &#x60;my&#x60;, the value should be:  &#x60;&#x60;&#x60; name like my%25 &#x60;&#x60;&#x60;  If the parameter isn&#39;t provided, or if the value is empty, then all the decisions that the user has permission to see will be returned.  Note. If the query is invalid, an error will be returned  | 

### Return type

[**DecisionList**](DecisionList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListServiceAccounts

> ServiceAccountList ListServiceAccounts(ctx).Execute()

List service accounts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ListServiceAccounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ListServiceAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListServiceAccounts`: ServiceAccountList
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ListServiceAccounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListServiceAccountsRequest struct via the builder pattern


### Return type

[**ServiceAccountList**](ServiceAccountList.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetServiceAccountCreds

> ServiceAccount ResetServiceAccountCreds(ctx, id).Execute()

reset credentials for the service account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The id of record

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ResetServiceAccountCreds(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ResetServiceAccountCreds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetServiceAccountCreds`: ServiceAccount
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ResetServiceAccountCreds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The id of record | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetServiceAccountCredsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ServiceAccount**](ServiceAccount.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ServiceStatus

> ServiceStatus ServiceStatus(ctx).Execute()

Retrieves the status of resources e.g whether we have reached maximum service capacity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.ServiceStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.ServiceStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ServiceStatus`: ServiceStatus
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.ServiceStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiServiceStatusRequest struct via the builder pattern


### Return type

[**ServiceStatus**](ServiceStatus.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VersionMetadata

> VersionMetadata VersionMetadata(ctx).Execute()

Retrieves the version metadata

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.VersionMetadata(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VersionMetadata``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VersionMetadata`: VersionMetadata
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VersionMetadata`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiVersionMetadataRequest struct via the builder pattern


### Return type

[**VersionMetadata**](VersionMetadata.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

