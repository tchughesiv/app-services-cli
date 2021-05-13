/*
 * Service Registry Service - Fleet Manager - v1
 *
 * Main entry point for the system, responsible for all sorts of management operations for the whole service of managed service registry.
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kasclient

import (
	"encoding/json"
)

// Registry Service Registry instance within a multi-tenant deployment.
type Registry struct {
	Id int32 `json:"id"`
	Status RegistryStatus `json:"status"`
	RegistryUrl string `json:"registryUrl"`
	// User-defined Registry name. Does not have to be unique.
	Name *string `json:"name,omitempty"`
	// Identifier of a multi-tenant deployment, where this Service Registry instance resides.
	RegistryDeploymentId *int32 `json:"registryDeploymentId,omitempty"`
}

// NewRegistry instantiates a new Registry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistry(id int32, status RegistryStatus, registryUrl string) *Registry {
	this := Registry{}
	this.Id = id
	this.Status = status
	this.RegistryUrl = registryUrl
	return &this
}

// NewRegistryWithDefaults instantiates a new Registry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistryWithDefaults() *Registry {
	this := Registry{}
	return &this
}

// GetId returns the Id field value
func (o *Registry) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Registry) GetIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Registry) SetId(v int32) {
	o.Id = v
}

// GetStatus returns the Status field value
func (o *Registry) GetStatus() RegistryStatus {
	if o == nil {
		var ret RegistryStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Registry) GetStatusOk() (*RegistryStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Registry) SetStatus(v RegistryStatus) {
	o.Status = v
}

// GetRegistryUrl returns the RegistryUrl field value
func (o *Registry) GetRegistryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RegistryUrl
}

// GetRegistryUrlOk returns a tuple with the RegistryUrl field value
// and a boolean to check if the value has been set.
func (o *Registry) GetRegistryUrlOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RegistryUrl, true
}

// SetRegistryUrl sets field value
func (o *Registry) SetRegistryUrl(v string) {
	o.RegistryUrl = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Registry) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Registry) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Registry) SetName(v string) {
	o.Name = &v
}

// GetRegistryDeploymentId returns the RegistryDeploymentId field value if set, zero value otherwise.
func (o *Registry) GetRegistryDeploymentId() int32 {
	if o == nil || o.RegistryDeploymentId == nil {
		var ret int32
		return ret
	}
	return *o.RegistryDeploymentId
}

// GetRegistryDeploymentIdOk returns a tuple with the RegistryDeploymentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Registry) GetRegistryDeploymentIdOk() (*int32, bool) {
	if o == nil || o.RegistryDeploymentId == nil {
		return nil, false
	}
	return o.RegistryDeploymentId, true
}

// HasRegistryDeploymentId returns a boolean if a field has been set.
func (o *Registry) HasRegistryDeploymentId() bool {
	if o != nil && o.RegistryDeploymentId != nil {
		return true
	}

	return false
}

// SetRegistryDeploymentId gets a reference to the given int32 and assigns it to the RegistryDeploymentId field.
func (o *Registry) SetRegistryDeploymentId(v int32) {
	o.RegistryDeploymentId = &v
}

func (o Registry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["registryUrl"] = o.RegistryUrl
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RegistryDeploymentId != nil {
		toSerialize["registryDeploymentId"] = o.RegistryDeploymentId
	}
	return json.Marshal(toSerialize)
}

type NullableRegistry struct {
	value *Registry
	isSet bool
}

func (v NullableRegistry) Get() *Registry {
	return v.value
}

func (v *NullableRegistry) Set(val *Registry) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistry) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistry(val *Registry) *NullableRegistry {
	return &NullableRegistry{value: val, isSet: true}
}

func (v NullableRegistry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


