# DecisionRequestPayloadAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Decision cluster. It must consist of lower-case alphanumeric characters or &#39;-&#39;, start with an alphabetic character, and end with an alphanumeric character, and can not be longer than 32 characters. | 
**Description** | **string** |  | 
**Model** | [**DecisionRequestPayloadAllOfModel**](DecisionRequestPayloadAllOfModel.md) |  | 
**Eventing** | Pointer to [**DecisionRequestAllOfEventing**](DecisionRequestAllOfEventing.md) |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewDecisionRequestPayloadAllOf

`func NewDecisionRequestPayloadAllOf(name string, description string, model DecisionRequestPayloadAllOfModel, ) *DecisionRequestPayloadAllOf`

NewDecisionRequestPayloadAllOf instantiates a new DecisionRequestPayloadAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionRequestPayloadAllOfWithDefaults

`func NewDecisionRequestPayloadAllOfWithDefaults() *DecisionRequestPayloadAllOf`

NewDecisionRequestPayloadAllOfWithDefaults instantiates a new DecisionRequestPayloadAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DecisionRequestPayloadAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DecisionRequestPayloadAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DecisionRequestPayloadAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *DecisionRequestPayloadAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DecisionRequestPayloadAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DecisionRequestPayloadAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetModel

`func (o *DecisionRequestPayloadAllOf) GetModel() DecisionRequestPayloadAllOfModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DecisionRequestPayloadAllOf) GetModelOk() (*DecisionRequestPayloadAllOfModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DecisionRequestPayloadAllOf) SetModel(v DecisionRequestPayloadAllOfModel)`

SetModel sets Model field to given value.


### GetEventing

`func (o *DecisionRequestPayloadAllOf) GetEventing() DecisionRequestAllOfEventing`

GetEventing returns the Eventing field if non-nil, zero value otherwise.

### GetEventingOk

`func (o *DecisionRequestPayloadAllOf) GetEventingOk() (*DecisionRequestAllOfEventing, bool)`

GetEventingOk returns a tuple with the Eventing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventing

`func (o *DecisionRequestPayloadAllOf) SetEventing(v DecisionRequestAllOfEventing)`

SetEventing sets Eventing field to given value.

### HasEventing

`func (o *DecisionRequestPayloadAllOf) HasEventing() bool`

HasEventing returns a boolean if a field has been set.

### GetConfiguration

`func (o *DecisionRequestPayloadAllOf) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DecisionRequestPayloadAllOf) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DecisionRequestPayloadAllOf) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DecisionRequestPayloadAllOf) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTags

`func (o *DecisionRequestPayloadAllOf) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DecisionRequestPayloadAllOf) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DecisionRequestPayloadAllOf) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DecisionRequestPayloadAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


