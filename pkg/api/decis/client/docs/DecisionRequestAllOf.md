# DecisionRequestAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Model** | Pointer to [**DecisionRequestAllOfModel**](DecisionRequestAllOfModel.md) |  | [optional] 
**Eventing** | Pointer to [**DecisionRequestAllOfEventing**](DecisionRequestAllOfEventing.md) |  | [optional] 
**SubmittedAt** | Pointer to **time.Time** |  | [optional] 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**Configuration** | Pointer to **map[string]string** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Version** | Pointer to **int** |  | [optional] 

## Methods

### NewDecisionRequestAllOf

`func NewDecisionRequestAllOf() *DecisionRequestAllOf`

NewDecisionRequestAllOf instantiates a new DecisionRequestAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionRequestAllOfWithDefaults

`func NewDecisionRequestAllOfWithDefaults() *DecisionRequestAllOf`

NewDecisionRequestAllOfWithDefaults instantiates a new DecisionRequestAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DecisionRequestAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DecisionRequestAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DecisionRequestAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DecisionRequestAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *DecisionRequestAllOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DecisionRequestAllOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DecisionRequestAllOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *DecisionRequestAllOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetDescription

`func (o *DecisionRequestAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DecisionRequestAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DecisionRequestAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DecisionRequestAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DecisionRequestAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DecisionRequestAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DecisionRequestAllOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DecisionRequestAllOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *DecisionRequestAllOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DecisionRequestAllOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DecisionRequestAllOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DecisionRequestAllOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetModel

`func (o *DecisionRequestAllOf) GetModel() DecisionRequestAllOfModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DecisionRequestAllOf) GetModelOk() (*DecisionRequestAllOfModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DecisionRequestAllOf) SetModel(v DecisionRequestAllOfModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *DecisionRequestAllOf) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetEventing

`func (o *DecisionRequestAllOf) GetEventing() DecisionRequestAllOfEventing`

GetEventing returns the Eventing field if non-nil, zero value otherwise.

### GetEventingOk

`func (o *DecisionRequestAllOf) GetEventingOk() (*DecisionRequestAllOfEventing, bool)`

GetEventingOk returns a tuple with the Eventing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventing

`func (o *DecisionRequestAllOf) SetEventing(v DecisionRequestAllOfEventing)`

SetEventing sets Eventing field to given value.

### HasEventing

`func (o *DecisionRequestAllOf) HasEventing() bool`

HasEventing returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *DecisionRequestAllOf) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *DecisionRequestAllOf) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *DecisionRequestAllOf) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *DecisionRequestAllOf) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *DecisionRequestAllOf) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *DecisionRequestAllOf) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *DecisionRequestAllOf) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *DecisionRequestAllOf) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetConfiguration

`func (o *DecisionRequestAllOf) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DecisionRequestAllOf) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DecisionRequestAllOf) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DecisionRequestAllOf) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTags

`func (o *DecisionRequestAllOf) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DecisionRequestAllOf) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DecisionRequestAllOf) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DecisionRequestAllOf) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *DecisionRequestAllOf) GetVersion() int`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecisionRequestAllOf) GetVersionOk() (*int, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecisionRequestAllOf) SetVersion(v int)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DecisionRequestAllOf) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


