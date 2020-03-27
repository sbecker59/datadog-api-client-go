/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// OrgResponse Response with an organization.
type OrgResponse struct {
	Org *Org `json:"org,omitempty"`
}

// NewOrgResponse instantiates a new OrgResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgResponse() *OrgResponse {
	this := OrgResponse{}
	return &this
}

// NewOrgResponseWithDefaults instantiates a new OrgResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgResponseWithDefaults() *OrgResponse {
	this := OrgResponse{}
	return &this
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *OrgResponse) GetOrg() Org {
	if o == nil || o.Org == nil {
		var ret Org
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *OrgResponse) GetOrgOk() (Org, bool) {
	if o == nil || o.Org == nil {
		var ret Org
		return ret, false
	}
	return *o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *OrgResponse) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given Org and assigns it to the Org field.
func (o *OrgResponse) SetOrg(v Org) {
	o.Org = &v
}

func (o OrgResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	return json.Marshal(toSerialize)
}

type NullableOrgResponse struct {
	value *OrgResponse
	isSet bool
}

func (v NullableOrgResponse) Get() *OrgResponse {
	return v.value
}

func (v NullableOrgResponse) Set(val *OrgResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgResponse) IsSet() bool {
	return v.isSet
}

func (v NullableOrgResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgResponse(val *OrgResponse) *NullableOrgResponse {
	return &NullableOrgResponse{value: val, isSet: true}
}

func (v NullableOrgResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
