package routes

import (
	"net/http"
	"strconv"

	"github.com/Benjam11n/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetCommunities(context *gin.Context) {
	communities, err := models.GetAllCommunities()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, communities)
}

func GetCommunity(context *gin.Context) {
	communityId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	community, err := models.GetCommunityById(uint(communityId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, community)
}

func CreateCommunity(context *gin.Context) {
	var community models.Community
	err := context.ShouldBindWith(&community, binding.FormMultipart)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the file is uploaded
	file, err := context.FormFile("iconUrl")
	var avatarFilename string

	if err != nil {
		// No file uploaded, set icon to default
		community.IconUrl = "/assets/communities/default.jpg"
	} else {
		// File uploaded, proceed with saving and updating
		avatarFilename = file.Filename
		community.IconUrl = "/assets/communities/" + avatarFilename

		err = context.SaveUploadedFile(file, "assets/communities/"+avatarFilename)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
			return
		}
	}

	userId := context.GetUint("userId")
	community.UserId = userId

	err = community.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "community created successfully", "community": community})
}

func UpdateCommunity(context *gin.Context) {
	communityId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userId := context.GetUint("userId")
	community, err := models.GetCommunityById(uint(communityId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "community not found"})
		return
	}

	// check if user is the owner of the community
	if community.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorised to update community"})
		return
	}

	var updatedCommunity models.Community
	err = context.ShouldBindJSON(&updatedCommunity)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updatedCommunity.ID = uint(communityId)

	err = updatedCommunity.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "community not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Community updated successfully"})
}

func DeleteCommunity(context *gin.Context) {
	communityId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")
	community, err := models.GetCommunityById(uint(communityId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "community not found"})
		return
	}

	// check if user is the owner of the community
	if community.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorised to update community"})
		return
	}

	err = community.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Community deleted successfully"})
}
