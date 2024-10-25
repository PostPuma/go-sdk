# PostPuma\MediaAPI

All URIs are relative to *https://app.postpuma.com/app/5afgg2-1egj4n-7612ng-g313ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMediaFiles**](MediaAPI.md#DeleteMediaFiles) | **Delete** /media | Delete media files
[**GetMediaFile**](MediaAPI.md#GetMediaFile) | **Get** /media/{mediaUuid} | Get media file
[**ListMediaFiles**](MediaAPI.md#ListMediaFiles) | **Get** /media | List media files
[**UploadMediaFile**](MediaAPI.md#UploadMediaFile) | **Post** /media | Upload media file



## DeleteMediaFiles

> DeleteMediaFiles200Response DeleteMediaFiles(ctx).DeleteMediaFilesRequest(deleteMediaFilesRequest).Execute()

Delete media files



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PostPuma/go"
)

func main() {
	deleteMediaFilesRequest := *openapiclient.NewDeleteMediaFilesRequest([]int32{int32(123)}) // DeleteMediaFilesRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.DeleteMediaFiles(context.Background()).DeleteMediaFilesRequest(deleteMediaFilesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.DeleteMediaFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteMediaFiles`: DeleteMediaFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.DeleteMediaFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMediaFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deleteMediaFilesRequest** | [**DeleteMediaFilesRequest**](DeleteMediaFilesRequest.md) |  | 

### Return type

[**DeleteMediaFiles200Response**](DeleteMediaFiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMediaFile

> MediaFile GetMediaFile(ctx, mediaUuid).Execute()

Get media file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PostPuma/go"
)

func main() {
	mediaUuid := "mediaUuid_example" // string | Media UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.GetMediaFile(context.Background(), mediaUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.GetMediaFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMediaFile`: MediaFile
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.GetMediaFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mediaUuid** | **string** | Media UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMediaFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MediaFile**](MediaFile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMediaFiles

> ListMediaFiles200Response ListMediaFiles(ctx).Page(page).Execute()

List media files



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PostPuma/go"
)

func main() {
	page := int32(56) // int32 | Page number (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.ListMediaFiles(context.Background()).Page(page).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.ListMediaFiles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMediaFiles`: ListMediaFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.ListMediaFiles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMediaFilesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | Page number | 

### Return type

[**ListMediaFiles200Response**](ListMediaFiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadMediaFile

> MediaFile UploadMediaFile(ctx).File(file).Execute()

Upload media file



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PostPuma/go"
)

func main() {
	file := os.NewFile(1234, "some_file") // *os.File |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MediaAPI.UploadMediaFile(context.Background()).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MediaAPI.UploadMediaFile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UploadMediaFile`: MediaFile
	fmt.Fprintf(os.Stdout, "Response from `MediaAPI.UploadMediaFile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadMediaFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

[**MediaFile**](MediaFile.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

