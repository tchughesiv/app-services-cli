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

// DecisionList struct for DecisionList
type DecisionList struct {
	Kind  string            `json:"kind"`
	Page  *int32            `json:"page,omitempty"`
	Size  *int32            `json:"size,omitempty"`
	Total *int32            `json:"total,omitempty"`
	Items []DecisionRequest `json:"items"`
}

// NewDecisionList instantiates a new DecisionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDecisionList(kind string, items []DecisionRequest) *DecisionList {
	this := DecisionList{}
	this.Kind = kind
	this.Items = items
	return &this
}

// NewDecisionListWithDefaults instantiates a new DecisionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDecisionListWithDefaults() *DecisionList {
	this := DecisionList{}
	return &this
}

// GetKind returns the Kind field value
func (o *DecisionList) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DecisionList) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *DecisionList) SetKind(v string) {
	o.Kind = v
}

// GetPage returns the Page field value if set, zero value otherwise.
func (o *DecisionList) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionList) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *DecisionList) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *DecisionList) SetPage(v int32) {
	o.Page = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *DecisionList) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionList) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *DecisionList) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *DecisionList) SetSize(v int32) {
	o.Size = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *DecisionList) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DecisionList) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *DecisionList) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *DecisionList) SetTotal(v int32) {
	o.Total = &v
}

// GetItems returns the Items field value
func (o *DecisionList) GetItems() []DecisionRequest {
	if o == nil {
		var ret []DecisionRequest
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *DecisionList) GetItemsOk() (*[]DecisionRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *DecisionList) SetItems(v []DecisionRequest) {
	o.Items = v
}

func (o DecisionList) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableDecisionList struct {
	value *DecisionList
	isSet bool
}

func (v NullableDecisionList) Get() *DecisionList {
	return v.value
}

func (v *NullableDecisionList) Set(val *DecisionList) {
	v.value = val
	v.isSet = true
}

func (v NullableDecisionList) IsSet() bool {
	return v.isSet
}

func (v *NullableDecisionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDecisionList(val *DecisionList) *NullableDecisionList {
	return &NullableDecisionList{value: val, isSet: true}
}

func (v NullableDecisionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDecisionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}