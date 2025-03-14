// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// CheckCanDeleteMonitorResponseData Wrapper object with the list of monitor IDs.
type CheckCanDeleteMonitorResponseData struct {
	// An array of Monitor IDs that can be safely deleted.
	Ok []int64 `json:"ok,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewCheckCanDeleteMonitorResponseData instantiates a new CheckCanDeleteMonitorResponseData object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewCheckCanDeleteMonitorResponseData() *CheckCanDeleteMonitorResponseData {
	this := CheckCanDeleteMonitorResponseData{}
	return &this
}

// NewCheckCanDeleteMonitorResponseDataWithDefaults instantiates a new CheckCanDeleteMonitorResponseData object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewCheckCanDeleteMonitorResponseDataWithDefaults() *CheckCanDeleteMonitorResponseData {
	this := CheckCanDeleteMonitorResponseData{}
	return &this
}

// GetOk returns the Ok field value if set, zero value otherwise.
func (o *CheckCanDeleteMonitorResponseData) GetOk() []int64 {
	if o == nil || o.Ok == nil {
		var ret []int64
		return ret
	}
	return o.Ok
}

// GetOkOk returns a tuple with the Ok field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckCanDeleteMonitorResponseData) GetOkOk() (*[]int64, bool) {
	if o == nil || o.Ok == nil {
		return nil, false
	}
	return &o.Ok, true
}

// HasOk returns a boolean if a field has been set.
func (o *CheckCanDeleteMonitorResponseData) HasOk() bool {
	return o != nil && o.Ok != nil
}

// SetOk gets a reference to the given []int64 and assigns it to the Ok field.
func (o *CheckCanDeleteMonitorResponseData) SetOk(v []int64) {
	o.Ok = v
}

// MarshalJSON serializes the struct using spec logic.
func (o CheckCanDeleteMonitorResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Ok != nil {
		toSerialize["ok"] = o.Ok
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *CheckCanDeleteMonitorResponseData) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Ok []int64 `json:"ok,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"ok"})
	} else {
		return err
	}
	o.Ok = all.Ok

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	return nil
}
