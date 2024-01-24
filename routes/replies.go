package routes

import (
	"net/http"
	"strconv"

	"github.com/Benjam11n/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetAllReplies(context *gin.Context) {
	replies, err := models.GetAllReplies()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, replies)
}

func GetReplies(context *gin.Context) {
	parentId, err := strconv.ParseUint(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	replies, err := models.GetAllPostReplies(uint(parentId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, replies)
}

func GetReply(context *gin.Context) {
	replyId, err := strconv.ParseUint(context.Param("replyid"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reply, err := models.GetReplyByID(uint(replyId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, reply)
}

func CreateReply(context *gin.Context) {
	parentID, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var reply models.Reply
	err = context.ShouldBindWith(&reply, binding.FormMultipart)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the file is uploaded
	file, err := context.FormFile("imagesUrl")
	var avatarFilename string

	if err != nil {
		// No file uploaded, continue
	} else {
		// File uploaded, proceed with saving and updating
		avatarFilename = file.Filename
		reply.ImagesUrl = "/assets/replies/" + avatarFilename

		err = context.SaveUploadedFile(file, "assets/replies/"+avatarFilename)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
			return
		}
	}

	userId := context.GetUint("userId")
	reply.UserId = userId
	reply.ParentId = uint(parentID)

	err = reply.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "reply created successfully", "reply": reply})
}

func UpdateReply(context *gin.Context) {
	replyId, err := strconv.ParseUint(context.Param("replyid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	parentID, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")
	reply, err := models.GetReplyByID(uint(replyId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "reply not found"})
		return
	}
	// check if user is the owner of the reply
	if reply.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorised to update reply"})
		return
	}

	var updatedReply models.Reply
	err = context.ShouldBindJSON(&updatedReply)
	updatedReply.ParentId = uint(parentID)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedReply.ID = uint(replyId)
	err = updatedReply.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Reply updated succcessfully"})
}

func DeleteReply(context *gin.Context) {
	replyId, err := strconv.ParseUint(context.Param("replyid"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get userID from context
	userId := context.GetUint("userId")
	reply, err := models.GetReplyByID(uint(replyId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// check if user is the owner of the reply
	if reply.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorised to delete reply"})
		return
	}
	// delete reply
	err = reply.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Reply deleted succcessfully"})
}
