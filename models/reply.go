package models

import (
	"example.com/benjamin/db"
	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	Content   string `binding:"required" json:"content"`
	UserId    uint   `json:"userId"`
	ParentId  uint   `json:"parentId"`
	ImagesUrl string `json:"imagesUrl"`
	Edited    bool   `json:"edited"`
	Views     uint   `json:"views"`
}

func (r *Reply) Save() error {
	db := db.DB
	result := db.Create(r)
	return result.Error
}

func GetAllReplies() ([]*Reply, error) {
	db := db.DB
	var replies []*Reply
	result := db.Find(&replies)
	return replies, result.Error
}

func GetAllPostReplies(parentId uint) ([]*Reply, error) {
	db := db.DB
	var replies []*Reply
	result := db.Model(&Reply{}).Where("parent_id = ?", parentId).Find(&replies)
	return replies, result.Error
}

func GetReplyByID(replyId uint) (*Reply, error) {
	db := db.DB
	var reply Reply
	result := db.First(&reply, replyId)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &reply, result.Error
}

func (r *Reply) Update() error {
	db := db.DB
	result := db.Model(r).Updates(Reply{Content: r.Content, Views: r.Views, Edited: true})
	return result.Error
}

func (r *Reply) Delete() error {
	db := db.DB
	result := db.Delete(r)
	return result.Error
}
