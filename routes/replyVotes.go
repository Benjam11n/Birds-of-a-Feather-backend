package routes

import (
	"net/http"
	"strconv"

	"github.com/Benjam11n/models"
	"github.com/gin-gonic/gin"
)

func GetReplyVotes(context *gin.Context) {
	replyId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalVotes, err := models.GetReplyVotes(uint(replyId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"totalVotes": totalVotes})
}

func GetAllReplyVotes(context *gin.Context) {
	totalVotes, err := models.GetAllReplyVotes()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"totalVotes": totalVotes})
}

func CreateReplyVote(context *gin.Context) {
	replyId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")

	var vote models.ReplyVote
	err = context.ShouldBindJSON(&vote)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vote.ReplyId = uint(replyId)
	vote.UserId = uint(userId)

	err = vote.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Vote created", "replyVote": vote})
}

func UpdateReplyVote(context *gin.Context) {
	var updatedVote models.ReplyVote
	err := context.ShouldBindJSON(&updatedVote)

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

	context.JSON(http.StatusOK, gin.H{"message": "Vote updated"})

	err = updatedVote.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Vote updated"})
}

func DeleteReplyVote(context *gin.Context) {
	replyId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")

	err = models.DeleteReplyVotes(userId, uint(replyId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Vote deleted"})
}
