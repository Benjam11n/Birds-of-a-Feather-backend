package routes

import (
	"net/http"

	"github.com/Benjam11n/models"
	"github.com/gin-gonic/gin"
)

func GetFollowees(context *gin.Context) {
	userId := context.GetUint("userId")

	followees, err := models.GetAllFollowees(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, followees)
}

func GetFollows(context *gin.Context) {
	follows, err := models.GetAllFollows()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, follows)
}

func CreateFollow(context *gin.Context) {
	var follow models.Follow
	err := context.ShouldBindJSON(&follow)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	followerId := context.GetUint("userId")
	follow.FollowerId = followerId

	err = follow.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "follow created successfully", "follow": follow})
}

func DeleteFollow(context *gin.Context) {
	var f models.Follow
	err := context.ShouldBindJSON(&f)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")
	follow, err := models.GetFollowsByIDs(userId, f.FolloweeId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if follow.FollowerId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorised to delete follow"})
		return
	}

	err = follow.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Unfollowed successfully"})
}
