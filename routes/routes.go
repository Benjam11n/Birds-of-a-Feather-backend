package routes

import (
	"github.com/Benjam11n/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.Static("/assets", "./assets")
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	// Serve static files from the "assets" folder

	// posts
	server.GET("/posts", GetPosts)
	server.GET("/posts/:id", GetPost)
	authenticated.POST("/posts", CreatePost)
	authenticated.PUT("/posts/:id", UpdatePost)
	authenticated.DELETE("/posts/:id", DeletePost)

	// replies to posts
	server.GET("/replies", GetAllReplies)
	server.GET("/posts/:id/replies", GetReplies)
	server.GET("/posts/:id/replies/:replyid", GetReply)
	authenticated.POST("/posts/:id/replies", CreateReply)
	authenticated.PUT("/posts/:id/replies/:replyid", UpdateReply)
	authenticated.DELETE("/posts/:id/replies/:replyid", DeleteReply)

	// users
	server.GET("/users", GetUsers)
	authenticated.GET("/", GetCurrentUser)
	server.GET("/users/:id", GetUser)
	authenticated.PUT("/users/:id", UpdateUser)
	authenticated.PUT("/users/:id/password", UpdateUserPassword)
	authenticated.PUT("/users/:id/avatar", UploadAvatar)
	authenticated.DELETE("/users/:id", DeleteUser)
	server.POST("/signup", Signup)
	server.POST("/login", Login)
	authenticated.POST("/logout", Logout)

	// follows
	authenticated.GET("/follows", GetFollowees)
	// server.GET("/follows", GetFollows)
	authenticated.POST("/follows", CreateFollow)
	authenticated.DELETE("/follows", DeleteFollow)

	// postVotes
	server.GET("/posts/:id/votes", GetPostVotes)
	server.GET("/postvotes", GetAllPostVotes)
	authenticated.POST("/posts/:id/votes", CreatePostVote)
	authenticated.PUT("/posts/:id/votes", UpdatePostVote)
	authenticated.DELETE("/posts/:id/votes", DeletePostVote)

	// replyVotes
	server.GET("replies/:id/votes", GetReplyVotes)
	server.GET("replyvotes", GetAllReplyVotes)
	authenticated.POST("replies/:id/votes", CreateReplyVote)
	authenticated.PUT("replies/:id/votes", UpdateReplyVote)
	authenticated.DELETE("replies/:id/votes", DeleteReplyVote)

	//communities
	server.GET("communities", GetCommunities)
	server.GET("communities/:id", GetCommunity)
	authenticated.GET("userCommunities", GetAllUserCommunities)
	authenticated.POST("communities", CreateCommunity)
	authenticated.PUT("communities/:id", UpdateCommunity)
	authenticated.DELETE("communities/:id", DeleteCommunity)

	//communityMembers
	server.GET("communitymembers", GetAllCommunityMembers)
	server.GET("communities/:id/members", GetCommunityMembers)
	authenticated.POST("communities/:id/members", CreateCommunityMember)
	authenticated.DELETE("communities/:id/members", DeleteCommunityMember)
}
