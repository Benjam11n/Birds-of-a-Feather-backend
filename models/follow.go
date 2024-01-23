package models

import (
	"time"

	"example.com/benjamin/db"
	"gorm.io/gorm"
)

type Follow struct {
	gorm.Model
	FollowerId uint      `json:"followerId"`
	FolloweeId uint      `json:"followeeId"`
	CreatedAt  time.Time `json:"createdAt"`
}

func (f *Follow) Save() error {
	db := db.DB
	result := db.Create(f)
	return result.Error
}

func GetAllFollows() ([]Follow, error) {
	db := db.DB
	var follows []Follow
	result := db.Find(&follows)
	return follows, result.Error
}

func GetAllFollowees(followerId uint) ([]*User, error) {
	db := db.DB

	// Select the FolloweeIds where FollowerId matches
	var followeeIds []uint
	result := db.Model(&Follow{}).Where("follower_id = ?", followerId).Pluck("followee_id", &followeeIds)
	if result.Error != nil {
		return nil, result.Error
	}

	// Retrieve the user details for each FolloweeId
	var users []*User
	result = db.Model(&User{}).Where("id IN ?", followeeIds).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func GetFollowsByIDs(followerId, followeeId uint) (*Follow, error) {
	db := db.DB
	var follow Follow

	result := db.Where("follower_id = ? AND followee_id = ?", followerId, followeeId).First(&follow)
	if result.Error != nil {
		return nil, result.Error
	}

	return &follow, nil
}

func (f *Follow) Delete() error {
	db := db.DB
	result := db.Delete(f)
	return result.Error
}
