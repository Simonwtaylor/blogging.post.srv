package repository

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"

	"post/entities"
	post "post/proto/post"
)

type Repository interface {
	CreateUser(user post.UserInput) (entities.User, error)
	GetAllUsers(query post.PaginatedQuery) ([]entities.User, error)
	CreatePost(post post.CreatePostRequest) (entities.Post, error)
	GetAllPosts(query post.PaginatedQuery) ([]entities.Post, error)
}

type PostRepository struct {
	db *gorm.DB
}

func (r *PostRepository) CreateUser(user post.UserInput) (*entities.User, error) {
	newUser := entities.User{
		Active:   user.Active,
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
	}

	isNew := r.db.NewRecord(newUser)

	if !isNew {
		return &entities.User{}, fmt.Errorf("unable to create user as it already contains an id, %v", newUser)
	}

	r.db.Create(&newUser)

	return &newUser, nil
}

func (r *PostRepository) CreatePost(post post.CreatePostRequest) (*entities.Post, error) {
	newPost := entities.Post{
		Content:    post.Content,
		DatePosted: time.Now(),
	}
	var user entities.User
	r.db.First(&user)

	newPost.User = user

	isNew := r.db.NewRecord(newPost)

	if !isNew {
		return &entities.Post{}, fmt.Errorf("unable to create post as it already contains an id, %v", newPost)
	}

	r.db.Create(&newPost)

	return &newPost, nil
}

func (r *PostRepository) GetAllUsers(query post.PaginatedQuery) ([]entities.User, error) {
	return []entities.User{}, nil
}

func (r *PostRepository) GetAllPosts(query post.PaginatedQuery) ([]entities.Post, error) {
	return []entities.Post{}, nil
}
