/*
 * Decision Service Fleet Manager
 *
 * Decision Service Fleet Manager is a Rest API to manage decision instances and connectors.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package decisclient

import (
	"encoding/json"
	"time"
)

// DecisionRequest struct for DecisionRequest
type DecisionRequest struct {
	Id            *string                       `json:"id,omitempty"`
	Kind          *string                       `json:"kind,omitempty"`
	Href          *string                       `json:"href,omitempty"`
	Status        *string                       `json:"status,omitempty"`
	StatusMessage *string                       `json:"status_message,omitempty"`
	Description   *string                       `json:"description,omitempty"`
	Name          *string                       `json:"name,omitempty"`
	Url           *string                       `json:"url,omitempty"`
	Model         *DecisionRequestAllOfModel    `json:"model,omitempty"`
	Eventing      *DecisionRequestAllOfEventing `json:"eventing,omitempty"`
	SubmittedAt   *time.Time                    `json:"submitted_at,omitempty"`
	PublishedAt   *time.Time                    `json:"published_at,omitempty"`
	Configuration *map[string]string            `json:"configuration,omitempty"`
	Tags          *map[string]string            `json:"tags,omitempty"`
	Version       *int                          `json:"version,omitempty"`
}

// NewDecisionRequest instantiates a new DecisionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecisionRequest() *DecisionRequest {
	this := DecisionRequest{}
	return &this
}

// NewDecisionRequestWithDefaults instantiates a new DecisionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecisionRequestWithDefaults() *DecisionRequest {
	this := DecisionRequest{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DecisionRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DecisionRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DecisionRequest) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *DecisionRequest) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *DecisionRequest) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *DecisionRequest) SetKind(v string) {
	o.Kind = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *DecisionRequest) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *DecisionRequest) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *DecisionRequest) SetHref(v string) {
	o.Href = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *DecisionRequest) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *DecisionRequest) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *DecisionRequest) SetStatus(v string) {
	o.Status = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *DecisionRequest) GetStatusMessage() string {
	if o == nil || o.StatusMessage == nil {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetStatusMessageOk() (*string, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *DecisionRequest) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *DecisionRequest) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DecisionRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DecisionRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DecisionRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DecisionRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DecisionRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DecisionRequest) SetName(v string) {
	o.Name = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *DecisionRequest) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *DecisionRequest) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *DecisionRequest) SetUrl(v string) {
	o.Url = &v
}

// GetModel returns the Model field value if set, zero value otherwise.
func (o *DecisionRequest) GetModel() DecisionRequestAllOfModel {
	if o == nil || o.Model == nil {
		var ret DecisionRequestAllOfModel
		return ret
	}
	return *o.Model
}

// GetModelOk returns a tuple with the Model field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetModelOk() (*DecisionRequestAllOfModel, bool) {
	if o == nil || o.Model == nil {
		return nil, false
	}
	return o.Model, true
}

// HasModel returns a boolean if a field has been set.
func (o *DecisionRequest) HasModel() bool {
	if o != nil && o.Model != nil {
		return true
	}

	return false
}

// SetModel gets a reference to the given DecisionRequestAllOfModel and assigns it to the Model field.
func (o *DecisionRequest) SetModel(v DecisionRequestAllOfModel) {
	o.Model = &v
}

// GetEventing returns the Eventing field value if set, zero value otherwise.
func (o *DecisionRequest) GetEventing() DecisionRequestAllOfEventing {
	if o == nil || o.Eventing == nil {
		var ret DecisionRequestAllOfEventing
		return ret
	}
	return *o.Eventing
}

// GetEventingOk returns a tuple with the Eventing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetEventingOk() (*DecisionRequestAllOfEventing, bool) {
	if o == nil || o.Eventing == nil {
		return nil, false
	}
	return o.Eventing, true
}

// HasEventing returns a boolean if a field has been set.
func (o *DecisionRequest) HasEventing() bool {
	if o != nil && o.Eventing != nil {
		return true
	}

	return false
}

// SetEventing gets a reference to the given DecisionRequestAllOfEventing and assigns it to the Eventing field.
func (o *DecisionRequest) SetEventing(v DecisionRequestAllOfEventing) {
	o.Eventing = &v
}

// GetSubmittedAt returns the SubmittedAt field value if set, zero value otherwise.
func (o *DecisionRequest) GetSubmittedAt() time.Time {
	if o == nil || o.SubmittedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.SubmittedAt
}

// GetSubmittedAtOk returns a tuple with the SubmittedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetSubmittedAtOk() (*time.Time, bool) {
	if o == nil || o.SubmittedAt == nil {
		return nil, false
	}
	return o.SubmittedAt, true
}

// HasSubmittedAt returns a boolean if a field has been set.
func (o *DecisionRequest) HasSubmittedAt() bool {
	if o != nil && o.SubmittedAt != nil {
		return true
	}

	return false
}

// SetSubmittedAt gets a reference to the given time.Time and assigns it to the SubmittedAt field.
func (o *DecisionRequest) SetSubmittedAt(v time.Time) {
	o.SubmittedAt = &v
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise.
func (o *DecisionRequest) GetPublishedAt() time.Time {
	if o == nil || o.PublishedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil || o.PublishedAt == nil {
		return nil, false
	}
	return o.PublishedAt, true
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *DecisionRequest) HasPublishedAt() bool {
	if o != nil && o.PublishedAt != nil {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given time.Time and assigns it to the PublishedAt field.
func (o *DecisionRequest) SetPublishedAt(v time.Time) {
	o.PublishedAt = &v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *DecisionRequest) GetConfiguration() map[string]string {
	if o == nil || o.Configuration == nil {
		var ret map[string]string
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetConfigurationOk() (*map[string]string, bool) {
	if o == nil || o.Configuration == nil {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *DecisionRequest) HasConfiguration() bool {
	if o != nil && o.Configuration != nil {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]string and assigns it to the Configuration field.
func (o *DecisionRequest) SetConfiguration(v map[string]string) {
	o.Configuration = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DecisionRequest) GetTags() map[string]string {
	if o == nil || o.Tags == nil {
		var ret map[string]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetTagsOk() (*map[string]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DecisionRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]string and assigns it to the Tags field.
func (o *DecisionRequest) SetTags(v map[string]string) {
	o.Tags = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *DecisionRequest) GetVersion() int {
	if o == nil || o.Version == nil {
		var ret int
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequest) GetVersionOk() (*int, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *DecisionRequest) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int and assigns it to the Version field.
func (o *DecisionRequest) SetVersion(v int) {
	o.Version = &v
}

func (o DecisionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Kind != nil {
		toSerialize["kind"] = o.Kind
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessage != nil {
		toSerialize["status_message"] = o.StatusMessage
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	if o.Model != nil {
		toSerialize["model"] = o.Model
	}
	if o.Eventing != nil {
		toSerialize["eventing"] = o.Eventing
	}
	if o.SubmittedAt != nil {
		toSerialize["submitted_at"] = o.SubmittedAt
	}
	if o.PublishedAt != nil {
		toSerialize["published_at"] = o.PublishedAt
	}
	if o.Configuration != nil {
		toSerialize["configuration"] = o.Configuration
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDecisionRequest struct {
	value *DecisionRequest
	isSet bool
}

func (v NullableDecisionRequest) Get() *DecisionRequest {
	return v.value
}

func (v *NullableDecisionRequest) Set(val *DecisionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionRequest(val *DecisionRequest) *NullableDecisionRequest {
	return &NullableDecisionRequest{value: val, isSet: true}
}

func (v NullableDecisionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}