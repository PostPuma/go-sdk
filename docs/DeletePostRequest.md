# DeletePostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Trash** | Pointer to **bool** | Whether to move items to trash or not. | [optional] 

## Methods

### NewDeletePostRequest

`func NewDeletePostRequest() *DeletePostRequest`

NewDeletePostRequest instantiates a new DeletePostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeletePostRequestWithDefaults

`func NewDeletePostRequestWithDefaults() *DeletePostRequest`

NewDeletePostRequestWithDefaults instantiates a new DeletePostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrash

`func (o *DeletePostRequest) GetTrash() bool`

GetTrash returns the Trash field if non-nil, zero value otherwise.

### GetTrashOk

`func (o *DeletePostRequest) GetTrashOk() (*bool, bool)`

GetTrashOk returns a tuple with the Trash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrash

`func (o *DeletePostRequest) SetTrash(v bool)`

SetTrash sets Trash field to given value.

### HasTrash

`func (o *DeletePostRequest) HasTrash() bool`

HasTrash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


