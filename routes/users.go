package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"example.com/benjamin/models"
	"example.com/benjamin/utils"
	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.AvatarUrl = "/assets/avatars/default.jpg"
	user.CreatedAt = time.Now().Local()
	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, user)
}

func Login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	retrievedUser, err := user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateToken(user.Email, retrievedUser.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
		return
	}

	//set the current userId to the context
	context.Set("userId", retrievedUser.ID)
	context.JSON(http.StatusOK, gin.H{"message": "User logged in successfully", "token": token})
}

func GetUsers(context *gin.Context) {
	users, err := models.GetAllUsers()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, users)
}

func GetUser(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByID(uint(userId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, user)
}

func UpdateUser(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user is the owner of the account
	if uint(userId) != context.GetUint("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	// Bind the JSON data
	var updatedUser models.User
	if err := context.ShouldBindJSON(&updatedUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser.ID = uint(userId)

	// Call the Update method to update the user in the database
	err = updatedUser.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func UpdateUserPassword(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user is the owner of the account
	if uint(userId) != context.GetUint("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}

	// Bind the JSON data
	var user models.User
	var updatedUser models.UpdateUserPassword
	if err := context.ShouldBindJSON(&updatedUser); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser.ID = uint(userId)
	user.Password = updatedUser.OldPassword
	user.Email = updatedUser.Email

	_, err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	// Call the Update method to update the user in the database
	err = updatedUser.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func UploadAvatar(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserByID(uint(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Delete the old image file
	if err := models.DeleteOldAvatar(user.AvatarUrl); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete old avatar"})
		return
	}

	// Get the uploaded file
	file, err := context.FormFile("avatarUrl")
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Print(err.Error())
		return
	}

	// Save the file with a unique filename
	avatarFilename := fmt.Sprintf("%d_%s", userId, file.Filename)
	err = context.SaveUploadedFile(file, "assets/avatars/"+avatarFilename)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save avatar"})
		return
	}
	var updatedUser models.User
	updatedUser.ID = uint(userId)
	updatedUser.AvatarUrl = "/assets/avatars/" + avatarFilename
	err = updatedUser.UpdateAvatar()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(context *gin.Context) {
	userId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// check if user is the owner of the account
	if uint(userId) != context.GetUint("userId") {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorised"})
		return
	}

	user, err := models.GetUserByID(uint(userId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = user.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "User deleted succcessfully"})
}

func GetCurrentUser(context *gin.Context) {
	userId := context.GetUint("userId")
	user, err := models.GetUserByID(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, user)
}

func Logout(context *gin.Context) {
	context.Set("userId", 0)
	context.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}
