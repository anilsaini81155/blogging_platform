go mod init github.com/anilsaini81155/blogging_platform

## installing the protoc and and setting the env variable

## below command to generate the go file from poroto file
protoc --go_out=. --go-grpc_out=. blog.proto

## installing the grpcurl and below command to check the grpcurl version installed

grpcurl --version

## Steps to run the project 

go build

go run blog_service.go

### showing the available list of services 

grpcurl -plaintext localhost:50051 list

## listing out methods available with BlogService

grpcurl -plaintext localhost:50051 list blog.BlogService

## Create post request

grpcurl -plaintext -d '{
    "title": "First Post",
    "content": "This is the content of the first post",
    "author": "Anil",
    "publication_date": "2024-09-05T00:00:00Z",
    "tags": ["go", "grpc"]
}' localhost:50051 blog.BlogService.CreatePost

## Read request

grpcurl -plaintext -d '{
    "post_id": "post-1"
}' localhost:50051 blog.BlogService.ReadPost

## Update Post

grpcurl -plaintext -d '{
    "post_id": "post-1",
    "title": "Updated Post Title",
    "content": "This is the updated content.",
    "author": "Anil Saini",
    "tags": ["updated", "golang", "grpc"]
}' localhost:50051 blog.BlogService.UpdatePost


## Delete post

grpcurl -plaintext -d '{
    "post_id": "post-1"
}' localhost:50051 blog.BlogService.DeletePost


## Below command to run the test cases

go test -v
