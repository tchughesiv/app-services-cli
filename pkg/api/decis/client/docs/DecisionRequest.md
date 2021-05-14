# DecisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
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
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewDecisionRequest

`func NewDecisionRequest() *DecisionRequest`

NewDecisionRequest instantiates a new DecisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionRequestWithDefaults

`func NewDecisionRequestWithDefaults() *DecisionRequest`

NewDecisionRequestWithDefaults instantiates a new DecisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DecisionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DecisionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DecisionRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DecisionRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *DecisionRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *DecisionRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *DecisionRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *DecisionRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetHref

`func (o *DecisionRequest) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *DecisionRequest) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *DecisionRequest) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *DecisionRequest) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetStatus

`func (o *DecisionRequest) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DecisionRequest) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DecisionRequest) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DecisionRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *DecisionRequest) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *DecisionRequest) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *DecisionRequest) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *DecisionRequest) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetDescription

`func (o *DecisionRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DecisionRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DecisionRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DecisionRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *DecisionRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DecisionRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DecisionRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DecisionRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *DecisionRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *DecisionRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *DecisionRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *DecisionRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetModel

`func (o *DecisionRequest) GetModel() DecisionRequestAllOfModel`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DecisionRequest) GetModelOk() (*DecisionRequestAllOfModel, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DecisionRequest) SetModel(v DecisionRequestAllOfModel)`

SetModel sets Model field to given value.

### HasModel

`func (o *DecisionRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetEventing

`func (o *DecisionRequest) GetEventing() DecisionRequestAllOfEventing`

GetEventing returns the Eventing field if non-nil, zero value otherwise.

### GetEventingOk

`func (o *DecisionRequest) GetEventingOk() (*DecisionRequestAllOfEventing, bool)`

GetEventingOk returns a tuple with the Eventing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventing

`func (o *DecisionRequest) SetEventing(v DecisionRequestAllOfEventing)`

SetEventing sets Eventing field to given value.

### HasEventing

`func (o *DecisionRequest) HasEventing() bool`

HasEventing returns a boolean if a field has been set.

### GetSubmittedAt

`func (o *DecisionRequest) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *DecisionRequest) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *DecisionRequest) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *DecisionRequest) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### GetPublishedAt

`func (o *DecisionRequest) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *DecisionRequest) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *DecisionRequest) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *DecisionRequest) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetConfiguration

`func (o *DecisionRequest) GetConfiguration() map[string]string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DecisionRequest) GetConfigurationOk() (*map[string]string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DecisionRequest) SetConfiguration(v map[string]string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DecisionRequest) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetTags

`func (o *DecisionRequest) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DecisionRequest) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DecisionRequest) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DecisionRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetVersion

`func (o *DecisionRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DecisionRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DecisionRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DecisionRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


