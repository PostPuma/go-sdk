/*
PostPuma - OpenAPI 3.0

Testing MediaAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package PostPuma

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/PostPuma/go"
)

func Test_PostPuma_MediaAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test MediaAPIService DeleteMediaFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MediaAPI.DeleteMediaFiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService GetMediaFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var mediaUuid string

		resp, httpRes, err := apiClient.MediaAPI.GetMediaFile(context.Background(), mediaUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService ListMediaFiles", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MediaAPI.ListMediaFiles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test MediaAPIService UploadMediaFile", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.MediaAPI.UploadMediaFile(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
