/*
PostPuma - OpenAPI 3.0

Testing PostsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package PostPuma

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_PostPuma_PostsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PostsAPIService CreatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostsAPI.CreatePost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostsAPIService DeletePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postUuid string

		resp, httpRes, err := apiClient.PostsAPI.DeletePost(context.Background(), postUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostsAPIService DeletePosts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostsAPI.DeletePosts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostsAPIService GetPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postUuid string

		resp, httpRes, err := apiClient.PostsAPI.GetPost(context.Background(), postUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostsAPIService ListPosts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PostsAPI.ListPosts(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostsAPIService QueuePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postUuid string

		resp, httpRes, err := apiClient.PostsAPI.QueuePost(context.Background(), postUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostsAPIService SchedulePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postUuid string

		resp, httpRes, err := apiClient.PostsAPI.SchedulePost(context.Background(), postUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PostsAPIService UpdatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var postUuid string

		resp, httpRes, err := apiClient.PostsAPI.UpdatePost(context.Background(), postUuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
