# PostPuma\TagsAPI

All URIs are relative to *https://app.postpuma.com/app/5afgg2-1egj4n-7612ng-g313ie*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTag**](TagsAPI.md#CreateTag) | **Post** /tags | Create tag
[**DeleteTag**](TagsAPI.md#DeleteTag) | **Delete** /tags/{tagUuid} | Delete tag
[**GetTag**](TagsAPI.md#GetTag) | **Get** /tags/{tagUuid} | Get tag
[**ListTags**](TagsAPI.md#ListTags) | **Get** /tags | List tags
[**UpdateTag**](TagsAPI.md#UpdateTag) | **Put** /tags/{tagUuid} | Update tag



## CreateTag

> Tag CreateTag(ctx).CreateTagRequest(createTagRequest).Execute()

Create tag



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
	createTagRequest := *openapiclient.NewCreateTagRequest("Name_example", "HexColor_example") // CreateTagRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.CreateTag(context.Background()).CreateTagRequest(createTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.CreateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateTag`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.CreateTag`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createTagRequest** | [**CreateTagRequest**](CreateTagRequest.md) |  | 

### Return type

[**Tag**](Tag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTag

> DeleteMediaFiles200Response DeleteTag(ctx, tagUuid).Execute()

Delete tag



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
	tagUuid := "tagUuid_example" // string | Tag UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.DeleteTag(context.Background(), tagUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.DeleteTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteTag`: DeleteMediaFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.DeleteTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagUuid** | **string** | Tag UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteMediaFiles200Response**](DeleteMediaFiles200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTag

> Tag GetTag(ctx, tagUuid).Execute()

Get tag



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
	tagUuid := "tagUuid_example" // string | Tag UUID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.GetTag(context.Background(), tagUuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.GetTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTag`: Tag
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.GetTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagUuid** | **string** | Tag UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Tag**](Tag.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTags

> ListTags200Response ListTags(ctx).Execute()

List tags



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.ListTags(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.ListTags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListTags`: ListTags200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.ListTags`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTagsRequest struct via the builder pattern


### Return type

[**ListTags200Response**](ListTags200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTag

> DeleteMediaFiles200Response UpdateTag(ctx, tagUuid).UpdateTagRequest(updateTagRequest).Execute()

Update tag



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
	tagUuid := "tagUuid_example" // string | Tag UUID
	updateTagRequest := *openapiclient.NewUpdateTagRequest("Name_example") // UpdateTagRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TagsAPI.UpdateTag(context.Background(), tagUuid).UpdateTagRequest(updateTagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TagsAPI.UpdateTag``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateTag`: DeleteMediaFiles200Response
	fmt.Fprintf(os.Stdout, "Response from `TagsAPI.UpdateTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tagUuid** | **string** | Tag UUID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateTagRequest** | [**UpdateTagRequest**](UpdateTagRequest.md) |  | 

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

