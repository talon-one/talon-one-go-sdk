# GeoJSONGeometry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The geometry type discriminator. | 
**Coordinates** | [**[][][][]float32**]([][][]float32.md) | The shapes in this group. Each one follows the same boundary structure as a polygon. | 
**Geometries** | [**[]GeoJSONGeometry**](GeoJSONGeometry.md) | The shapes contained in this group. | 

## Methods

### NewGeoJSONGeometry

`func NewGeoJSONGeometry(type_ string, coordinates [][][][]float32, geometries []GeoJSONGeometry, ) *GeoJSONGeometry`

NewGeoJSONGeometry instantiates a new GeoJSONGeometry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeoJSONGeometryWithDefaults

`func NewGeoJSONGeometryWithDefaults() *GeoJSONGeometry`

NewGeoJSONGeometryWithDefaults instantiates a new GeoJSONGeometry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GeoJSONGeometry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeoJSONGeometry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeoJSONGeometry) SetType(v string)`

SetType sets Type field to given value.


### GetCoordinates

`func (o *GeoJSONGeometry) GetCoordinates() [][][][]float32`

GetCoordinates returns the Coordinates field if non-nil, zero value otherwise.

### GetCoordinatesOk

`func (o *GeoJSONGeometry) GetCoordinatesOk() (*[][][][]float32, bool)`

GetCoordinatesOk returns a tuple with the Coordinates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoordinates

`func (o *GeoJSONGeometry) SetCoordinates(v [][][][]float32)`

SetCoordinates sets Coordinates field to given value.


### GetGeometries

`func (o *GeoJSONGeometry) GetGeometries() []GeoJSONGeometry`

GetGeometries returns the Geometries field if non-nil, zero value otherwise.

### GetGeometriesOk

`func (o *GeoJSONGeometry) GetGeometriesOk() (*[]GeoJSONGeometry, bool)`

GetGeometriesOk returns a tuple with the Geometries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometries

`func (o *GeoJSONGeometry) SetGeometries(v []GeoJSONGeometry)`

SetGeometries sets Geometries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


