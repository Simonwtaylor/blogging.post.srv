package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	post "post/proto/post"
	"post/repository"
)

type Post struct {
	repo *repository.PostRepository
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Post) Call(ctx context.Context, req *post.Request, rsp *post.Response) error {
	log.Info("Received Post.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *Post) CreateUser(ctx context.Context, req *post.CreateUserRequest, rsp *post.CreateUserResponse) error {
	log.Info("Received Post.CreateUser request")
	// e.repo.CreatePost(req.)
	rsp.User = &post.User{
		Active: true,
	}
	return nil
}

func (e *Post) CreatePost(ctx context.Context, req *post.CreatePostRequest, rsp *post.CreatePostResponse) error {
	log.Info("Received Post.CreatePost request")
	rsp.Post = &post.Post{
		Content: "",
	}
	return nil
}

func (e *Post) GetAllUsers(ctx context.Context, req *post.GetAllUsersRequest, rsp *post.GetAllUsersResponse) error {
	log.Info("Received Post.GetAllUsers request")
	rsp.Meta = &post.Meta{
		Cursor:     1,
		TotalItems: 12,
	}

	rsp.Users = []*post.User{
		&post.User{
			Email: "asdasda",
		},
	}
	return nil
}

func (e *Post) GetAllPosts(ctx context.Context, req *post.GetAllPostsRequest, rsp *post.GetAllPostsResponse) error {
	log.Info("Received Post.GetAllPosts request")
	rsp.Meta = &post.Meta{
		Cursor:     1,
		TotalItems: 12,
	}

	rsp.Posts = []*post.Post{
		&post.Post{
			Content: "asdasda",
		},
	}
	return nil
}
