# LogsPipelinesOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PipelineIds** | Pointer to **[]string** | Ordered Array of &lt;PIPELINE_ID&gt; strings, the order of pipeline IDs in the array define the overall Pipelines order for Datadog. | 

## Methods

### NewLogsPipelinesOrder

`func NewLogsPipelinesOrder(pipelineIds []string, ) *LogsPipelinesOrder`

NewLogsPipelinesOrder instantiates a new LogsPipelinesOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsPipelinesOrderWithDefaults

`func NewLogsPipelinesOrderWithDefaults() *LogsPipelinesOrder`

NewLogsPipelinesOrderWithDefaults instantiates a new LogsPipelinesOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPipelineIds

`func (o *LogsPipelinesOrder) GetPipelineIds() []string`

GetPipelineIds returns the PipelineIds field if non-nil, zero value otherwise.

### GetPipelineIdsOk

`func (o *LogsPipelinesOrder) GetPipelineIdsOk() ([]string, bool)`

GetPipelineIdsOk returns a tuple with the PipelineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPipelineIds

`func (o *LogsPipelinesOrder) HasPipelineIds() bool`

HasPipelineIds returns a boolean if a field has been set.

### SetPipelineIds

`func (o *LogsPipelinesOrder) SetPipelineIds(v []string)`

SetPipelineIds gets a reference to the given []string and assigns it to the PipelineIds field.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

