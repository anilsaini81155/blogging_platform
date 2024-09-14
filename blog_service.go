package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	// Import reflection package
	pb "github.com/anilsaini81155/blogging_platform/blogpb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type BlogServiceServer struct {
	pb.UnimplementedBlogServiceServer
	posts map[string]*pb.Post
	mu    sync.Mutex // Ensures safe concurrent access

}

func NewBlogServiceServer() *BlogServiceServer {
	return &BlogServiceServer{
		posts: make(map[string]*pb.Post),
	}
}

// Create a new blog post
func (s *BlogServiceServer) CreatePost(ctx context.Context, req *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	s.mu.Lock()         // Lock to ensure thread safety
	defer s.mu.Unlock() // Unlock after the function completes

	postID := fmt.Sprintf("post-%d", len(s.posts)+1)
	newPost := &pb.Post{
		PostId:          postID,
		Title:           req.Title,
		Content:         req.Content,
		Author:          req.Author,
		PublicationDate: req.PublicationDate,
		Tags:            req.Tags,
	}
	s.posts[postID] = newPost

	return &pb.CreatePostResponse{
		Post: newPost,
	}, nil
}

// Read a blog post by its ID
func (s *BlogServiceServer) ReadPost(ctx context.Context, req *pb.ReadPostRequest) (*pb.ReadPostResponse, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[req.PostId]
	if !exists {
		return &pb.ReadPostResponse{
			Error: "Post not found",
		}, nil
	}
	return &pb.ReadPostResponse{
		Post: post,
	}, nil
}

// Update an existing blog post
func (s *BlogServiceServer) UpdatePost(ctx context.Context, req *pb.UpdatePostRequest) (*pb.UpdatePostResponse, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	post, exists := s.posts[req.PostId]
	if !exists {
		return &pb.UpdatePostResponse{
			Error: "Post not found",
		}, nil
	}

	post.Title = req.Title
	post.Content = req.Content
	post.Author = req.Author
	post.Tags = req.Tags

	return &pb.UpdatePostResponse{
		Post: post,
	}, nil
}

// Delete a blog post by its ID
func (s *BlogServiceServer) DeletePost(ctx context.Context, req *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.posts[req.PostId]
	if !exists {
		return &pb.DeletePostResponse{
			Message: "Post not found",
		}, nil
	}

	delete(s.posts, req.PostId)
	return &pb.DeletePostResponse{
		Message: "Post deleted successfully",
	}, nil
}

func main() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
		panic(err)
	}

	server := grpc.NewServer()
	blogService := NewBlogServiceServer()

	pb.RegisterBlogServiceServer(server, blogService)

	log.Println("Blog gRPC server started...")

	reflection.Register(server)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
