// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2019-Present Datadog, Inc.

package datadogV1

import (
	"github.com/DataDog/datadog-api-client-go/v2/api/datadog"
)

// MonitorSearchResult Holds search results.
type MonitorSearchResult struct {
	// Classification of the monitor.
	Classification *string `json:"classification,omitempty"`
	// Object describing the creator of the shared element.
	Creator *Creator `json:"creator,omitempty"`
	// ID of the monitor.
	Id *int64 `json:"id,omitempty"`
	// Latest timestamp the monitor triggered.
	LastTriggeredTs datadog.NullableInt64 `json:"last_triggered_ts,omitempty"`
	// Metrics used by the monitor.
	Metrics []string `json:"metrics,omitempty"`
	// The monitor name.
	Name *string `json:"name,omitempty"`
	// The notification triggered by the monitor.
	Notifications []MonitorSearchResultNotification `json:"notifications,omitempty"`
	// The ID of the organization.
	OrgId *int64 `json:"org_id,omitempty"`
	// Quality issues detected with the monitor.
	QualityIssues []string `json:"quality_issues,omitempty"`
	// The monitor query.
	Query *string `json:"query,omitempty"`
	// The scope(s) to which the downtime applies, for example `host:app2`.
	// Provide multiple scopes as a comma-separated list, for example `env:dev,env:prod`.
	// The resulting downtime applies to sources that matches ALL provided scopes
	// (that is `env:dev AND env:prod`), NOT any of them.
	Scopes []string `json:"scopes,omitempty"`
	// The different states your monitor can be in.
	Status *MonitorOverallStates `json:"status,omitempty"`
	// Tags associated with the monitor.
	Tags []string `json:"tags,omitempty"`
	// The type of the monitor. For more information about `type`, see the [monitor options](https://docs.datadoghq.com/monitors/guide/monitor_api_options/) docs.
	Type *MonitorType `json:"type,omitempty"`
	// UnparsedObject contains the raw value of the object if there was an error when deserializing into the struct
	UnparsedObject       map[string]interface{} `json:"-"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// NewMonitorSearchResult instantiates a new MonitorSearchResult object.
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed.
func NewMonitorSearchResult() *MonitorSearchResult {
	this := MonitorSearchResult{}
	return &this
}

// NewMonitorSearchResultWithDefaults instantiates a new MonitorSearchResult object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set.
func NewMonitorSearchResultWithDefaults() *MonitorSearchResult {
	this := MonitorSearchResult{}
	return &this
}

// GetClassification returns the Classification field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetClassification() string {
	if o == nil || o.Classification == nil {
		var ret string
		return ret
	}
	return *o.Classification
}

// GetClassificationOk returns a tuple with the Classification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetClassificationOk() (*string, bool) {
	if o == nil || o.Classification == nil {
		return nil, false
	}
	return o.Classification, true
}

// HasClassification returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasClassification() bool {
	return o != nil && o.Classification != nil
}

// SetClassification gets a reference to the given string and assigns it to the Classification field.
func (o *MonitorSearchResult) SetClassification(v string) {
	o.Classification = &v
}

// GetCreator returns the Creator field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetCreator() Creator {
	if o == nil || o.Creator == nil {
		var ret Creator
		return ret
	}
	return *o.Creator
}

// GetCreatorOk returns a tuple with the Creator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetCreatorOk() (*Creator, bool) {
	if o == nil || o.Creator == nil {
		return nil, false
	}
	return o.Creator, true
}

// HasCreator returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasCreator() bool {
	return o != nil && o.Creator != nil
}

// SetCreator gets a reference to the given Creator and assigns it to the Creator field.
func (o *MonitorSearchResult) SetCreator(v Creator) {
	o.Creator = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasId() bool {
	return o != nil && o.Id != nil
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *MonitorSearchResult) SetId(v int64) {
	o.Id = &v
}

// GetLastTriggeredTs returns the LastTriggeredTs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MonitorSearchResult) GetLastTriggeredTs() int64 {
	if o == nil || o.LastTriggeredTs.Get() == nil {
		var ret int64
		return ret
	}
	return *o.LastTriggeredTs.Get()
}

// GetLastTriggeredTsOk returns a tuple with the LastTriggeredTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned.
func (o *MonitorSearchResult) GetLastTriggeredTsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastTriggeredTs.Get(), o.LastTriggeredTs.IsSet()
}

// HasLastTriggeredTs returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasLastTriggeredTs() bool {
	return o != nil && o.LastTriggeredTs.IsSet()
}

// SetLastTriggeredTs gets a reference to the given datadog.NullableInt64 and assigns it to the LastTriggeredTs field.
func (o *MonitorSearchResult) SetLastTriggeredTs(v int64) {
	o.LastTriggeredTs.Set(&v)
}

// SetLastTriggeredTsNil sets the value for LastTriggeredTs to be an explicit nil.
func (o *MonitorSearchResult) SetLastTriggeredTsNil() {
	o.LastTriggeredTs.Set(nil)
}

// UnsetLastTriggeredTs ensures that no value is present for LastTriggeredTs, not even an explicit nil.
func (o *MonitorSearchResult) UnsetLastTriggeredTs() {
	o.LastTriggeredTs.Unset()
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetMetrics() []string {
	if o == nil || o.Metrics == nil {
		var ret []string
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetMetricsOk() (*[]string, bool) {
	if o == nil || o.Metrics == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasMetrics() bool {
	return o != nil && o.Metrics != nil
}

// SetMetrics gets a reference to the given []string and assigns it to the Metrics field.
func (o *MonitorSearchResult) SetMetrics(v []string) {
	o.Metrics = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasName() bool {
	return o != nil && o.Name != nil
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MonitorSearchResult) SetName(v string) {
	o.Name = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetNotifications() []MonitorSearchResultNotification {
	if o == nil || o.Notifications == nil {
		var ret []MonitorSearchResultNotification
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetNotificationsOk() (*[]MonitorSearchResultNotification, bool) {
	if o == nil || o.Notifications == nil {
		return nil, false
	}
	return &o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasNotifications() bool {
	return o != nil && o.Notifications != nil
}

// SetNotifications gets a reference to the given []MonitorSearchResultNotification and assigns it to the Notifications field.
func (o *MonitorSearchResult) SetNotifications(v []MonitorSearchResultNotification) {
	o.Notifications = v
}

// GetOrgId returns the OrgId field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetOrgId() int64 {
	if o == nil || o.OrgId == nil {
		var ret int64
		return ret
	}
	return *o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetOrgIdOk() (*int64, bool) {
	if o == nil || o.OrgId == nil {
		return nil, false
	}
	return o.OrgId, true
}

// HasOrgId returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasOrgId() bool {
	return o != nil && o.OrgId != nil
}

// SetOrgId gets a reference to the given int64 and assigns it to the OrgId field.
func (o *MonitorSearchResult) SetOrgId(v int64) {
	o.OrgId = &v
}

// GetQualityIssues returns the QualityIssues field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetQualityIssues() []string {
	if o == nil || o.QualityIssues == nil {
		var ret []string
		return ret
	}
	return o.QualityIssues
}

// GetQualityIssuesOk returns a tuple with the QualityIssues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetQualityIssuesOk() (*[]string, bool) {
	if o == nil || o.QualityIssues == nil {
		return nil, false
	}
	return &o.QualityIssues, true
}

// HasQualityIssues returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasQualityIssues() bool {
	return o != nil && o.QualityIssues != nil
}

// SetQualityIssues gets a reference to the given []string and assigns it to the QualityIssues field.
func (o *MonitorSearchResult) SetQualityIssues(v []string) {
	o.QualityIssues = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetQuery() string {
	if o == nil || o.Query == nil {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetQueryOk() (*string, bool) {
	if o == nil || o.Query == nil {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasQuery() bool {
	return o != nil && o.Query != nil
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *MonitorSearchResult) SetQuery(v string) {
	o.Query = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetScopesOk() (*[]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return &o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasScopes() bool {
	return o != nil && o.Scopes != nil
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *MonitorSearchResult) SetScopes(v []string) {
	o.Scopes = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetStatus() MonitorOverallStates {
	if o == nil || o.Status == nil {
		var ret MonitorOverallStates
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetStatusOk() (*MonitorOverallStates, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasStatus() bool {
	return o != nil && o.Status != nil
}

// SetStatus gets a reference to the given MonitorOverallStates and assigns it to the Status field.
func (o *MonitorSearchResult) SetStatus(v MonitorOverallStates) {
	o.Status = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return &o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasTags() bool {
	return o != nil && o.Tags != nil
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *MonitorSearchResult) SetTags(v []string) {
	o.Tags = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *MonitorSearchResult) GetType() MonitorType {
	if o == nil || o.Type == nil {
		var ret MonitorType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MonitorSearchResult) GetTypeOk() (*MonitorType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *MonitorSearchResult) HasType() bool {
	return o != nil && o.Type != nil
}

// SetType gets a reference to the given MonitorType and assigns it to the Type field.
func (o *MonitorSearchResult) SetType(v MonitorType) {
	o.Type = &v
}

// MarshalJSON serializes the struct using spec logic.
func (o MonitorSearchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UnparsedObject != nil {
		return datadog.Marshal(o.UnparsedObject)
	}
	if o.Classification != nil {
		toSerialize["classification"] = o.Classification
	}
	if o.Creator != nil {
		toSerialize["creator"] = o.Creator
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastTriggeredTs.IsSet() {
		toSerialize["last_triggered_ts"] = o.LastTriggeredTs.Get()
	}
	if o.Metrics != nil {
		toSerialize["metrics"] = o.Metrics
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Notifications != nil {
		toSerialize["notifications"] = o.Notifications
	}
	if o.OrgId != nil {
		toSerialize["org_id"] = o.OrgId
	}
	if o.QualityIssues != nil {
		toSerialize["quality_issues"] = o.QualityIssues
	}
	if o.Query != nil {
		toSerialize["query"] = o.Query
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}
	return datadog.Marshal(toSerialize)
}

// UnmarshalJSON deserializes the given payload.
func (o *MonitorSearchResult) UnmarshalJSON(bytes []byte) (err error) {
	all := struct {
		Classification  *string                           `json:"classification,omitempty"`
		Creator         *Creator                          `json:"creator,omitempty"`
		Id              *int64                            `json:"id,omitempty"`
		LastTriggeredTs datadog.NullableInt64             `json:"last_triggered_ts,omitempty"`
		Metrics         []string                          `json:"metrics,omitempty"`
		Name            *string                           `json:"name,omitempty"`
		Notifications   []MonitorSearchResultNotification `json:"notifications,omitempty"`
		OrgId           *int64                            `json:"org_id,omitempty"`
		QualityIssues   []string                          `json:"quality_issues,omitempty"`
		Query           *string                           `json:"query,omitempty"`
		Scopes          []string                          `json:"scopes,omitempty"`
		Status          *MonitorOverallStates             `json:"status,omitempty"`
		Tags            []string                          `json:"tags,omitempty"`
		Type            *MonitorType                      `json:"type,omitempty"`
	}{}
	if err = datadog.Unmarshal(bytes, &all); err != nil {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}
	additionalProperties := make(map[string]interface{})
	if err = datadog.Unmarshal(bytes, &additionalProperties); err == nil {
		datadog.DeleteKeys(additionalProperties, &[]string{"classification", "creator", "id", "last_triggered_ts", "metrics", "name", "notifications", "org_id", "quality_issues", "query", "scopes", "status", "tags", "type"})
	} else {
		return err
	}

	hasInvalidField := false
	o.Classification = all.Classification
	if all.Creator != nil && all.Creator.UnparsedObject != nil && o.UnparsedObject == nil {
		hasInvalidField = true
	}
	o.Creator = all.Creator
	o.Id = all.Id
	o.LastTriggeredTs = all.LastTriggeredTs
	o.Metrics = all.Metrics
	o.Name = all.Name
	o.Notifications = all.Notifications
	o.OrgId = all.OrgId
	o.QualityIssues = all.QualityIssues
	o.Query = all.Query
	o.Scopes = all.Scopes
	if all.Status != nil && !all.Status.IsValid() {
		hasInvalidField = true
	} else {
		o.Status = all.Status
	}
	o.Tags = all.Tags
	if all.Type != nil && !all.Type.IsValid() {
		hasInvalidField = true
	} else {
		o.Type = all.Type
	}

	if len(additionalProperties) > 0 {
		o.AdditionalProperties = additionalProperties
	}

	if hasInvalidField {
		return datadog.Unmarshal(bytes, &o.UnparsedObject)
	}

	return nil
}
