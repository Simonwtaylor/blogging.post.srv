package repository

import "github.com/jinzhu/gorm"

type Repository interface {
	CreateUser(user interface{}) error
	GetAllUsers() error
	CreatePost(post interface{}) error
	GetAllPosts() error
}

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) CreateUser(user interface{}) error {
	return nil
}
