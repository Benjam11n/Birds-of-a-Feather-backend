package db

import (
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	UserBio   string    `json:"userBio"`
	Birthday  time.Time `json:"birthday"`
	AvatarUrl string    `json:"avatarUrl"`
}

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

type PostVote struct {
	gorm.Model
	UserId    uint      `json:"userId"`
	PostId    uint      `json:"postId"`
	VoteValue int64     `json:"voteValue"`
	CreatedAt time.Time `json:"createdAt"`
}

type Reply struct {
	gorm.Model
	Content   string `binding:"required" json:"content"`
	UserId    uint   `json:"userId"`
	ParentId  uint   `json:"parentId"`
	ImagesUrl string `json:"imagesUrl"`
	Edited    bool   `json:"edited"`
	Views     uint   `json:"views"`
}

type ReplyVote struct {
	gorm.Model
	UserId    uint      `json:"userId"`
	ReplyId   uint      `json:"replyId"`
	VoteValue int64     `json:"voteValue"`
	CreatedAt time.Time `json:"createdAt"`
}

type Follow struct {
	gorm.Model
	FollowerId uint      `json:"followerId"`
	FolloweeId uint      `json:"followeeId"`
	CreatedAt  time.Time `json:"createdAt"`
}

type Community struct {
	gorm.Model
	UserId      uint   `json:"userId"`
	Title       string `binding:"required" json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	IconUrl     string `json:"iconUrl"`
}

type CommunityMember struct {
	ID          uint `gorm:"primaryKey;autoIncrement" json:"id"`
	UserId      uint `gorm:"not null" binding:"required" json:"userId"`
	CommunityId uint `gorm:"not null" binding:"required" json:"communityId"`
}

func InitDB() {
	var err error
	databaseURL := os.Getenv("DATABASE_URL")
	// databaseURL := "postgres://postgres:(Passcode336133)@db:5432/Birds-of-a-Feather"

	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  databaseURL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database" + err.Error())
	}

	// AutoMigrate tables
	migrateModels()

	// Optionally, set connection pool parameters
	// DB.SetMaxOpenConns(10)
	// DB.SetMaxIdleConns(5)
	// DB.SetConnMaxLifetime(time.Second * 15)
}

// migrateModels performs auto-migration of models.
func migrateModels() {
	if err := DB.AutoMigrate(
		&User{},
		&Post{},
		&PostVote{},
		&Reply{},
		&ReplyVote{},
		&Follow{},
		&Community{},
		&CommunityMember{},
	); err != nil {
		panic("Could not create tables" + err.Error())
	}
}
