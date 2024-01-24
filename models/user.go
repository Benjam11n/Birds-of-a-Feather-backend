package models

import (
	"errors"
	"os"
	"path/filepath"
	"time"

	"github.com/Benjam11n/db"
	"github.com/Benjam11n/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `gorm:"password" json:"password"`
	UserBio   string    `json:"userBio"`
	Birthday  time.Time `json:"birthday"`
	AvatarUrl string    `json:"avatarUrl"`
}

type UpdateUserPassword struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	OldPassword string `json:"oldPassword" gorm:"column:old_password"`
	NewPassword string `json:"newPassword" gorm:"column:new_password"`
}

func (u *User) Save() error {
	db := db.DB
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = hashedPassword

	result := db.Create(u)
	return result.Error
}

func (u *User) ValidateCredentials() (*User, error) {
	db := db.DB
	var retrievedUser User
	result := db.Model(u).Where("email = ?", u.Email).First(&retrievedUser)
	if result.Error != nil {
		return nil, result.Error
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedUser.Password)

	if !passwordIsValid {
		return nil, errors.New("credentials invalid")
	}

	return &retrievedUser, nil
}

func (u *User) Update() error {
	db := db.DB
	result := db.Model(u).Updates(User{Name: u.Name, Email: u.Email, AvatarUrl: u.AvatarUrl, UserBio: u.UserBio, Birthday: u.Birthday})
	return result.Error
}

func (u *UpdateUserPassword) Update() error {
	db := db.DB
	hashedPassword, err := utils.HashPassword(u.NewPassword)
	if err != nil {
		return err
	}

	result := db.Model(&User{}).Where("id = ? AND email = ?", u.ID, u.Email).Update("password", hashedPassword)

	return result.Error
}

func (u *User) UpdateAvatar() error {
	db := db.DB
	result := db.Model(u).Update("avatar_url", u.AvatarUrl)
	return result.Error
}

func DeleteOldAvatar(avatarURL string) error {
	// Extract the filename from the avatar URL
	filename := filepath.Base(avatarURL)

	// Construct the full path to the old avatar file
	oldAvatarPath := filepath.Join("assets/avatars", filename)

	// Attempt to delete the old avatar file
	if err := os.Remove(oldAvatarPath); err != nil && !os.IsNotExist(err) {
		return err
	}

	return nil
}

func (u *User) Delete() error {
	db := db.DB
	result := db.Delete(u)
	return result.Error
}

func GetAllUsers() ([]*User, error) {
	db := db.DB
	var users []*User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByID(id uint) (*User, error) {
	db := db.DB
	var retrievedUser User
	result := db.First(&retrievedUser, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &retrievedUser, nil
}
