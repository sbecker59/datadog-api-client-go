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

// SyntheticsTestMonitorStatus the model 'SyntheticsTestMonitorStatus'
type SyntheticsTestMonitorStatus int64

// List of SyntheticsTestMonitorStatus
const (
	SYNTHETICSTESTMONITORSTATUS_UNTRIGGERED SyntheticsTestMonitorStatus = 0
	SYNTHETICSTESTMONITORSTATUS_TRIGGERED   SyntheticsTestMonitorStatus = 1
	SYNTHETICSTESTMONITORSTATUS_NO_DATA     SyntheticsTestMonitorStatus = 2
)

// Ptr returns reference to SyntheticsTestMonitorStatus value
func (v SyntheticsTestMonitorStatus) Ptr() *SyntheticsTestMonitorStatus {
	return &v
}

type NullableSyntheticsTestMonitorStatus struct {
	value *SyntheticsTestMonitorStatus
	isSet bool
}

func (v NullableSyntheticsTestMonitorStatus) Get() *SyntheticsTestMonitorStatus {
	return v.value
}

func (v NullableSyntheticsTestMonitorStatus) Set(val *SyntheticsTestMonitorStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSyntheticsTestMonitorStatus) IsSet() bool {
	return v.isSet
}

func (v NullableSyntheticsTestMonitorStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSyntheticsTestMonitorStatus(val *SyntheticsTestMonitorStatus) *NullableSyntheticsTestMonitorStatus {
	return &NullableSyntheticsTestMonitorStatus{value: val, isSet: true}
}

func (v NullableSyntheticsTestMonitorStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSyntheticsTestMonitorStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}