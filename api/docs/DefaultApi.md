# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestCategoriesGet**](DefaultApi.md#RestCategoriesGet) | **Get** /rest/categories | fetching categories with pagination and optional filter queries
[**RestCategoriesPost**](DefaultApi.md#RestCategoriesPost) | **Post** /rest/categories | creating new categories
[**RestElasticSyncSyncSyncIdRunPost**](DefaultApi.md#RestElasticSyncSyncSyncIdRunPost) | **Post** /rest/elastic-sync/sync/{syncId}/run | runs an existing elastic sync
[**RestItemsAttributeValuesValueIdNamesPost**](DefaultApi.md#RestItemsAttributeValuesValueIdNamesPost) | **Post** /rest/items/attribute_values/{valueId}/names | creating attribute values names
[**RestItemsAttributesAttributeIdNamesPost**](DefaultApi.md#RestItemsAttributesAttributeIdNamesPost) | **Post** /rest/items/attributes/{attributeId}/names | creating names for attribute
[**RestItemsAttributesAttributeIdValuesGet**](DefaultApi.md#RestItemsAttributesAttributeIdValuesGet) | **Get** /rest/items/attributes/{attributeId}/values | fetching attributes values with pagination
[**RestItemsAttributesAttributeIdValuesPost**](DefaultApi.md#RestItemsAttributesAttributeIdValuesPost) | **Post** /rest/items/attributes/{attributeId}/values | creating new attribute values
[**RestItemsAttributesGet**](DefaultApi.md#RestItemsAttributesGet) | **Get** /rest/items/attributes | fetching attributes with pagination
[**RestItemsAttributesPost**](DefaultApi.md#RestItemsAttributesPost) | **Post** /rest/items/attributes | creating new attributes
[**RestItemsGet**](DefaultApi.md#RestItemsGet) | **Get** /rest/items | fetching items with optional extra data or filters given in query parameter
[**RestItemsItemIdVariationsPost**](DefaultApi.md#RestItemsItemIdVariationsPost) | **Post** /rest/items/{itemId}/variations | creating variations for an item
[**RestItemsManufacturersGet**](DefaultApi.md#RestItemsManufacturersGet) | **Get** /rest/items/manufacturers | list manufacturers
[**RestItemsManufacturersPost**](DefaultApi.md#RestItemsManufacturersPost) | **Post** /rest/items/manufacturers | creates a manufacturer
[**RestItemsPost**](DefaultApi.md#RestItemsPost) | **Post** /rest/items | creates items
[**RestItemsVariationsGet**](DefaultApi.md#RestItemsVariationsGet) | **Get** /rest/items/variations | fetching variations with pagination and the possibility to fetch additional related data by &#39;with&#39; query param
[**RestItemsVariationsPut**](DefaultApi.md#RestItemsVariationsPut) | **Put** /rest/items/variations | updating up to 50 variations
[**RestLoginPost**](DefaultApi.md#RestLoginPost) | **Post** /rest/login | login for authentication at rest api
[**RestPimVariationsGet**](DefaultApi.md#RestPimVariationsGet) | **Get** /rest/pim/variations | Lists variations
[**RestPimVariationsPut**](DefaultApi.md#RestPimVariationsPut) | **Put** /rest/pim/variations | Create a list of variations and their related data
[**RestStockmanagementWarehousesWarehouseIdStockCorrectionPut**](DefaultApi.md#RestStockmanagementWarehousesWarehouseIdStockCorrectionPut) | **Put** /rest/stockmanagement/warehouses/{warehouseId}/stock/correction | set stocks



## RestCategoriesGet

> CategoryPagination RestCategoriesGet(ctx).Page(page).Execute()

fetching categories with pagination and optional filter queries



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
    page := int32(56) // int32 | Limits the results to a specific page. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestCategoriesGet(context.Background()).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestCategoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestCategoriesGet`: CategoryPagination
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestCategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Limits the results to a specific page. | 

### Return type

[**CategoryPagination**](CategoryPagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestCategoriesPost

> []Category RestCategoriesPost(ctx).Category(category).Execute()

creating new categories



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
    category := []openapiclient.Category{*openapiclient.NewCategory()} // []Category |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestCategoriesPost(context.Background()).Category(category).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestCategoriesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestCategoriesPost`: []Category
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestCategoriesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestCategoriesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | [**[]Category**](Category.md) |  | 

### Return type

[**[]Category**](Category.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestElasticSyncSyncSyncIdRunPost

> RestElasticSyncSyncSyncIdRunPost(ctx, syncId).Execute()

runs an existing elastic sync



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
    syncId := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestElasticSyncSyncSyncIdRunPost(context.Background(), syncId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestElasticSyncSyncSyncIdRunPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**syncId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestElasticSyncSyncSyncIdRunPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsAttributeValuesValueIdNamesPost

> AttributeName RestItemsAttributeValuesValueIdNamesPost(ctx, valueId).AttributeName(attributeName).Execute()

creating attribute values names



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
    valueId := int32(56) // int32 | 
    attributeName := *openapiclient.NewAttributeName() // AttributeName |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsAttributeValuesValueIdNamesPost(context.Background(), valueId).AttributeName(attributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsAttributeValuesValueIdNamesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsAttributeValuesValueIdNamesPost`: AttributeName
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsAttributeValuesValueIdNamesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**valueId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsAttributeValuesValueIdNamesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeName** | [**AttributeName**](AttributeName.md) |  | 

### Return type

[**AttributeName**](AttributeName.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsAttributesAttributeIdNamesPost

> AttributeName RestItemsAttributesAttributeIdNamesPost(ctx, attributeId).AttributeName(attributeName).Execute()

creating names for attribute



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
    attributeId := int32(56) // int32 | 
    attributeName := *openapiclient.NewAttributeName() // AttributeName |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsAttributesAttributeIdNamesPost(context.Background(), attributeId).AttributeName(attributeName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsAttributesAttributeIdNamesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsAttributesAttributeIdNamesPost`: AttributeName
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsAttributesAttributeIdNamesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsAttributesAttributeIdNamesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeName** | [**AttributeName**](AttributeName.md) |  | 

### Return type

[**AttributeName**](AttributeName.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsAttributesAttributeIdValuesGet

> AttributeValuePagination RestItemsAttributesAttributeIdValuesGet(ctx, attributeId).Page(page).With(with).Execute()

fetching attributes values with pagination



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
    attributeId := int32(56) // int32 | 
    page := int32(56) // int32 | Limits the results to a specific page. (optional)
    with := []string{"With_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsAttributesAttributeIdValuesGet(context.Background(), attributeId).Page(page).With(with).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsAttributesAttributeIdValuesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsAttributesAttributeIdValuesGet`: AttributeValuePagination
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsAttributesAttributeIdValuesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsAttributesAttributeIdValuesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Limits the results to a specific page. | 
 **with** | **[]string** |  | 

### Return type

[**AttributeValuePagination**](AttributeValuePagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsAttributesAttributeIdValuesPost

> AttributeValue RestItemsAttributesAttributeIdValuesPost(ctx, attributeId).AttributeValue(attributeValue).Execute()

creating new attribute values



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
    attributeId := int32(56) // int32 | 
    attributeValue := *openapiclient.NewAttributeValue() // AttributeValue |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsAttributesAttributeIdValuesPost(context.Background(), attributeId).AttributeValue(attributeValue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsAttributesAttributeIdValuesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsAttributesAttributeIdValuesPost`: AttributeValue
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsAttributesAttributeIdValuesPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**attributeId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsAttributesAttributeIdValuesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributeValue** | [**AttributeValue**](AttributeValue.md) |  | 

### Return type

[**AttributeValue**](AttributeValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsAttributesGet

> AttributePagination RestItemsAttributesGet(ctx).Page(page).With(with).Execute()

fetching attributes with pagination



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
    page := int32(56) // int32 | Limits the results to a specific page. (optional)
    with := []string{"With_example"} // []string |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsAttributesGet(context.Background()).Page(page).With(with).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsAttributesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsAttributesGet`: AttributePagination
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsAttributesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsAttributesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Limits the results to a specific page. | 
 **with** | **[]string** |  | 

### Return type

[**AttributePagination**](AttributePagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsAttributesPost

> Attribute RestItemsAttributesPost(ctx).Attribute(attribute).Execute()

creating new attributes



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
    attribute := *openapiclient.NewAttribute() // Attribute |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsAttributesPost(context.Background()).Attribute(attribute).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsAttributesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsAttributesPost`: Attribute
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsAttributesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsAttributesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **attribute** | [**Attribute**](Attribute.md) |  | 

### Return type

[**Attribute**](Attribute.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsGet

> []ItemPagination RestItemsGet(ctx).With(with).Lang(lang).Page(page).ItemsPerPage(itemsPerPage).Name(name).ManufacturerId(manufacturerId).Id(id).FlagOne(flagOne).FlagTwo(flagTwo).UpdatedBetween(updatedBetween).VariationUpdatedBetween(variationUpdatedBetween).VariationRelatedUpdatedBetween(variationRelatedUpdatedBetween).Or(or).Execute()

fetching items with optional extra data or filters given in query parameter



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
    with := []string{"With_example"} // []string | Includes the specified variation information in the results (optional)
    lang := "lang_example" // string | The language of the variation information. (optional)
    page := int32(56) // int32 | Limits the results to a specific page. (optional)
    itemsPerPage := int32(56) // int32 | Limits the number of results listed per page to a specific number. (optional)
    name := "name_example" // string | Filter restricts the list of results to items with the specified item name. (optional)
    manufacturerId := int32(56) // int32 | Filter restricts the list of results to items with the specified manufacturerId. (optional)
    id := []int32{int32(123)} // []int32 | Filter restricts the list of results to items with the specified ID. More than one ID should be separated by commas. (optional)
    flagOne := int32(56) // int32 | Filter restricts the list of results to items with the specified flagOne. (optional)
    flagTwo := int32(56) // int32 | Filter restricts the list of results to items with the specified flagTwo. (optional)
    updatedBetween := "updatedBetween_example" // string | Filter restricts the list of results to items updated during the specified period. The end date (to) is optional. If no end date is specified, items updated between the start date (from) and the present will be listed. The dates can be specified as unix timestamps or in the ISO 8601 date format. Start date and optional end date are separated by a comma. For example, .../items?updatedBetween=1451606400,1456790400 will list items updated between 2016-01-01 and 2016-03-01. .../items?updatedBetween=1451606400 will list items updated since 2016-01-01. The PHP function strtotime is also supported. (optional)
    variationUpdatedBetween := "variationUpdatedBetween_example" // string | Filter restricts the list of results to items with variations that were updated during the specified period. The end date (to) is optional. If no end date is specified, items with variations updated between the start date (from) and the present will be listed. The dates can be specified as unix timestamps or in the ISO 8601 date format. Start date and optional end date are separated by a comma. For example, .../items?updatedBetween=1451606400,1456790400 will list items with variations that were updated between 2016-01-01 and 2016-03-01. .../items?updatedBetween=1451606400 will list items with variations that were updated since 2016-01-01. The PHP function strtotime is also supported. (optional)
    variationRelatedUpdatedBetween := "variationRelatedUpdatedBetween_example" // string | Filter restricts the list of results to items with variations for which related information was updated during the specified period. Related information is defined as information linked to the variation, i.e. barcodes, categories, images, markets, clients (stores), prices, suppliers, warehouses and the default category. See variationUpdatedBetween for supported formats. (optional)
    or := "or_example" // string | Filters can be defined in this param to link them via OR instead of AND. The syntax looks like the following: or=(updatedBetween=1573050718&varitionUpdatedBetween=1573050718). Everything in the brackets is written in the normal URL-Syntax. The or-param can be used multiple times if it is used like this: or[]=(...)&or[]=(...) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsGet(context.Background()).With(with).Lang(lang).Page(page).ItemsPerPage(itemsPerPage).Name(name).ManufacturerId(manufacturerId).Id(id).FlagOne(flagOne).FlagTwo(flagTwo).UpdatedBetween(updatedBetween).VariationUpdatedBetween(variationUpdatedBetween).VariationRelatedUpdatedBetween(variationRelatedUpdatedBetween).Or(or).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsGet`: []ItemPagination
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **with** | **[]string** | Includes the specified variation information in the results | 
 **lang** | **string** | The language of the variation information. | 
 **page** | **int32** | Limits the results to a specific page. | 
 **itemsPerPage** | **int32** | Limits the number of results listed per page to a specific number. | 
 **name** | **string** | Filter restricts the list of results to items with the specified item name. | 
 **manufacturerId** | **int32** | Filter restricts the list of results to items with the specified manufacturerId. | 
 **id** | **[]int32** | Filter restricts the list of results to items with the specified ID. More than one ID should be separated by commas. | 
 **flagOne** | **int32** | Filter restricts the list of results to items with the specified flagOne. | 
 **flagTwo** | **int32** | Filter restricts the list of results to items with the specified flagTwo. | 
 **updatedBetween** | **string** | Filter restricts the list of results to items updated during the specified period. The end date (to) is optional. If no end date is specified, items updated between the start date (from) and the present will be listed. The dates can be specified as unix timestamps or in the ISO 8601 date format. Start date and optional end date are separated by a comma. For example, .../items?updatedBetween&#x3D;1451606400,1456790400 will list items updated between 2016-01-01 and 2016-03-01. .../items?updatedBetween&#x3D;1451606400 will list items updated since 2016-01-01. The PHP function strtotime is also supported. | 
 **variationUpdatedBetween** | **string** | Filter restricts the list of results to items with variations that were updated during the specified period. The end date (to) is optional. If no end date is specified, items with variations updated between the start date (from) and the present will be listed. The dates can be specified as unix timestamps or in the ISO 8601 date format. Start date and optional end date are separated by a comma. For example, .../items?updatedBetween&#x3D;1451606400,1456790400 will list items with variations that were updated between 2016-01-01 and 2016-03-01. .../items?updatedBetween&#x3D;1451606400 will list items with variations that were updated since 2016-01-01. The PHP function strtotime is also supported. | 
 **variationRelatedUpdatedBetween** | **string** | Filter restricts the list of results to items with variations for which related information was updated during the specified period. Related information is defined as information linked to the variation, i.e. barcodes, categories, images, markets, clients (stores), prices, suppliers, warehouses and the default category. See variationUpdatedBetween for supported formats. | 
 **or** | **string** | Filters can be defined in this param to link them via OR instead of AND. The syntax looks like the following: or&#x3D;(updatedBetween&#x3D;1573050718&amp;varitionUpdatedBetween&#x3D;1573050718). Everything in the brackets is written in the normal URL-Syntax. The or-param can be used multiple times if it is used like this: or[]&#x3D;(...)&amp;or[]&#x3D;(...) | 

### Return type

[**[]ItemPagination**](ItemPagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsItemIdVariationsPost

> Variation RestItemsItemIdVariationsPost(ctx, itemId).Variation(variation).Execute()

creating variations for an item



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
    itemId := int32(56) // int32 | numeric itemId of the item where the variations belongs to
    variation := *openapiclient.NewVariation() // Variation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsItemIdVariationsPost(context.Background(), itemId).Variation(variation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsItemIdVariationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsItemIdVariationsPost`: Variation
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsItemIdVariationsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**itemId** | **int32** | numeric itemId of the item where the variations belongs to | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsItemIdVariationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **variation** | [**Variation**](Variation.md) |  | 

### Return type

[**Variation**](Variation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsManufacturersGet

> ManufacturersPagination RestItemsManufacturersGet(ctx).Page(page).With(with).UpdatedAt(updatedAt).Name(name).Execute()

list manufacturers



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
    page := int32(56) // int32 | Limits the results to a specific page. (optional)
    with := []string{"With_example"} // []string |  (optional)
    updatedAt := "updatedAt_example" // string | Filter restricts the list of results to records updated after the specified date. The date can be specified as unix timestamps or in the ISO 8601 date format. The PHP function strtotime is also supported. (optional)
    name := "name_example" // string | Filter restricts the list of results to records with specified name. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsManufacturersGet(context.Background()).Page(page).With(with).UpdatedAt(updatedAt).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsManufacturersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsManufacturersGet`: ManufacturersPagination
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsManufacturersGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsManufacturersGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Limits the results to a specific page. | 
 **with** | **[]string** |  | 
 **updatedAt** | **string** | Filter restricts the list of results to records updated after the specified date. The date can be specified as unix timestamps or in the ISO 8601 date format. The PHP function strtotime is also supported. | 
 **name** | **string** | Filter restricts the list of results to records with specified name. | 

### Return type

[**ManufacturersPagination**](ManufacturersPagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsManufacturersPost

> Manufacturer RestItemsManufacturersPost(ctx).Manufacturer(manufacturer).Execute()

creates a manufacturer



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
    manufacturer := *openapiclient.NewManufacturer() // Manufacturer |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsManufacturersPost(context.Background()).Manufacturer(manufacturer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsManufacturersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsManufacturersPost`: Manufacturer
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsManufacturersPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsManufacturersPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **manufacturer** | [**Manufacturer**](Manufacturer.md) |  | 

### Return type

[**Manufacturer**](Manufacturer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsPost

> CreateItemsResponse RestItemsPost(ctx).ItemRequest(itemRequest).Execute()

creates items



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
    itemRequest := []openapiclient.ItemRequest{*openapiclient.NewItemRequest(int32(123), int32(123), int32(123), []openapiclient.Variation{*openapiclient.NewVariation()})} // []ItemRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsPost(context.Background()).ItemRequest(itemRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsPost`: CreateItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemRequest** | [**[]ItemRequest**](ItemRequest.md) |  | 

### Return type

[**CreateItemsResponse**](CreateItemsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsVariationsGet

> VariationPagination RestItemsVariationsGet(ctx).With(with).Lang(lang).Page(page).ItemsPerPage(itemsPerPage).Id(id).ItemId(itemId).VariationTagId(variationTagId).ItemName(itemName).FlagOne(flagOne).FlagTwo(flagTwo).StoreSpecial(storeSpecial).CategoryId(categoryId).IsMain(isMain).IsActive(isActive).Barcode(barcode).NumberExact(numberExact).NumberFuzzy(numberFuzzy).IsBundle(isBundle).PlentyId(plentyId).ReferrerId(referrerId).SupplierNumber(supplierNumber).Sku(sku).ManufacturerId(manufacturerId).UpdatedBetween(updatedBetween).CreatedBetween(createdBetween).RelatedUpdatedBetween(relatedUpdatedBetween).ItemDescription(itemDescription).StockWarehouseId(stockWarehouseId).SupplierId(supplierId).Execute()

fetching variations with pagination and the possibility to fetch additional related data by 'with' query param



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
    with := []string{"With_example"} // []string |  (optional)
    lang := "lang_example" // string | The language of the variation information. (optional)
    page := int32(56) // int32 | Limits the results to a specific page. (optional)
    itemsPerPage := int32(56) // int32 | Limits the number of results listed per page to a specific number. (optional)
    id := []int32{int32(123)} // []int32 | Filter restricts the list of results to variations with the specified variation ID. An variation ID must be specified. More than one ID should be separated by commas. (optional)
    itemId := []int32{int32(123)} // []int32 | Filter restricts the list of results to variations with the specified item ID. An item ID must be specified. More than one ID should be separated by commas. (optional)
    variationTagId := []int32{int32(123)} // []int32 | Filter restricts the list of results to variations with the specified item name. An item name must be specified. (optional)
    itemName := "itemName_example" // string | Filter restricts the list of results to variations with the specified item name (optional)
    flagOne := int32(56) // int32 | Filter restricts the list of results to variations of items with the flag one. (optional)
    flagTwo := int32(56) // int32 | Filter restricts the list of results to variations of items with the flag two. (optional)
    storeSpecial := float32(8.14) // float32 | Filter restricts the list of results to variations of items with the specified store special. The following values are allowed: 0 (None), 1 (Special offer), 2 (New item), 3 (Top item) (optional)
    categoryId := int32(56) // int32 | Filter restricts the list of results to variations with the specified category id (optional)
    isMain := true // bool | Filter restricts the list of results to variations that are main variations. (optional)
    isActive := true // bool | Filter restricts the list of results to variations that are active. (optional)
    barcode := "barcode_example" // string | Filter restricts the list of results to variations with the specified barcode. (optional)
    numberExact := "numberExact_example" // string | Filter restricts the list of results to the variation with the variation number specified. (optional)
    numberFuzzy := "numberFuzzy_example" // string | Filter restricts the list of results to variations with numbers that contain the variation number specified (SQL LIKE operator). For example, if variations with variation numbers 1 to 400 exist in the system, filtering by 12 will list variation numbers 12, 112, 120-129, 212 and 312. (optional)
    isBundle := true // bool | Filter restricts the list of results to variations to which variations were added to create a bundle. (optional)
    plentyId := []int32{int32(123)} // []int32 | Filter restricts the list of results to variations that are visible in specified clients. Separate more than one client by commas. (optional)
    referrerId := []int32{int32(123)} // []int32 | Filter restricts the list of results to variations that are visible in specified markets. Separate more than one referrer by commas. (optional)
    supplierNumber := "supplierNumber_example" // string | Filter restricts the list of results to variations with the specified supplier number. (optional)
    sku := "sku_example" // string | Filter restricts the list of results to variations with the specified SKU. In additional, results can also be restricted to a specific referrer by specifying the referrer ID after a colon. Example: L0R3MIP5UM:104.1 (optional)
    manufacturerId := int32(56) // int32 | Filter restricts the list of results to variations with the specified manufacturer ID. (optional)
    updatedBetween := "updatedBetween_example" // string | Filter restricts the list of results to variations updated during the specified period. The end date (to) is optional. If no end date is specified, variations updated between the start date (from) and the present will be listed. The dates can be specified as unix timestamps or in the ISO 8601 date format. Start date and optional end date are separated by a comma. For example, .../variations?updatedBetween=1451606400,1456790400 will list variations updated between 2016-01-01 and 2016-03-01. .../variations?updatedBetween=1451606400 will list variations updated since 2016-01-01. The PHP function strtotime is also supported. (optional)
    createdBetween := "createdBetween_example" // string | Filter restricts the list of results to variations created during the specified period. The end date (to) is optional. If no end date is specified, variations created between the start date (from) and the present will be listed. The dates can be specified as unix timestamps or in the ISO 8601 date format. Start date and optional end date are separated by a comma. For example, .../variations?createdBetween=1451606400,1456790400 will list variations created between 2016-01-01 and 2016-03-01. .../variations?createdBetween=1451606400 will list variations created since 2016-01-01. The PHP function strtotime is also supported. (optional)
    relatedUpdatedBetween := "relatedUpdatedBetween_example" // string | Filter restricts the list of results to those variations for which related information was updated during the specified period. Related information is defined as information linked to the variation, i.e. barcodes, categories, images, markets, clients (stores), prices, suppliers, warehouses and the default category. See variationUpdatedBetween for supported formats. (optional)
    itemDescription := "itemDescription_example" // string | Filter restricts the list of results to variations with descriptions that contain the specified string. (optional)
    stockWarehouseId := "stockWarehouseId_example" // string | Filter restricts the list of results to variations which have physical stock on the given warehouse. (optional)
    supplierId := int32(56) // int32 | Filter restricts the list of results to variations with the specified supplier ID. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsVariationsGet(context.Background()).With(with).Lang(lang).Page(page).ItemsPerPage(itemsPerPage).Id(id).ItemId(itemId).VariationTagId(variationTagId).ItemName(itemName).FlagOne(flagOne).FlagTwo(flagTwo).StoreSpecial(storeSpecial).CategoryId(categoryId).IsMain(isMain).IsActive(isActive).Barcode(barcode).NumberExact(numberExact).NumberFuzzy(numberFuzzy).IsBundle(isBundle).PlentyId(plentyId).ReferrerId(referrerId).SupplierNumber(supplierNumber).Sku(sku).ManufacturerId(manufacturerId).UpdatedBetween(updatedBetween).CreatedBetween(createdBetween).RelatedUpdatedBetween(relatedUpdatedBetween).ItemDescription(itemDescription).StockWarehouseId(stockWarehouseId).SupplierId(supplierId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsVariationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsVariationsGet`: VariationPagination
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsVariationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsVariationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **with** | **[]string** |  | 
 **lang** | **string** | The language of the variation information. | 
 **page** | **int32** | Limits the results to a specific page. | 
 **itemsPerPage** | **int32** | Limits the number of results listed per page to a specific number. | 
 **id** | **[]int32** | Filter restricts the list of results to variations with the specified variation ID. An variation ID must be specified. More than one ID should be separated by commas. | 
 **itemId** | **[]int32** | Filter restricts the list of results to variations with the specified item ID. An item ID must be specified. More than one ID should be separated by commas. | 
 **variationTagId** | **[]int32** | Filter restricts the list of results to variations with the specified item name. An item name must be specified. | 
 **itemName** | **string** | Filter restricts the list of results to variations with the specified item name | 
 **flagOne** | **int32** | Filter restricts the list of results to variations of items with the flag one. | 
 **flagTwo** | **int32** | Filter restricts the list of results to variations of items with the flag two. | 
 **storeSpecial** | **float32** | Filter restricts the list of results to variations of items with the specified store special. The following values are allowed: 0 (None), 1 (Special offer), 2 (New item), 3 (Top item) | 
 **categoryId** | **int32** | Filter restricts the list of results to variations with the specified category id | 
 **isMain** | **bool** | Filter restricts the list of results to variations that are main variations. | 
 **isActive** | **bool** | Filter restricts the list of results to variations that are active. | 
 **barcode** | **string** | Filter restricts the list of results to variations with the specified barcode. | 
 **numberExact** | **string** | Filter restricts the list of results to the variation with the variation number specified. | 
 **numberFuzzy** | **string** | Filter restricts the list of results to variations with numbers that contain the variation number specified (SQL LIKE operator). For example, if variations with variation numbers 1 to 400 exist in the system, filtering by 12 will list variation numbers 12, 112, 120-129, 212 and 312. | 
 **isBundle** | **bool** | Filter restricts the list of results to variations to which variations were added to create a bundle. | 
 **plentyId** | **[]int32** | Filter restricts the list of results to variations that are visible in specified clients. Separate more than one client by commas. | 
 **referrerId** | **[]int32** | Filter restricts the list of results to variations that are visible in specified markets. Separate more than one referrer by commas. | 
 **supplierNumber** | **string** | Filter restricts the list of results to variations with the specified supplier number. | 
 **sku** | **string** | Filter restricts the list of results to variations with the specified SKU. In additional, results can also be restricted to a specific referrer by specifying the referrer ID after a colon. Example: L0R3MIP5UM:104.1 | 
 **manufacturerId** | **int32** | Filter restricts the list of results to variations with the specified manufacturer ID. | 
 **updatedBetween** | **string** | Filter restricts the list of results to variations updated during the specified period. The end date (to) is optional. If no end date is specified, variations updated between the start date (from) and the present will be listed. The dates can be specified as unix timestamps or in the ISO 8601 date format. Start date and optional end date are separated by a comma. For example, .../variations?updatedBetween&#x3D;1451606400,1456790400 will list variations updated between 2016-01-01 and 2016-03-01. .../variations?updatedBetween&#x3D;1451606400 will list variations updated since 2016-01-01. The PHP function strtotime is also supported. | 
 **createdBetween** | **string** | Filter restricts the list of results to variations created during the specified period. The end date (to) is optional. If no end date is specified, variations created between the start date (from) and the present will be listed. The dates can be specified as unix timestamps or in the ISO 8601 date format. Start date and optional end date are separated by a comma. For example, .../variations?createdBetween&#x3D;1451606400,1456790400 will list variations created between 2016-01-01 and 2016-03-01. .../variations?createdBetween&#x3D;1451606400 will list variations created since 2016-01-01. The PHP function strtotime is also supported. | 
 **relatedUpdatedBetween** | **string** | Filter restricts the list of results to those variations for which related information was updated during the specified period. Related information is defined as information linked to the variation, i.e. barcodes, categories, images, markets, clients (stores), prices, suppliers, warehouses and the default category. See variationUpdatedBetween for supported formats. | 
 **itemDescription** | **string** | Filter restricts the list of results to variations with descriptions that contain the specified string. | 
 **stockWarehouseId** | **string** | Filter restricts the list of results to variations which have physical stock on the given warehouse. | 
 **supplierId** | **int32** | Filter restricts the list of results to variations with the specified supplier ID. | 

### Return type

[**VariationPagination**](VariationPagination.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestItemsVariationsPut

> MultipleVariationsUpdatedResponse RestItemsVariationsPut(ctx).Variation(variation).Execute()

updating up to 50 variations



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
    variation := []openapiclient.Variation{*openapiclient.NewVariation()} // []Variation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsVariationsPut(context.Background()).Variation(variation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsVariationsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsVariationsPut`: MultipleVariationsUpdatedResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsVariationsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsVariationsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variation** | [**[]Variation**](Variation.md) |  | 

### Return type

[**MultipleVariationsUpdatedResponse**](MultipleVariationsUpdatedResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestLoginPost

> RestLoginResponse RestLoginPost(ctx).RestLoginBody(restLoginBody).Execute()

login for authentication at rest api



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
    restLoginBody := *openapiclient.NewRestLoginBody("Username_example", "Password_example") // RestLoginBody |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestLoginPost(context.Background()).RestLoginBody(restLoginBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestLoginPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestLoginPost`: RestLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestLoginPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestLoginPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **restLoginBody** | [**RestLoginBody**](RestLoginBody.md) |  | 

### Return type

[**RestLoginResponse**](RestLoginResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestPimVariationsGet

> InlineResponse200 RestPimVariationsGet(ctx).With(with).SortBy(sortBy).GroupBy(groupBy).Ids(ids).ItemId(itemId).ItemIds(itemIds).IsActive(isActive).IsMain(isMain).IsSalable(isSalable).SupplierId(supplierId).AvailabilityIds(availabilityIds).HasChildren(hasChildren).HasActiveChildren(hasActiveChildren).AttributeId(attributeId).AnyAttributeId(anyAttributeId).AllAttributeIds(allAttributeIds).AttributeValueId(attributeValueId).AnyAttributeValueId(anyAttributeValueId).AllAttributeValueIds(allAttributeValueIds).BarcodeCode(barcodeCode).BarcodeId(barcodeId).BundleType(bundleType).CategoryId(categoryId).AnyCategoryId(anyCategoryId).AllCategoryIds(allCategoryIds).AnyCharacteristicId(anyCharacteristicId).AllCharacteristicIds(allCharacteristicIds).ClientId(clientId).AnyClientId(anyClientId).AllClientIds(allClientIds).AutomaticClientVisibilities(automaticClientVisibilities).ImageHasMarketId(imageHasMarketId).Flag1(flag1).Flag2(flag2).ManufacturerId(manufacturerId).AnyManufacturerId(anyManufacturerId).ItemType(itemType).MarketId(marketId).AnyMarketId(anyMarketId).AllMarketIds(allMarketIds).PriceBetween(priceBetween).PriceBetweenById(priceBetweenById).AnySalesPriceId(anySalesPriceId).PropertySelectionId(propertySelectionId).AnyPropertySelectionId(anyPropertySelectionId).AllPropertySelectionIds(allPropertySelectionIds).HasNameInLanguage(hasNameInLanguage).CreatedAt(createdAt).UpdatedAt(updatedAt).ItemCreatedAt(itemCreatedAt).ItemUpdatedAt(itemUpdatedAt).AvailabilityUpdatedAt(availabilityUpdatedAt).StockUpdatedAt(stockUpdatedAt).Page(page).ItemsPerPage(itemsPerPage).Execute()

Lists variations



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
    with := []string{"With_example"} // []string | Includes the specified information in the results. More than one parameter should be separated by commas. The following parameters are available:<ul><li>'additionalSkus' = The additional skus of the variation.</li><li>'attributeValues' = The attribute values of the variation.</li><li>'attributeValues.attribute'/b> = Includes attributeValues. The attribute data to the related attribute ID.</li><li>'attributeValues.attributeValue' = Includes attributeValues. The attribute value data to the related attribute value ID.</li><li>'barcodes' = The barcodes of the variation.</li><li>'barcodes.barcode' = Includes barcodes. The barcode data to the related barcode ID.</li><li>'base' = The variation base.</li><li>'base.item' = Includes base. The item data of the variation.</li><li>'base.itemSerialNumber' = Includes base. The item serial numbers of the variation.</li><li>'base.feedback' = Includes base. The feedback of the variation.</li><li>'base.characteristics' = Includes base. The characteristics of the variation.</li><li>'base.crossSelling' = Includes base. The cross selling items of the variation.</li><li>'base.texts' = Includes base. The texts of the variation.</li><li>'base.availability' = Includes base. The availability data related to the variation's availability ID.</li><li>'base.images' = Includes base. The images linked to the item.</li><li>'base.shippingProfiles' = Includes base. The shipping profiles linked to the item.</li><li>'base.stock' = Includes base. The stock of the variation.</li><li>'base.stockStorageLocations' = Includes base. The stock storage locations of the variation.</li><li>'bundleComponents' = The bundle components of the variation.</li><li>'categories' = The categories of the variation.</li><li>'categories.category' = Includes categories. The related category data for each category ID.</li><li>'categories.categoryBranch' = Includes categories. The related category branch data for each category ID.</li><li>'clients' = The clients of the variation.</li><li>'defaultCategories' = The default categories of the variation</li><li>'defaultCategories.category' = Includes defaultCategories. The category data to the related category ID.</li><li>'images' = The images of the variation</li><li>'images.image' = Includes images. The image data to the related image ID.</li><li>'markets' = The markets of the variation.</li><li>'marketIdentNumbers' = The market ident numbers of the variation</li><li>'salesPrices' = The sales prices of the variation.</li><li>'salesPrices.salesPrice' = Includes salesPrices. The sales price data to the related sales price ID.</li><li>'skus' = The skus of the variation.</li><li>'supplier' = The supplier of the variation.</li><li>'supplier.supplier' = Includes supplier. The contact data to the related supplier ID.</li><li>'timestamps' = The timetamps of the variation.</li><li>'warehouses' = The warehouses of the variation</li><li>'warehouses.warehouse' = Includes warehouses. The warehouse data to the related warehouse ID.</li><li>'unit' = The unit of the variation</li><li>'unit.unit' = Includes unit. The unit data of the related unit ID.</li><li>'tags' = The tags of the variation.</li><li>'tags.tag' = Includes tags. The tag data to the related tag ID.</li><li>'properties' = The properties of the variation.</li><li>'properties.property' = Includes properties. The property data to the related property ID.</li></ul> (optional)
    sortBy := "sortBy_example" // string | Sorts the results. Append '_asc' or '_desc' to specify the sorting order. '_desc' is the default value if no other is specified. More than one parameter should be separated by commas. The following parameters are available:<ul><li>'id'</li><li>'itemId'</li><li>'isMain'</li><li>'position'</li><li>'availabilityId'</li><li>'createdAt'</li><li>'updatedAt'</li><li>'itemUpdatedAt'</li><li>'relatedUpdatedAt'</li><li>'variationName'</li><li>'number'</li></ul> (optional)
    groupBy := "groupBy_example" // string | Groups the result. The following parameters are available:<ul><li>'itemId' = Groups the result by the item ID.</li><li>'itemAttributeValue' = Groups the result by the attribute with the flag 'isGroupable'.</li></ul> (optional)
    ids := "ids_example" // string | Filter restricts the list of results to variations with the specified IDs. More than one parameter should be separated by commas. (optional)
    itemId := int32(56) // int32 | Filter restricts the list of results to variations with the specified item ID. (optional)
    itemIds := "itemIds_example" // string | Filter restricts the list of results to variations with the specified item IDs. More than one parameter should be separated by commas. (optional)
    isActive := true // bool | Filter restricts the list of results to variations which are active/inactive. (optional)
    isMain := true // bool | Filter restricts the list of results to variations which are main/not main. (optional)
    isSalable := true // bool | Filter restricts the list of results to variations which are salable. (optional)
    supplierId := int32(56) // int32 | Filter restricts the list of results to variations which have the given supplier ID. (optional)
    availabilityIds := "availabilityIds_example" // string | Filter restricts the list of results to variations with the specified availability IDs. More than one parameter should be separated by commas. (optional)
    hasChildren := true // bool | Filter restricts the list of results to variations which have children. (optional)
    hasActiveChildren := true // bool | Filter restricts the list of results to variations which have active children. (optional)
    attributeId := int32(56) // int32 |  Filter restricts the list of results to variations which have the specified attribute ID. (optional)
    anyAttributeId := "anyAttributeId_example" // string | Filter restricts the list of results to variations which have any of the specified attribute IDs. More than one parameter should be separated by commas. (optional)
    allAttributeIds := "allAttributeIds_example" // string | Filter restricts the list of results to variations which have all specified attribute IDs. More than one parameter should be separated by commas. (optional)
    attributeValueId := int32(56) // int32 |  Filter restricts the list of results to variations which have the specified attribute value ID. (optional)
    anyAttributeValueId := "anyAttributeValueId_example" // string |  Filter restricts the list of results to variations which have the any of the specified attribute value IDs. More than one parameter should be separated by commas. (optional)
    allAttributeValueIds := "allAttributeValueIds_example" // string | Filter restricts the list of results to variations which have all specified attribute value IDs. More than one parameter should be separated by commas. (optional)
    barcodeCode := "barcodeCode_example" // string | Filter restricts the list of results to variations which have a barcode with the specified code. (optional)
    barcodeId := int32(56) // int32 | Filter restricts the list of results to variations which have a barcode with the specified ID. (optional)
    bundleType := "bundleType_example" // string | Filter restricts the list of results to variations with the specified bundle type. (optional)
    categoryId := int32(56) // int32 | Filter restricts the list of results to variations which have the specified category ID. (optional)
    anyCategoryId := "anyCategoryId_example" // string | Filter restricts the list of results to variations which have any of the specified category IDs. More than one parameter should be separated by commas. (optional)
    allCategoryIds := "allCategoryIds_example" // string | Filter restricts the list of results to variations which have all specified category IDs. More than one parameter should be separated by commas. (optional)
    anyCharacteristicId := "anyCharacteristicId_example" // string | Filter restricts the list of results to variations which have any of the specified characteristic IDs. (optional)
    allCharacteristicIds := "allCharacteristicIds_example" // string | Filter restricts the list of results to variations which have all specified characteristic IDs. (optional)
    clientId := int32(56) // int32 | Filter restricts the list of results to variations which have the specified client ID. (optional)
    anyClientId := "anyClientId_example" // string | Filter restricts the list of results to variations which have any of the specified client IDs. More than one parameter should be separated by commas. (optional)
    allClientIds := "allClientIds_example" // string | Filter restricts the list of results to variations which have all specified client IDs. More than one parameter should be separated by commas. (optional)
    automaticClientVisibilities := "automaticClientVisibilities_example" // string | Filter restricts the list of results to variations which have any of the specified automatic client visibilities. More than one parameter should be separated by commas. (optional)
    imageHasMarketId := float32(8.14) // float32 |  Filter restricts the list of results to variations which have an image available for the specified market ID. (optional)
    flag1 := int32(56) // int32 | Filter restricts the list of results to variations with the specified flag one. (optional)
    flag2 := int32(56) // int32 | Filter restricts the list of results to variations with the specified flag two. (optional)
    manufacturerId := int32(56) // int32 | Filter restricts the list of results to variations with the specified manufacturer ID. (optional)
    anyManufacturerId := "anyManufacturerId_example" // string | Filter restricts the list of results to variations with any of the specified manufacturer IDs. (optional)
    itemType := "itemType_example" // string | Filter restricts the list of results to variations which have the specified item type. (optional)
    marketId := float32(8.14) // float32 | Filter restricts the list of results to variations which have the specified market ID. (optional)
    anyMarketId := "anyMarketId_example" // string | Filter restricts the list of results to variations which have any of the specified market IDs. (optional)
    allMarketIds := "allMarketIds_example" // string | Filter restricts the list of results to variations which have all specified market IDs. (optional)
    priceBetween := "priceBetween_example" // string | Filter restricts the list of results to variations which have a sales price between the specified minimum and maximum value. Minimum and maximum value should be separated by a comma. (optional)
    priceBetweenById := "priceBetweenById_example" // string | Filter restricts the list of results to variations where the specified sales price is between the specified minimum and maximum value. Sales price ID, Minimum and maximum value should be separated by a comma. (optional)
    anySalesPriceId := "anySalesPriceId_example" // string | Filter restricts the list of results to variations which have any of the specified sales price IDs. More than one parameter should be separated by commas. (optional)
    propertySelectionId := int32(56) // int32 | Filter restricts the list of results to variations which have the specified property selection ID. (optional)
    anyPropertySelectionId := "anyPropertySelectionId_example" // string | Filter restricts the list of results to variations which have any of the specified property selection IDs. (optional)
    allPropertySelectionIds := "allPropertySelectionIds_example" // string | Filter restricts the list of results to variations which have all specified property selection IDs. (optional)
    hasNameInLanguage := "hasNameInLanguage_example" // string | Filter restricts the list of results to variations which have a name in the specified language. (optional)
    createdAt := "createdAt_example" // string | Filter restricts the list of results to variations which have been created in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. (optional)
    updatedAt := "updatedAt_example" // string | Filter restricts the list of results to variations which have been updated in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. (optional)
    itemCreatedAt := "itemCreatedAt_example" // string | Filter restricts the list of results to variations whose item has been created in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. (optional)
    itemUpdatedAt := "itemUpdatedAt_example" // string | Filter restricts the list of results to variations whose item has been updated in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. (optional)
    availabilityUpdatedAt := "availabilityUpdatedAt_example" // string | Filter restricts the list of results to variations whose availablity has been updated in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. (optional)
    stockUpdatedAt := "stockUpdatedAt_example" // string | Filter restricts the list of results to variations whose stock has been updated in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. (optional)
    page := int32(56) // int32 | The requested page of results. Default value is 1. (optional)
    itemsPerPage := int32(56) // int32 | The number of results per page. Maximum value is 250. Default value is 50. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestPimVariationsGet(context.Background()).With(with).SortBy(sortBy).GroupBy(groupBy).Ids(ids).ItemId(itemId).ItemIds(itemIds).IsActive(isActive).IsMain(isMain).IsSalable(isSalable).SupplierId(supplierId).AvailabilityIds(availabilityIds).HasChildren(hasChildren).HasActiveChildren(hasActiveChildren).AttributeId(attributeId).AnyAttributeId(anyAttributeId).AllAttributeIds(allAttributeIds).AttributeValueId(attributeValueId).AnyAttributeValueId(anyAttributeValueId).AllAttributeValueIds(allAttributeValueIds).BarcodeCode(barcodeCode).BarcodeId(barcodeId).BundleType(bundleType).CategoryId(categoryId).AnyCategoryId(anyCategoryId).AllCategoryIds(allCategoryIds).AnyCharacteristicId(anyCharacteristicId).AllCharacteristicIds(allCharacteristicIds).ClientId(clientId).AnyClientId(anyClientId).AllClientIds(allClientIds).AutomaticClientVisibilities(automaticClientVisibilities).ImageHasMarketId(imageHasMarketId).Flag1(flag1).Flag2(flag2).ManufacturerId(manufacturerId).AnyManufacturerId(anyManufacturerId).ItemType(itemType).MarketId(marketId).AnyMarketId(anyMarketId).AllMarketIds(allMarketIds).PriceBetween(priceBetween).PriceBetweenById(priceBetweenById).AnySalesPriceId(anySalesPriceId).PropertySelectionId(propertySelectionId).AnyPropertySelectionId(anyPropertySelectionId).AllPropertySelectionIds(allPropertySelectionIds).HasNameInLanguage(hasNameInLanguage).CreatedAt(createdAt).UpdatedAt(updatedAt).ItemCreatedAt(itemCreatedAt).ItemUpdatedAt(itemUpdatedAt).AvailabilityUpdatedAt(availabilityUpdatedAt).StockUpdatedAt(stockUpdatedAt).Page(page).ItemsPerPage(itemsPerPage).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestPimVariationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestPimVariationsGet`: InlineResponse200
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestPimVariationsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestPimVariationsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **with** | **[]string** | Includes the specified information in the results. More than one parameter should be separated by commas. The following parameters are available:&lt;ul&gt;&lt;li&gt;&#39;additionalSkus&#39; &#x3D; The additional skus of the variation.&lt;/li&gt;&lt;li&gt;&#39;attributeValues&#39; &#x3D; The attribute values of the variation.&lt;/li&gt;&lt;li&gt;&#39;attributeValues.attribute&#39;/b&gt; &#x3D; Includes attributeValues. The attribute data to the related attribute ID.&lt;/li&gt;&lt;li&gt;&#39;attributeValues.attributeValue&#39; &#x3D; Includes attributeValues. The attribute value data to the related attribute value ID.&lt;/li&gt;&lt;li&gt;&#39;barcodes&#39; &#x3D; The barcodes of the variation.&lt;/li&gt;&lt;li&gt;&#39;barcodes.barcode&#39; &#x3D; Includes barcodes. The barcode data to the related barcode ID.&lt;/li&gt;&lt;li&gt;&#39;base&#39; &#x3D; The variation base.&lt;/li&gt;&lt;li&gt;&#39;base.item&#39; &#x3D; Includes base. The item data of the variation.&lt;/li&gt;&lt;li&gt;&#39;base.itemSerialNumber&#39; &#x3D; Includes base. The item serial numbers of the variation.&lt;/li&gt;&lt;li&gt;&#39;base.feedback&#39; &#x3D; Includes base. The feedback of the variation.&lt;/li&gt;&lt;li&gt;&#39;base.characteristics&#39; &#x3D; Includes base. The characteristics of the variation.&lt;/li&gt;&lt;li&gt;&#39;base.crossSelling&#39; &#x3D; Includes base. The cross selling items of the variation.&lt;/li&gt;&lt;li&gt;&#39;base.texts&#39; &#x3D; Includes base. The texts of the variation.&lt;/li&gt;&lt;li&gt;&#39;base.availability&#39; &#x3D; Includes base. The availability data related to the variation&#39;s availability ID.&lt;/li&gt;&lt;li&gt;&#39;base.images&#39; &#x3D; Includes base. The images linked to the item.&lt;/li&gt;&lt;li&gt;&#39;base.shippingProfiles&#39; &#x3D; Includes base. The shipping profiles linked to the item.&lt;/li&gt;&lt;li&gt;&#39;base.stock&#39; &#x3D; Includes base. The stock of the variation.&lt;/li&gt;&lt;li&gt;&#39;base.stockStorageLocations&#39; &#x3D; Includes base. The stock storage locations of the variation.&lt;/li&gt;&lt;li&gt;&#39;bundleComponents&#39; &#x3D; The bundle components of the variation.&lt;/li&gt;&lt;li&gt;&#39;categories&#39; &#x3D; The categories of the variation.&lt;/li&gt;&lt;li&gt;&#39;categories.category&#39; &#x3D; Includes categories. The related category data for each category ID.&lt;/li&gt;&lt;li&gt;&#39;categories.categoryBranch&#39; &#x3D; Includes categories. The related category branch data for each category ID.&lt;/li&gt;&lt;li&gt;&#39;clients&#39; &#x3D; The clients of the variation.&lt;/li&gt;&lt;li&gt;&#39;defaultCategories&#39; &#x3D; The default categories of the variation&lt;/li&gt;&lt;li&gt;&#39;defaultCategories.category&#39; &#x3D; Includes defaultCategories. The category data to the related category ID.&lt;/li&gt;&lt;li&gt;&#39;images&#39; &#x3D; The images of the variation&lt;/li&gt;&lt;li&gt;&#39;images.image&#39; &#x3D; Includes images. The image data to the related image ID.&lt;/li&gt;&lt;li&gt;&#39;markets&#39; &#x3D; The markets of the variation.&lt;/li&gt;&lt;li&gt;&#39;marketIdentNumbers&#39; &#x3D; The market ident numbers of the variation&lt;/li&gt;&lt;li&gt;&#39;salesPrices&#39; &#x3D; The sales prices of the variation.&lt;/li&gt;&lt;li&gt;&#39;salesPrices.salesPrice&#39; &#x3D; Includes salesPrices. The sales price data to the related sales price ID.&lt;/li&gt;&lt;li&gt;&#39;skus&#39; &#x3D; The skus of the variation.&lt;/li&gt;&lt;li&gt;&#39;supplier&#39; &#x3D; The supplier of the variation.&lt;/li&gt;&lt;li&gt;&#39;supplier.supplier&#39; &#x3D; Includes supplier. The contact data to the related supplier ID.&lt;/li&gt;&lt;li&gt;&#39;timestamps&#39; &#x3D; The timetamps of the variation.&lt;/li&gt;&lt;li&gt;&#39;warehouses&#39; &#x3D; The warehouses of the variation&lt;/li&gt;&lt;li&gt;&#39;warehouses.warehouse&#39; &#x3D; Includes warehouses. The warehouse data to the related warehouse ID.&lt;/li&gt;&lt;li&gt;&#39;unit&#39; &#x3D; The unit of the variation&lt;/li&gt;&lt;li&gt;&#39;unit.unit&#39; &#x3D; Includes unit. The unit data of the related unit ID.&lt;/li&gt;&lt;li&gt;&#39;tags&#39; &#x3D; The tags of the variation.&lt;/li&gt;&lt;li&gt;&#39;tags.tag&#39; &#x3D; Includes tags. The tag data to the related tag ID.&lt;/li&gt;&lt;li&gt;&#39;properties&#39; &#x3D; The properties of the variation.&lt;/li&gt;&lt;li&gt;&#39;properties.property&#39; &#x3D; Includes properties. The property data to the related property ID.&lt;/li&gt;&lt;/ul&gt; | 
 **sortBy** | **string** | Sorts the results. Append &#39;_asc&#39; or &#39;_desc&#39; to specify the sorting order. &#39;_desc&#39; is the default value if no other is specified. More than one parameter should be separated by commas. The following parameters are available:&lt;ul&gt;&lt;li&gt;&#39;id&#39;&lt;/li&gt;&lt;li&gt;&#39;itemId&#39;&lt;/li&gt;&lt;li&gt;&#39;isMain&#39;&lt;/li&gt;&lt;li&gt;&#39;position&#39;&lt;/li&gt;&lt;li&gt;&#39;availabilityId&#39;&lt;/li&gt;&lt;li&gt;&#39;createdAt&#39;&lt;/li&gt;&lt;li&gt;&#39;updatedAt&#39;&lt;/li&gt;&lt;li&gt;&#39;itemUpdatedAt&#39;&lt;/li&gt;&lt;li&gt;&#39;relatedUpdatedAt&#39;&lt;/li&gt;&lt;li&gt;&#39;variationName&#39;&lt;/li&gt;&lt;li&gt;&#39;number&#39;&lt;/li&gt;&lt;/ul&gt; | 
 **groupBy** | **string** | Groups the result. The following parameters are available:&lt;ul&gt;&lt;li&gt;&#39;itemId&#39; &#x3D; Groups the result by the item ID.&lt;/li&gt;&lt;li&gt;&#39;itemAttributeValue&#39; &#x3D; Groups the result by the attribute with the flag &#39;isGroupable&#39;.&lt;/li&gt;&lt;/ul&gt; | 
 **ids** | **string** | Filter restricts the list of results to variations with the specified IDs. More than one parameter should be separated by commas. | 
 **itemId** | **int32** | Filter restricts the list of results to variations with the specified item ID. | 
 **itemIds** | **string** | Filter restricts the list of results to variations with the specified item IDs. More than one parameter should be separated by commas. | 
 **isActive** | **bool** | Filter restricts the list of results to variations which are active/inactive. | 
 **isMain** | **bool** | Filter restricts the list of results to variations which are main/not main. | 
 **isSalable** | **bool** | Filter restricts the list of results to variations which are salable. | 
 **supplierId** | **int32** | Filter restricts the list of results to variations which have the given supplier ID. | 
 **availabilityIds** | **string** | Filter restricts the list of results to variations with the specified availability IDs. More than one parameter should be separated by commas. | 
 **hasChildren** | **bool** | Filter restricts the list of results to variations which have children. | 
 **hasActiveChildren** | **bool** | Filter restricts the list of results to variations which have active children. | 
 **attributeId** | **int32** |  Filter restricts the list of results to variations which have the specified attribute ID. | 
 **anyAttributeId** | **string** | Filter restricts the list of results to variations which have any of the specified attribute IDs. More than one parameter should be separated by commas. | 
 **allAttributeIds** | **string** | Filter restricts the list of results to variations which have all specified attribute IDs. More than one parameter should be separated by commas. | 
 **attributeValueId** | **int32** |  Filter restricts the list of results to variations which have the specified attribute value ID. | 
 **anyAttributeValueId** | **string** |  Filter restricts the list of results to variations which have the any of the specified attribute value IDs. More than one parameter should be separated by commas. | 
 **allAttributeValueIds** | **string** | Filter restricts the list of results to variations which have all specified attribute value IDs. More than one parameter should be separated by commas. | 
 **barcodeCode** | **string** | Filter restricts the list of results to variations which have a barcode with the specified code. | 
 **barcodeId** | **int32** | Filter restricts the list of results to variations which have a barcode with the specified ID. | 
 **bundleType** | **string** | Filter restricts the list of results to variations with the specified bundle type. | 
 **categoryId** | **int32** | Filter restricts the list of results to variations which have the specified category ID. | 
 **anyCategoryId** | **string** | Filter restricts the list of results to variations which have any of the specified category IDs. More than one parameter should be separated by commas. | 
 **allCategoryIds** | **string** | Filter restricts the list of results to variations which have all specified category IDs. More than one parameter should be separated by commas. | 
 **anyCharacteristicId** | **string** | Filter restricts the list of results to variations which have any of the specified characteristic IDs. | 
 **allCharacteristicIds** | **string** | Filter restricts the list of results to variations which have all specified characteristic IDs. | 
 **clientId** | **int32** | Filter restricts the list of results to variations which have the specified client ID. | 
 **anyClientId** | **string** | Filter restricts the list of results to variations which have any of the specified client IDs. More than one parameter should be separated by commas. | 
 **allClientIds** | **string** | Filter restricts the list of results to variations which have all specified client IDs. More than one parameter should be separated by commas. | 
 **automaticClientVisibilities** | **string** | Filter restricts the list of results to variations which have any of the specified automatic client visibilities. More than one parameter should be separated by commas. | 
 **imageHasMarketId** | **float32** |  Filter restricts the list of results to variations which have an image available for the specified market ID. | 
 **flag1** | **int32** | Filter restricts the list of results to variations with the specified flag one. | 
 **flag2** | **int32** | Filter restricts the list of results to variations with the specified flag two. | 
 **manufacturerId** | **int32** | Filter restricts the list of results to variations with the specified manufacturer ID. | 
 **anyManufacturerId** | **string** | Filter restricts the list of results to variations with any of the specified manufacturer IDs. | 
 **itemType** | **string** | Filter restricts the list of results to variations which have the specified item type. | 
 **marketId** | **float32** | Filter restricts the list of results to variations which have the specified market ID. | 
 **anyMarketId** | **string** | Filter restricts the list of results to variations which have any of the specified market IDs. | 
 **allMarketIds** | **string** | Filter restricts the list of results to variations which have all specified market IDs. | 
 **priceBetween** | **string** | Filter restricts the list of results to variations which have a sales price between the specified minimum and maximum value. Minimum and maximum value should be separated by a comma. | 
 **priceBetweenById** | **string** | Filter restricts the list of results to variations where the specified sales price is between the specified minimum and maximum value. Sales price ID, Minimum and maximum value should be separated by a comma. | 
 **anySalesPriceId** | **string** | Filter restricts the list of results to variations which have any of the specified sales price IDs. More than one parameter should be separated by commas. | 
 **propertySelectionId** | **int32** | Filter restricts the list of results to variations which have the specified property selection ID. | 
 **anyPropertySelectionId** | **string** | Filter restricts the list of results to variations which have any of the specified property selection IDs. | 
 **allPropertySelectionIds** | **string** | Filter restricts the list of results to variations which have all specified property selection IDs. | 
 **hasNameInLanguage** | **string** | Filter restricts the list of results to variations which have a name in the specified language. | 
 **createdAt** | **string** | Filter restricts the list of results to variations which have been created in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. | 
 **updatedAt** | **string** | Filter restricts the list of results to variations which have been updated in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. | 
 **itemCreatedAt** | **string** | Filter restricts the list of results to variations whose item has been created in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. | 
 **itemUpdatedAt** | **string** | Filter restricts the list of results to variations whose item has been updated in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. | 
 **availabilityUpdatedAt** | **string** | Filter restricts the list of results to variations whose availablity has been updated in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. | 
 **stockUpdatedAt** | **string** | Filter restricts the list of results to variations whose stock has been updated in the specified time frame. The from and to parameter should be separated by a comma. If there is no to value, the current time is used instead. | 
 **page** | **int32** | The requested page of results. Default value is 1. | 
 **itemsPerPage** | **int32** | The number of results per page. Maximum value is 250. Default value is 50. | 

### Return type

[**InlineResponse200**](InlineResponse200.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json;charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestPimVariationsPut

> RestPimVariationsPut(ctx).PimVariation(pimVariation).Execute()

Create a list of variations and their related data



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
    pimVariation := []openapiclient.PimVariation{*openapiclient.NewPimVariation()} // []PimVariation |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestPimVariationsPut(context.Background()).PimVariation(pimVariation).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestPimVariationsPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestPimVariationsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pimVariation** | [**[]PimVariation**](PimVariation.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestStockmanagementWarehousesWarehouseIdStockCorrectionPut

> RestStockmanagementWarehousesWarehouseIdStockCorrectionPut(ctx, warehouseId).StockCorrections(stockCorrections).Execute()

set stocks



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
    warehouseId := int32(56) // int32 | 
    stockCorrections := *openapiclient.NewStockCorrections() // StockCorrections |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestStockmanagementWarehousesWarehouseIdStockCorrectionPut(context.Background(), warehouseId).StockCorrections(stockCorrections).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestStockmanagementWarehousesWarehouseIdStockCorrectionPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**warehouseId** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestStockmanagementWarehousesWarehouseIdStockCorrectionPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stockCorrections** | [**StockCorrections**](StockCorrections.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json;charset=utf-8
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

