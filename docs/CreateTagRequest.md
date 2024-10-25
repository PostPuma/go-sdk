# CreateTagRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**HexColor** | **string** |  | 

## Methods

### NewCreateTagRequest

`func NewCreateTagRequest(name string, hexColor string, ) *CreateTagRequest`

NewCreateTagRequest instantiates a new CreateTagRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTagRequestWithDefaults

`func NewCreateTagRequestWithDefaults() *CreateTagRequest`

NewCreateTagRequestWithDefaults instantiates a new CreateTagRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTagRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTagRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTagRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHexColor

`func (o *CreateTagRequest) GetHexColor() string`

GetHexColor returns the HexColor field if non-nil, zero value otherwise.

### GetHexColorOk

`func (o *CreateTagRequest) GetHexColorOk() (*string, bool)`

GetHexColorOk returns a tuple with the HexColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexColor

`func (o *CreateTagRequest) SetHexColor(v string)`

SetHexColor sets HexColor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


