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
)

// DecisionRequestAllOfModel struct for DecisionRequestAllOfModel
type DecisionRequestAllOfModel struct {
	Md5  *string `json:"md5,omitempty"`
	Href *string `json:"href,omitempty"`
}

// NewDecisionRequestAllOfModel instantiates a new DecisionRequestAllOfModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecisionRequestAllOfModel() *DecisionRequestAllOfModel {
	this := DecisionRequestAllOfModel{}
	return &this
}

// NewDecisionRequestAllOfModelWithDefaults instantiates a new DecisionRequestAllOfModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecisionRequestAllOfModelWithDefaults() *DecisionRequestAllOfModel {
	this := DecisionRequestAllOfModel{}
	return &this
}

// GetMd5 returns the Md5 field value if set, zero value otherwise.
func (o *DecisionRequestAllOfModel) GetMd5() string {
	if o == nil || o.Md5 == nil {
		var ret string
		return ret
	}
	return *o.Md5
}

// GetMd5Ok returns a tuple with the Md5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOfModel) GetMd5Ok() (*string, bool) {
	if o == nil || o.Md5 == nil {
		return nil, false
	}
	return o.Md5, true
}

// HasMd5 returns a boolean if a field has been set.
func (o *DecisionRequestAllOfModel) HasMd5() bool {
	if o != nil && o.Md5 != nil {
		return true
	}

	return false
}

// SetMd5 gets a reference to the given string and assigns it to the Md5 field.
func (o *DecisionRequestAllOfModel) SetMd5(v string) {
	o.Md5 = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *DecisionRequestAllOfModel) GetHref() string {
	if o == nil || o.Href == nil {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionRequestAllOfModel) GetHrefOk() (*string, bool) {
	if o == nil || o.Href == nil {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *DecisionRequestAllOfModel) HasHref() bool {
	if o != nil && o.Href != nil {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *DecisionRequestAllOfModel) SetHref(v string) {
	o.Href = &v
}

func (o DecisionRequestAllOfModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Md5 != nil {
		toSerialize["md5"] = o.Md5
	}
	if o.Href != nil {
		toSerialize["href"] = o.Href
	}
	return json.Marshal(toSerialize)
}

type NullableDecisionRequestAllOfModel struct {
	value *DecisionRequestAllOfModel
	isSet bool
}

func (v NullableDecisionRequestAllOfModel) Get() *DecisionRequestAllOfModel {
	return v.value
}

func (v *NullableDecisionRequestAllOfModel) Set(val *DecisionRequestAllOfModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionRequestAllOfModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionRequestAllOfModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionRequestAllOfModel(val *DecisionRequestAllOfModel) *NullableDecisionRequestAllOfModel {
	return &NullableDecisionRequestAllOfModel{value: val, isSet: true}
}

func (v NullableDecisionRequestAllOfModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionRequestAllOfModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
