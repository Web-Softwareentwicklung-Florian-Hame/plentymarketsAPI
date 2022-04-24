# \DefaultApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RestCategoriesGet**](DefaultApi.md#RestCategoriesGet) | **Get** /rest/categories | fetching categories with pagination and optional filter queries
[**RestCategoriesPost**](DefaultApi.md#RestCategoriesPost) | **Post** /rest/categories | creating new categories
[**RestItemsAttributeValuesValueIdNamesPost**](DefaultApi.md#RestItemsAttributeValuesValueIdNamesPost) | **Post** /rest/items/attribute_values/{valueId}/names | creating attribute values names
[**RestItemsAttributesAttributeIdNamesPost**](DefaultApi.md#RestItemsAttributesAttributeIdNamesPost) | **Post** /rest/items/attributes/{attributeId}/names | creating names for attribute
[**RestItemsAttributesAttributeIdValuesGet**](DefaultApi.md#RestItemsAttributesAttributeIdValuesGet) | **Get** /rest/items/attributes/{attributeId}/values | fetching attributes values with pagination
[**RestItemsAttributesAttributeIdValuesPost**](DefaultApi.md#RestItemsAttributesAttributeIdValuesPost) | **Post** /rest/items/attributes/{attributeId}/values | creating new attribute values
[**RestItemsAttributesGet**](DefaultApi.md#RestItemsAttributesGet) | **Get** /rest/items/attributes | fetching attributes with pagination
[**RestItemsAttributesPost**](DefaultApi.md#RestItemsAttributesPost) | **Post** /rest/items/attributes | creating new attributes
[**RestItemsGet**](DefaultApi.md#RestItemsGet) | **Get** /rest/items | fetching items with optional extra data or filters given in query parameter
[**RestItemsItemIdVariationsPost**](DefaultApi.md#RestItemsItemIdVariationsPost) | **Post** /rest/items/{itemId}/variations | creating variations for an item
[**RestItemsPost**](DefaultApi.md#RestItemsPost) | **Post** /rest/items | creates items
[**RestItemsVariationsGet**](DefaultApi.md#RestItemsVariationsGet) | **Get** /rest/items/variations | fetching variations with pagination and the possibility to fetch additional related data by &#39;with&#39; query param
[**RestLoginPost**](DefaultApi.md#RestLoginPost) | **Post** /rest/login | login for authentication at rest api



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
    variation := []openapiclient.Variation{*openapiclient.NewVariation()} // []Variation |  (optional)

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

 **variation** | [**[]Variation**](Variation.md) |  | 

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


## RestItemsPost

> []Item RestItemsPost(ctx).Item(item).Execute()

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
    item := []openapiclient.Item{*openapiclient.NewItem(int32(123), int32(123), int32(123), []openapiclient.Variation{*openapiclient.NewVariation()})} // []Item |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DefaultApi.RestItemsPost(context.Background()).Item(item).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.RestItemsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestItemsPost`: []Item
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.RestItemsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestItemsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **item** | [**[]Item**](Item.md) |  | 

### Return type

[**[]Item**](Item.md)

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

