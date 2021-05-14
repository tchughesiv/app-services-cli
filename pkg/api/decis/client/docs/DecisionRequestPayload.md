# DecisionRequestPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the Decision cluster. It must consist of lower-case alphanumeric characters or &#39;-&#39;, start with an alphabetic character, and end with an alphanumeric character, and can not be longer than 32 characters. | 

## Methods

### NewDecisionRequestPayload

`func NewDecisionRequestPayload(name string, ) *DecisionRequestPayload`

NewDecisionRequestPayload instantiates a new DecisionRequestPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionRequestPayloadWithDefaults

`func NewDecisionRequestPayloadWithDefaults() *DecisionRequestPayload`

NewDecisionRequestPayloadWithDefaults instantiates a new DecisionRequestPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


