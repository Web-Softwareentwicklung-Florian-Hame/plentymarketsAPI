# Property

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PropertyId** | Pointer to **int32** |  | [optional] 
**RelationTypeIdentifier** | Pointer to **string** |  | [optional] 
**RelationTargetId** | Pointer to **int32** |  | [optional] 
**SelectionRelationId** | Pointer to **int32** |  | [optional] 
**GroupId** | Pointer to **int32** |  | [optional] 
**Markup** | Pointer to **int32** |  | [optional] 
**RelationValues** | Pointer to [**[]PropertyRelationValues**](PropertyRelationValues.md) |  | [optional] 
**PropertyRelation** | Pointer to [**PropertyRelation**](PropertyRelation.md) |  | [optional] 
**Id** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewProperty

`func NewProperty() *Property`

NewProperty instantiates a new Property object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyWithDefaults

`func NewPropertyWithDefaults() *Property`

NewPropertyWithDefaults instantiates a new Property object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPropertyId

`func (o *Property) GetPropertyId() int32`

GetPropertyId returns the PropertyId field if non-nil, zero value otherwise.

### GetPropertyIdOk

`func (o *Property) GetPropertyIdOk() (*int32, bool)`

GetPropertyIdOk returns a tuple with the PropertyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyId

`func (o *Property) SetPropertyId(v int32)`

SetPropertyId sets PropertyId field to given value.

### HasPropertyId

`func (o *Property) HasPropertyId() bool`

HasPropertyId returns a boolean if a field has been set.

### GetRelationTypeIdentifier

`func (o *Property) GetRelationTypeIdentifier() string`

GetRelationTypeIdentifier returns the RelationTypeIdentifier field if non-nil, zero value otherwise.

### GetRelationTypeIdentifierOk

`func (o *Property) GetRelationTypeIdentifierOk() (*string, bool)`

GetRelationTypeIdentifierOk returns a tuple with the RelationTypeIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationTypeIdentifier

`func (o *Property) SetRelationTypeIdentifier(v string)`

SetRelationTypeIdentifier sets RelationTypeIdentifier field to given value.

### HasRelationTypeIdentifier

`func (o *Property) HasRelationTypeIdentifier() bool`

HasRelationTypeIdentifier returns a boolean if a field has been set.

### GetRelationTargetId

`func (o *Property) GetRelationTargetId() int32`

GetRelationTargetId returns the RelationTargetId field if non-nil, zero value otherwise.

### GetRelationTargetIdOk

`func (o *Property) GetRelationTargetIdOk() (*int32, bool)`

GetRelationTargetIdOk returns a tuple with the RelationTargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationTargetId

`func (o *Property) SetRelationTargetId(v int32)`

SetRelationTargetId sets RelationTargetId field to given value.

### HasRelationTargetId

`func (o *Property) HasRelationTargetId() bool`

HasRelationTargetId returns a boolean if a field has been set.

### GetSelectionRelationId

`func (o *Property) GetSelectionRelationId() int32`

GetSelectionRelationId returns the SelectionRelationId field if non-nil, zero value otherwise.

### GetSelectionRelationIdOk

`func (o *Property) GetSelectionRelationIdOk() (*int32, bool)`

GetSelectionRelationIdOk returns a tuple with the SelectionRelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionRelationId

`func (o *Property) SetSelectionRelationId(v int32)`

SetSelectionRelationId sets SelectionRelationId field to given value.

### HasSelectionRelationId

`func (o *Property) HasSelectionRelationId() bool`

HasSelectionRelationId returns a boolean if a field has been set.

### GetGroupId

`func (o *Property) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Property) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Property) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Property) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetMarkup

`func (o *Property) GetMarkup() int32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *Property) GetMarkupOk() (*int32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *Property) SetMarkup(v int32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *Property) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetRelationValues

`func (o *Property) GetRelationValues() []PropertyRelationValues`

GetRelationValues returns the RelationValues field if non-nil, zero value otherwise.

### GetRelationValuesOk

`func (o *Property) GetRelationValuesOk() (*[]PropertyRelationValues, bool)`

GetRelationValuesOk returns a tuple with the RelationValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationValues

`func (o *Property) SetRelationValues(v []PropertyRelationValues)`

SetRelationValues sets RelationValues field to given value.

### HasRelationValues

`func (o *Property) HasRelationValues() bool`

HasRelationValues returns a boolean if a field has been set.

### GetPropertyRelation

`func (o *Property) GetPropertyRelation() PropertyRelation`

GetPropertyRelation returns the PropertyRelation field if non-nil, zero value otherwise.

### GetPropertyRelationOk

`func (o *Property) GetPropertyRelationOk() (*PropertyRelation, bool)`

GetPropertyRelationOk returns a tuple with the PropertyRelation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertyRelation

`func (o *Property) SetPropertyRelation(v PropertyRelation)`

SetPropertyRelation sets PropertyRelation field to given value.

### HasPropertyRelation

`func (o *Property) HasPropertyRelation() bool`

HasPropertyRelation returns a boolean if a field has been set.

### GetId

`func (o *Property) GetId() float32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Property) GetIdOk() (*float32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Property) SetId(v float32)`

SetId sets Id field to given value.

### HasId

`func (o *Property) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Property) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Property) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Property) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Property) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Property) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Property) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Property) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Property) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


