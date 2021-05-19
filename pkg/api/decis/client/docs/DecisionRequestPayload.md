# DecisionRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**Name** | **string** | The name of the Decision cluster. It must consist of lower-case alphanumeric characters or &#39;-&#39;, start with an alphabetic character, and end with an alphanumeric character, and can not be longer than 32 characters. | 
**Description** | **string** |  | 
**Model** | [**DecisionRequestPayloadAllOfModel**](DecisionRequestPayloadAllOfModel.md) |  | 
**Eventing** | Pointer to [**DecisionRequestAllOfEventing**](DecisionRequestAllOfEventing.md) |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDecisionRequestPayload

`func NewDecisionRequestPayload(name string, description string, model DecisionRequestPayloadAllOfModel, ) *DecisionRequestPayload`

NewDecisionRequestPayload instantiates a new DecisionRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionRequestPayloadWithDefaults

`func NewDecisionRequestPayloadWithDefaults() *DecisionRequestPayload`

NewDecisionRequestPayloadWithDefaults instantiates a new DecisionRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DecisionRequestPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DecisionRequestPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DecisionRequestPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DecisionRequestPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *DecisionRequestPayload) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DecisionRequestPayload) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DecisionRequestPayload) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DecisionRequestPayload) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *DecisionRequestPayload) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DecisionRequestPayload) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DecisionRequestPayload) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DecisionRequestPayload) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetName

`func (o *DecisionRequestPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DecisionRequestPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DecisionRequestPayload) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DecisionRequestPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DecisionRequestPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DecisionRequestPayload) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModel

`func (o *DecisionRequestPayload) GetModel() DecisionRequestPayloadAllOfModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DecisionRequestPayload) GetModelOk() (*DecisionRequestPayloadAllOfModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DecisionRequestPayload) SetModel(v DecisionRequestPayloadAllOfModel)`

SetModel sets Model field to given value.


### GetEventing

`func (o *DecisionRequestPayload) GetEventing() DecisionRequestAllOfEventing`

GetEventing returns the Eventing field if non-nil, zero value otherwise.

### GetEventingOk

`func (o *DecisionRequestPayload) GetEventingOk() (*DecisionRequestAllOfEventing, bool)`

GetEventingOk returns a tuple with the Eventing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventing

`func (o *DecisionRequestPayload) SetEventing(v DecisionRequestAllOfEventing)`

SetEventing sets Eventing field to given value.

### HasEventing

`func (o *DecisionRequestPayload) HasEventing() bool`

HasEventing returns a boolean if a field has been set.

### GetConfiguration

`func (o *DecisionRequestPayload) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DecisionRequestPayload) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DecisionRequestPayload) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DecisionRequestPayload) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTags

`func (o *DecisionRequestPayload) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DecisionRequestPayload) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DecisionRequestPayload) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DecisionRequestPayload) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


