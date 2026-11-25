# ExportJobDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Destination type can be Google Cloud Storage (gcs) or file. | 
**Compression** | Pointer to **string** | Compression applied to the exported file. | [optional] [default to "none"]
**Path** | Pointer to **string** | Directory on the exporter host where the file is written. Required for file destinations; must be omitted for gcs. | [optional] 
**Bucket** | Pointer to **string** | GCS bucket name for storage. Absent for file. | [optional] 

## Methods

### NewExportJobDestination

`func NewExportJobDestination(type_ string, ) *ExportJobDestination`

NewExportJobDestination instantiates a new ExportJobDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportJobDestinationWithDefaults

`func NewExportJobDestinationWithDefaults() *ExportJobDestination`

NewExportJobDestinationWithDefaults instantiates a new ExportJobDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ExportJobDestination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportJobDestination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportJobDestination) SetType(v string)`

SetType sets Type field to given value.


### GetCompression

`func (o *ExportJobDestination) GetCompression() string`

GetCompression returns the Compression field if non-nil, zero value otherwise.

### GetCompressionOk

`func (o *ExportJobDestination) GetCompressionOk() (*string, bool)`

GetCompressionOk returns a tuple with the Compression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompression

`func (o *ExportJobDestination) SetCompression(v string)`

SetCompression sets Compression field to given value.

### HasCompression

`func (o *ExportJobDestination) HasCompression() bool`

HasCompression returns a boolean if a field has been set.

### GetPath

`func (o *ExportJobDestination) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ExportJobDestination) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ExportJobDestination) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ExportJobDestination) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetBucket

`func (o *ExportJobDestination) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ExportJobDestination) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ExportJobDestination) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *ExportJobDestination) HasBucket() bool`

HasBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


