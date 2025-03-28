// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV2

import (
	"fmt"

	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// ApplicationSecurityWafCustomRuleAttributes A WAF custom rule.
type ApplicationSecurityWafCustomRuleAttributes struct {
	// The definition of `ApplicationSecurityWafCustomRuleAction` object.
	Action *ApplicationSecurityWafCustomRuleAction `json:"action,omitempty"`
	// Indicates whether the WAF custom rule will block the request.
	Blocking bool `json:"blocking"`
	// Conditions for which the WAF Custom Rule will triggers, all conditions needs to match in order for the WAF
	// rule to trigger.
	Conditions []ApplicationSecurityWafCustomRuleCondition `json:"conditions"`
	// Indicates whether the WAF custom rule is enabled.
	Enabled bool `json:"enabled"`
	// Metadata associated with the WAF Custom Rule.
	Metadata *ApplicationSecurityWafCustomRuleMetadata `json:"metadata,omitempty"`
	// The Name of the WAF custom rule.
	Name string `json:"name"`
	// The path glob for the WAF custom rule.
	PathGlob *string `json:"path_glob,omitempty"`
	// The scope of the WAF custom rule.
	Scope []ApplicationSecurityWafCustomRuleScope `json:"scope,omitempty"`
	// Tags associated with the WAF Custom Rule. The concatenation of category and type will form the security
	// activity field associated with the traces.
	Tags ApplicationSecurityWafCustomRuleTags `json:"tags"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewApplicationSecurityWafCustomRuleAttributes instantiates a new ApplicationSecurityWafCustomRuleAttributes object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewApplicationSecurityWafCustomRuleAttributes(blocking bool, conditions []ApplicationSecurityWafCustomRuleCondition, enabled bool, name string, tags ApplicationSecurityWafCustomRuleTags) *ApplicationSecurityWafCustomRuleAttributes {
	this := ApplicationSecurityWafCustomRuleAttributes{}
	this.Blocking = blocking
	this.Conditions = conditions
	this.Enabled = enabled
	this.Name = name
	this.Tags = tags
	return &this
}

// NewApplicationSecurityWafCustomRuleAttributesWithDefaults instantiates a new ApplicationSecurityWafCustomRuleAttributes object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewApplicationSecurityWafCustomRuleAttributesWithDefaults() *ApplicationSecurityWafCustomRuleAttributes {
	this := ApplicationSecurityWafCustomRuleAttributes{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetAction() ApplicationSecurityWafCustomRuleAction {
	if o == nil || o.Action == nil {
		var ret ApplicationSecurityWafCustomRuleAction
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetActionOk() (*ApplicationSecurityWafCustomRuleAction, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) HasAction() bool {
	return o != nil && o.Action != nil
}

// SetAction gets a reference to the given ApplicationSecurityWafCustomRuleAction and assigns it to the Action field.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetAction(v ApplicationSecurityWafCustomRuleAction) {
	o.Action = &v
}

// GetBlocking returns the Blocking field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetBlocking() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Blocking
}

// GetBlockingOk returns a tuple with the Blocking field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetBlockingOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blocking, true
}

// SetBlocking sets field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetBlocking(v bool) {
	o.Blocking = v
}

// GetConditions returns the Conditions field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetConditions() []ApplicationSecurityWafCustomRuleCondition {
	if o == nil {
		var ret []ApplicationSecurityWafCustomRuleCondition
		return ret
	}
	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetConditionsOk() (*[]ApplicationSecurityWafCustomRuleCondition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Conditions, true
}

// SetConditions sets field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetConditions(v []ApplicationSecurityWafCustomRuleCondition) {
	o.Conditions = v
}

// GetEnabled returns the Enabled field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}
	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetMetadata() ApplicationSecurityWafCustomRuleMetadata {
	if o == nil || o.Metadata == nil {
		var ret ApplicationSecurityWafCustomRuleMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetMetadataOk() (*ApplicationSecurityWafCustomRuleMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) HasMetadata() bool {
	return o != nil && o.Metadata != nil
}

// SetMetadata gets a reference to the given ApplicationSecurityWafCustomRuleMetadata and assigns it to the Metadata field.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetMetadata(v ApplicationSecurityWafCustomRuleMetadata) {
	o.Metadata = &v
}

// GetName returns the Name field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}
	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetName(v string) {
	o.Name = v
}

// GetPathGlob returns the PathGlob field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetPathGlob() string {
	if o == nil || o.PathGlob == nil {
		var ret string
		return ret
	}
	return *o.PathGlob
}

// GetPathGlobOk returns a tuple with the PathGlob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetPathGlobOk() (*string, bool) {
	if o == nil || o.PathGlob == nil {
		return nil, false
	}
	return o.PathGlob, true
}

// HasPathGlob returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) HasPathGlob() bool {
	return o != nil && o.PathGlob != nil
}

// SetPathGlob gets a reference to the given string and assigns it to the PathGlob field.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetPathGlob(v string) {
	o.PathGlob = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetScope() []ApplicationSecurityWafCustomRuleScope {
	if o == nil || o.Scope == nil {
		var ret []ApplicationSecurityWafCustomRuleScope
		return ret
	}
	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetScopeOk() (*[]ApplicationSecurityWafCustomRuleScope, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return &o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) HasScope() bool {
	return o != nil && o.Scope != nil
}

// SetScope gets a reference to the given []ApplicationSecurityWafCustomRuleScope and assigns it to the Scope field.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetScope(v []ApplicationSecurityWafCustomRuleScope) {
	o.Scope = v
}

// GetTags returns the Tags field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetTags() ApplicationSecurityWafCustomRuleTags {
	if o == nil {
		var ret ApplicationSecurityWafCustomRuleTags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ApplicationSecurityWafCustomRuleAttributes) GetTagsOk() (*ApplicationSecurityWafCustomRuleTags, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tags, true
}

// SetTags sets field value.
func (o *ApplicationSecurityWafCustomRuleAttributes) SetTags(v ApplicationSecurityWafCustomRuleTags) {
	o.Tags = v
}

// MarshalJSON serializes the struct using spec logic.
func (o ApplicationSecurityWafCustomRuleAttributes) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	toSerialize["blocking"] = o.Blocking
	toSerialize["conditions"] = o.Conditions
	toSerialize["enabled"] = o.Enabled
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["name"] = o.Name
	if o.PathGlob != nil {
		toSerialize["path_glob"] = o.PathGlob
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	toSerialize["tags"] = o.Tags

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *ApplicationSecurityWafCustomRuleAttributes) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Action     *ApplicationSecurityWafCustomRuleAction      `json:"action,omitempty"`
		Blocking   *bool                                        `json:"blocking"`
		Conditions *[]ApplicationSecurityWafCustomRuleCondition `json:"conditions"`
		Enabled    *bool                                        `json:"enabled"`
		Metadata   *ApplicationSecurityWafCustomRuleMetadata    `json:"metadata,omitempty"`
		Name       *string                                      `json:"name"`
		PathGlob   *string                                      `json:"path_glob,omitempty"`
		Scope      []ApplicationSecurityWafCustomRuleScope      `json:"scope,omitempty"`
		Tags       *ApplicationSecurityWafCustomRuleTags        `json:"tags"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	if all.Blocking == nil {
		return fmt.Errorf("required field blocking missing")
	}
	if all.Conditions == nil {
		return fmt.Errorf("required field conditions missing")
	}
	if all.Enabled == nil {
		return fmt.Errorf("required field enabled missing")
	}
	if all.Name == nil {
		return fmt.Errorf("required field name missing")
	}
	if all.Tags == nil {
		return fmt.Errorf("required field tags missing")
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"action", "blocking", "conditions", "enabled", "metadata", "name", "path_glob", "scope", "tags"})
	} else {
		return err
	}

	hasInvalidField := false
	if all.Action != nil && all.Action.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Action = all.Action
	o.Blocking = *all.Blocking
	o.Conditions = *all.Conditions
	o.Enabled = *all.Enabled
	if all.Metadata != nil && all.Metadata.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Metadata = all.Metadata
	o.Name = *all.Name
	o.PathGlob = all.PathGlob
	o.Scope = all.Scope
	if all.Tags.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Tags = *all.Tags

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
