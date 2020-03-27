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

// UsageTopAvgMetricsResponse Response containing the number of hourly recorded custom metrics for a given organization.
type UsageTopAvgMetricsResponse struct {
	// Number of hourly recorded custom metrics for a given organization.
	Usage *[]UsageTopAvgMetricsHour `json:"usage,omitempty"`
}

// NewUsageTopAvgMetricsResponse instantiates a new UsageTopAvgMetricsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageTopAvgMetricsResponse() *UsageTopAvgMetricsResponse {
	this := UsageTopAvgMetricsResponse{}
	return &this
}

// NewUsageTopAvgMetricsResponseWithDefaults instantiates a new UsageTopAvgMetricsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageTopAvgMetricsResponseWithDefaults() *UsageTopAvgMetricsResponse {
	this := UsageTopAvgMetricsResponse{}
	return &this
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *UsageTopAvgMetricsResponse) GetUsage() []UsageTopAvgMetricsHour {
	if o == nil || o.Usage == nil {
		var ret []UsageTopAvgMetricsHour
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *UsageTopAvgMetricsResponse) GetUsageOk() ([]UsageTopAvgMetricsHour, bool) {
	if o == nil || o.Usage == nil {
		var ret []UsageTopAvgMetricsHour
		return ret, false
	}
	return *o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *UsageTopAvgMetricsResponse) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given []UsageTopAvgMetricsHour and assigns it to the Usage field.
func (o *UsageTopAvgMetricsResponse) SetUsage(v []UsageTopAvgMetricsHour) {
	o.Usage = &v
}

func (o UsageTopAvgMetricsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Usage != nil {
		toSerialize["usage"] = o.Usage
	}
	return json.Marshal(toSerialize)
}

type NullableUsageTopAvgMetricsResponse struct {
	value *UsageTopAvgMetricsResponse
	isSet bool
}

func (v NullableUsageTopAvgMetricsResponse) Get() *UsageTopAvgMetricsResponse {
	return v.value
}

func (v NullableUsageTopAvgMetricsResponse) Set(val *UsageTopAvgMetricsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageTopAvgMetricsResponse) IsSet() bool {
	return v.isSet
}

func (v NullableUsageTopAvgMetricsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageTopAvgMetricsResponse(val *UsageTopAvgMetricsResponse) *NullableUsageTopAvgMetricsResponse {
	return &NullableUsageTopAvgMetricsResponse{value: val, isSet: true}
}

func (v NullableUsageTopAvgMetricsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageTopAvgMetricsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
