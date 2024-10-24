# UpdateTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**HexColor** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateTagRequest

`func NewUpdateTagRequest(name string, ) *UpdateTagRequest`

NewUpdateTagRequest instantiates a new UpdateTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTagRequestWithDefaults

`func NewUpdateTagRequestWithDefaults() *UpdateTagRequest`

NewUpdateTagRequestWithDefaults instantiates a new UpdateTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateTagRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHexColor

`func (o *UpdateTagRequest) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *UpdateTagRequest) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *UpdateTagRequest) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.

### HasHexColor

`func (o *UpdateTagRequest) HasHexColor() bool`

HasHexColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


