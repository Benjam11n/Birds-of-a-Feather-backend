package routes

import (
	"net/http"
	"strconv"

	"example.com/benjamin/models"
	"github.com/gin-gonic/gin"
)

func GetCommunityMembers(context *gin.Context) {
	communityId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	communityMembers, err := models.GetCommunityMembers(uint(communityId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, communityMembers)
}

func GetAllCommunityMembers(context *gin.Context) {
	communityMembers, err := models.GetAllCommunityMembers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, communityMembers)
}

func GetAllUserCommunities(context *gin.Context) {
	userId := context.GetUint("userId")
	communityMembers, err := models.GetAllUserCommunities(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, communityMembers)
}

func CreateCommunityMember(context *gin.Context) {
	var member models.CommunityMember

	userId := context.GetUint("userId")
	communityId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	member.CommunityId = uint(communityId)
	member.UserId = userId

	err = member.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "community member created successfully", "community": member})
}

func DeleteCommunityMember(context *gin.Context) {
	var member models.CommunityMember
	communityId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")
	member.CommunityId = uint(communityId)
	member.UserId = userId

	err = member.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Community member deleted succcessfully"})
}
