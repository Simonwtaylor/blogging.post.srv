package repository

import (
	"github.com/jinzhu/gorm"

	post "post/proto/post"
)

type Repository interface {
	CreateUser(user interface{}) error
	GetAllUsers() error
	CreatePost(post interface{}) error
	GetAllPosts() error
}

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) CreateUser(user post.UserInput) error {
	return nil
}

func (r *PostRepository) CreatePost(user post.CreatePostRequest) error {
	return nil
}
