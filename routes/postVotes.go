package routes

import (
	"net/http"
	"strconv"

	"example.com/benjamin/models"
	"github.com/gin-gonic/gin"
)

func GetPostVoteValue(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalVotes, err := models.GetAllPostVoteValue(uint(postId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"totalVotes": totalVotes})
}

func GetPostVotes(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	votes, err := models.GetPostVotes(uint(postId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"votes": votes})
}

func GetAllPostVotes(context *gin.Context) {
	votes, err := models.GetAllPostVotes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"votes": votes})
}

func CreatePostVote(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")

	var vote models.PostVote
	err = context.ShouldBindJSON(&vote)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	vote.PostId = uint(postId)
	vote.UserId = uint(userId)

	err = vote.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Vote created", "postVote": vote})
}

func UpdatePostVote(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedVote models.PostVote
	err = context.ShouldBindJSON(&updatedVote)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")

	// check if user is the owner of the vote
	if updatedVote.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorised"})
		return
	}

	updatedVote.PostId = uint(postId)
	err = updatedVote.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Vote updated"})
}

func DeletePostVote(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")

	err = models.DeletePostVotes(userId, uint(postId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Vote deleted"})
}
