package models

import (
	"time"

	"example.com/benjamin/db"
	"gorm.io/gorm"
)

type ReplyVote struct {
	gorm.Model
	UserId    uint      `json:"userId"`
	ReplyId   uint      `json:"replyId"`
	VoteValue int64     `json:"voteValue"`
	CreatedAt time.Time `json:"createdAt"`
}

func (pv *ReplyVote) Save() error {
	db := db.DB
	result := db.Create(pv)
	return result.Error
}

func (rv *ReplyVote) Update() error {
	db := db.DB
	result := db.Model(rv).Where("user_id = ?", rv.UserId).Update("vote_value", rv.VoteValue)
	return result.Error
}

func DeleteReplyVotes(userId, replyId uint) error {
	db := db.DB
	result := db.Where("user_id = ? AND reply_id = ?", userId, replyId).Delete(&ReplyVote{})
	return result.Error
}

func GetAllReplyVotes() ([]ReplyVote, error) {
	db := db.DB
	var totalVotes []ReplyVote
	result := db.Find(&totalVotes)
	return totalVotes, result.Error
}

func GetReplyVotes(replyId uint) ([]ReplyVote, error) {
	db := db.DB
	var totalVotes []ReplyVote
	result := db.Where("reply_id = ?", replyId).Find(&totalVotes)
	return totalVotes, result.Error
}
