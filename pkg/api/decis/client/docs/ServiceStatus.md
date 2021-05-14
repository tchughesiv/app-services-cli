# ServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decisions** | Pointer to [**ServiceStatusDecisions**](ServiceStatusDecisions.md) |  | [optional] 

## Methods

### NewServiceStatus

`func NewServiceStatus() *ServiceStatus`

NewServiceStatus instantiates a new ServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceStatusWithDefaults

`func NewServiceStatusWithDefaults() *ServiceStatus`

NewServiceStatusWithDefaults instantiates a new ServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecisions

`func (o *ServiceStatus) GetDecisions() ServiceStatusDecisions`

GetDecisions returns the Decisions field if non-nil, zero value otherwise.

### GetDecisionsOk

`func (o *ServiceStatus) GetDecisionsOk() (*ServiceStatusDecisions, bool)`

GetDecisionsOk returns a tuple with the Decisions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisions

`func (o *ServiceStatus) SetDecisions(v ServiceStatusDecisions)`

SetDecisions sets Decisions field to given value.

### HasDecisions

`func (o *ServiceStatus) HasDecisions() bool`

HasDecisions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


