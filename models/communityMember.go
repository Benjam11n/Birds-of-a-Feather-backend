package models

import "github.com/Benjam11n/db"

type CommunityMember struct {
	ID          uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId      uint `gorm:"not null" binding:"required" json:"userId"`
	CommunityId uint `gorm:"not null" binding:"required" json:"communityId"`
}

func (c *CommunityMember) Save() error {
	db := db.DB
	result := db.Create(c)
	return result.Error
}

func GetCommunityMembers(communityId uint) ([]CommunityMember, error) {
	db := db.DB
	var members []CommunityMember
	result := db.Where("community_id = ?", communityId).Find(&members)
	return members, result.Error
}

func GetAllCommunityMembers() ([]CommunityMember, error) {
	db := db.DB
	var members []CommunityMember
	result := db.Find(&members)
	return members, result.Error
}

func GetAllUserCommunities(userId uint) ([]Community, error) {
	db := db.DB

	var communityIDs []int64
	result := db.Model(&CommunityMember{}).Select("community_id").Where("user_id = ?", userId).Pluck("community_id", &communityIDs)
	if result.Error != nil {
		return nil, result.Error
	}

	var communities []Community
	result = db.Where("id IN ?", communityIDs).Find(&communities)
	if result.Error != nil {
		return nil, result.Error
	}

	return communities, nil
}

func (c *CommunityMember) Delete() error {
	db := db.DB
	result := db.Delete(c, "user_id = ? AND community_id = ?", c.UserId, c.CommunityId)
	return result.Error
}
