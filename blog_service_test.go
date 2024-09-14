package main

import (
	"context"
	"testing"

	pb "github.com/anilsaini81155/blogging_platform/blogpb"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// TestCreatePost tests the CreatePost gRPC method
func TestCreatePost(t *testing.T) {
	blogService := NewBlogServiceServer()

	req := &pb.CreatePostRequest{
		Title:           "Test Post",
		Content:         "This is a test post.",
		Author:          "Test Author",
		PublicationDate: timestamppb.Now(),
		Tags:            []string{"test", "grpc"},
	}

	resp, err := blogService.CreatePost(context.Background(), req)
	assert.NoError(t, err, "CreatePost should not return an error")
	assert.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, "Test Post", resp.Post.Title, "Post title should match")
	assert.Equal(t, "Test Author", resp.Post.Author, "Post author should match")
	assert.Equal(t, []string{"test", "grpc"}, resp.Post.Tags, "Post tags should match")
}

// TestReadPost tests the ReadPost gRPC method
func TestReadPost(t *testing.T) {
	blogService := NewBlogServiceServer()

	// Create a post first
	createReq := &pb.CreatePostRequest{
		Title:           "Test Post",
		Content:         "This is a test post.",
		Author:          "Test Author",
		PublicationDate: timestamppb.Now(),
		Tags:            []string{"test", "grpc"},
	}
	createResp, _ := blogService.CreatePost(context.Background(), createReq)

	// Test ReadPost
	req := &pb.ReadPostRequest{
		PostId: createResp.Post.PostId,
	}
	resp, err := blogService.ReadPost(context.Background(), req)
	assert.NoError(t, err, "ReadPost should not return an error")
	assert.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, createResp.Post.PostId, resp.Post.PostId, "Post ID should match")
	assert.Equal(t, "Test Post", resp.Post.Title, "Post title should match")
	assert.Equal(t, "Test Author", resp.Post.Author, "Post author should match")
}

// TestUpdatePost tests the UpdatePost gRPC method
func TestUpdatePost(t *testing.T) {
	blogService := NewBlogServiceServer()

	// Create a post first
	createReq := &pb.CreatePostRequest{
		Title:           "Old Title",
		Content:         "Old Content",
		Author:          "Test Author",
		PublicationDate: timestamppb.Now(),
		Tags:            []string{"old"},
	}
	createResp, _ := blogService.CreatePost(context.Background(), createReq)

	// Test UpdatePost
	updateReq := &pb.UpdatePostRequest{
		PostId:  createResp.Post.PostId,
		Title:   "New Title",
		Content: "New Content",
		Author:  "Updated Author",
		Tags:    []string{"updated"},
	}
	resp, err := blogService.UpdatePost(context.Background(), updateReq)
	assert.NoError(t, err, "UpdatePost should not return an error")
	assert.NotNil(t, resp, "Response should not be nil")
	assert.Equal(t, "New Title", resp.Post.Title, "Post title should be updated")
	assert.Equal(t, "New Content", resp.Post.Content, "Post content should be updated")
	assert.Equal(t, "Updated Author", resp.Post.Author, "Post author should be updated")
	assert.Equal(t, []string{"updated"}, resp.Post.Tags, "Post tags should be updated")
}

// TestDeletePost tests the DeletePost gRPC method
func TestDeletePost(t *testing.T) {
	blogService := NewBlogServiceServer()

	// Create a post first
	createReq := &pb.CreatePostRequest{
		Title:           "Post to Delete",
		Content:         "Content of post to delete.",
		Author:          "Author",
		PublicationDate: timestamppb.Now(),
		Tags:            []string{"delete"},
	}
	createResp, _ := blogService.CreatePost(context.Background(), createReq)

	// Test DeletePost
	deleteReq := &pb.DeletePostRequest{
		PostId: createResp.Post.PostId,
	}
	deleteResp, err := blogService.DeletePost(context.Background(), deleteReq)
	assert.NoError(t, err, "DeletePost should not return an error")
	assert.NotNil(t, deleteResp, "Response should not be nil")
	assert.Equal(t, "Post deleted successfully", deleteResp.Message, "Delete message should match")

	// Try to read the deleted post to verify it's deleted
	readReq := &pb.ReadPostRequest{
		PostId: createResp.Post.PostId,
	}
	readResp, err := blogService.ReadPost(context.Background(), readReq)
	assert.NoError(t, err, "ReadPost should not return an error for deleted post")
	assert.NotNil(t, readResp, "Response should not be nil")
	assert.Equal(t, "Post not found", readResp.Error, "Post should not be found after deletion")
}
