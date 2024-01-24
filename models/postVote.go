package models

import (
	"fmt"
	"time"

	"github.com/Benjam11n/db"
	"gorm.io/gorm"
)

type PostVote struct {
	gorm.Model
	UserId    uint      `json:"userId"`
	PostId    uint      `json:"postId"`
	VoteValue int64     `json:"voteValue"`
	CreatedAt time.Time `json:"createdAt"`
}

func (pv *PostVote) Save() error {
	db := db.DB
	result := db.Create(pv)
	return result.Error
}

func (pv *PostVote) Update() error {
	db := db.DB
	fmt.Println(pv.VoteValue, pv.ID)
	result := db.Model(pv).Where("user_id = ? AND post_id = ?", pv.UserId, pv.PostId).Updates(PostVote{VoteValue: pv.VoteValue})
	return result.Error
}

func GetAllPostVoteValue(postId uint) (uint, error) {
	db := db.DB
	var totalVotes uint
	result := db.Table("post_votes").Where("post_id = ?", postId).Pluck("SUM(vote_value)", &totalVotes)
	return totalVotes, result.Error
}

func GetPostVotes(postId uint) ([]PostVote, error) {
	db := db.DB
	var votes []PostVote
	result := db.Where("post_id = ?", postId).Find(&votes)
	return votes, result.Error
}

func GetAllPostVotes() ([]PostVote, error) {
	db := db.DB
	var votes []PostVote
	result := db.Find(&votes)
	return votes, result.Error
}

func DeletePostVotes(userId, postId uint) error {
	db := db.DB
	result := db.Where("user_id = ? AND post_id = ?", userId, postId).Delete(&PostVote{})
	return result.Error
}
