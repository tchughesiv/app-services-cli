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

// DecisionRequestPayloadAllOfModel struct for DecisionRequestPayloadAllOfModel
type DecisionRequestPayloadAllOfModel struct {
	Dmn string `json:"dmn"`
}

// NewDecisionRequestPayloadAllOfModel instantiates a new DecisionRequestPayloadAllOfModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecisionRequestPayloadAllOfModel(dmn string) *DecisionRequestPayloadAllOfModel {
	this := DecisionRequestPayloadAllOfModel{}
	this.Dmn = dmn
	return &this
}

// NewDecisionRequestPayloadAllOfModelWithDefaults instantiates a new DecisionRequestPayloadAllOfModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecisionRequestPayloadAllOfModelWithDefaults() *DecisionRequestPayloadAllOfModel {
	this := DecisionRequestPayloadAllOfModel{}
	return &this
}

// GetDmn returns the Dmn field value
func (o *DecisionRequestPayloadAllOfModel) GetDmn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dmn
}

// GetDmnOk returns a tuple with the Dmn field value
// and a boolean to check if the value has been set.
func (o *DecisionRequestPayloadAllOfModel) GetDmnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dmn, true
}

// SetDmn sets field value
func (o *DecisionRequestPayloadAllOfModel) SetDmn(v string) {
	o.Dmn = v
}

func (o DecisionRequestPayloadAllOfModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["dmn"] = o.Dmn
	}
	return json.Marshal(toSerialize)
}

type NullableDecisionRequestPayloadAllOfModel struct {
	value *DecisionRequestPayloadAllOfModel
	isSet bool
}

func (v NullableDecisionRequestPayloadAllOfModel) Get() *DecisionRequestPayloadAllOfModel {
	return v.value
}

func (v *NullableDecisionRequestPayloadAllOfModel) Set(val *DecisionRequestPayloadAllOfModel) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionRequestPayloadAllOfModel) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionRequestPayloadAllOfModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionRequestPayloadAllOfModel(val *DecisionRequestPayloadAllOfModel) *NullableDecisionRequestPayloadAllOfModel {
	return &NullableDecisionRequestPayloadAllOfModel{value: val, isSet: true}
}

func (v NullableDecisionRequestPayloadAllOfModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionRequestPayloadAllOfModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
