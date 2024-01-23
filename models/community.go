package models

import (
	"example.com/benjamin/db"
	"gorm.io/gorm"
)

type Community struct {
	gorm.Model
	UserId      uint   `json:"userId"`
	Title       string `binding:"required" json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IconUrl     string `json:"iconUrl"`
}

// Migrate the Community model to create the corresponding database table
func MigrateCommunity() {
	db := db.DB
	db.AutoMigrate(&Community{})
}

func (c *Community) Save() error {
	db := db.DB
	result := db.Create(c)
	return result.Error
}

func GetAllCommunities() ([]Community, error) {
	db := db.DB
	var communities []Community
	result := db.Find(&communities)
	return communities, result.Error
}

func GetCommunityById(communityId uint) (*Community, error) {
	db := db.DB
	var community Community
	result := db.First(&community, communityId)
	if result.Error != nil {
		return nil, result.Error
	}
	return &community, nil
}

func (c *Community) Update() error {
	db := db.DB
	result := db.Model(c).Updates(Community{Title: c.Title, Description: c.Description, Category: c.Category})
	return result.Error
}

func (c *Community) Delete() error {
	db := db.DB
	result := db.Delete(c)
	return result.Error
}
