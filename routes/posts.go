package routes

import (
	"fmt"
	"net/http"
	"strconv"

	"example.com/benjamin/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func GetPosts(context *gin.Context) {
	posts, err := models.GetAllPosts()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, posts)
}

func GetPost(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post, err := models.GetPostByID(uint(postId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, post)
}

func CreatePost(context *gin.Context) {
	var post models.Post
	err := context.ShouldBindWith(&post, binding.FormMultipart)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println(err.Error())
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
		post.ImagesUrl = "/assets/posts/" + avatarFilename

		err = context.SaveUploadedFile(file, "assets/posts/"+avatarFilename)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save image"})
			return
		}
	}

	userId := context.GetUint("userId")
	post.UserId = userId

	err = post.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "post created successfully", "post": post})
}

func UpdatePost(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := context.GetUint("userId")
	post, err := models.GetPostByID(uint(postId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "post not found"})
		return
	}

	// check if user is the owner of the post
	if post.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorised to update post"})
		return
	}

	var updatedPost models.Post
	err = context.ShouldBindJSON(&updatedPost)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedPost.ID = uint(postId)
	err = updatedPost.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Post updated succcessfully"})
}

func DeletePost(context *gin.Context) {
	postId, err := strconv.ParseUint(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get userID from context
	userId := context.GetUint("userId")
	post, err := models.GetPostByID(uint(postId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// check if user is the owner of the post
	if post.UserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorised to delete post"})
		return
	}

	postReplies, err := models.GetAllPostReplies(uint(postId))

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Delete all replies associated with the post
	for _, reply := range postReplies {
		if err := reply.Delete(); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete a reply"})
			return
		}
	}

	// delete post
	err = post.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Post deleted succcessfully"})
}
