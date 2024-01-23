package models

import (
	"example.com/benjamin/db"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	UserId      uint   `json:"userId"`
	CommunityId uint   `json:"communityId"`
	Title       string `binding:"required" json:"title"`
	Content     string `binding:"required" json:"content"`
	ImagesUrl   string `json:"imagesUrl"`
	Edited      bool   `json:"edited"`
	Views       uint   `json:"views"`
	Tags        string `json:"tags"`
}

func (p *Post) Save() error {
	db := db.DB
	result := db.Create(p)
	return result.Error
}

func GetAllPosts() ([]Post, error) {
	db := db.DB
	var posts []Post
	result := db.Find(&posts)
	return posts, result.Error
}

func GetPostByID(id uint) (*Post, error) {
	db := db.DB
	var post Post
	result := db.First(&post, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &post, result.Error
}

func (post *Post) Update() error {
	db := db.DB
	result := db.Model(post).Updates(Post{Title: post.Title, Content: post.Content, Edited: true, Tags: post.Tags})
	return result.Error
}

func (post *Post) Delete() error {
	db := db.DB
	result := db.Delete(post, "id = ? AND user_id = ?", post.ID, post.UserId)
	return result.Error
}
