# DecisionList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**Page** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Items** | [**[]DecisionRequest**](DecisionRequest.md) |  | 

## Methods

### NewDecisionList

`func NewDecisionList(kind string, items []DecisionRequest, ) *DecisionList`

NewDecisionList instantiates a new DecisionList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionListWithDefaults

`func NewDecisionListWithDefaults() *DecisionList`

NewDecisionListWithDefaults instantiates a new DecisionList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *DecisionList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DecisionList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DecisionList) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPage

`func (o *DecisionList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DecisionList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DecisionList) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *DecisionList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *DecisionList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DecisionList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DecisionList) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *DecisionList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *DecisionList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *DecisionList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *DecisionList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *DecisionList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *DecisionList) GetItems() []DecisionRequest`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *DecisionList) GetItemsOk() (*[]DecisionRequest, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *DecisionList) SetItems(v []DecisionRequest)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


