package main

import (
	"net/http"
	"simple-douyin/controller"
	"simple-douyin/repository"

	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources
	r.Static("/videos", "./videos/")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.GET("/feed/", func(c *gin.Context) {
		latestTime := c.Query("latest_time")
		token := c.Query("token")
		data := controller.Feed(latestTime, token)
		c.JSON(http.StatusOK, data)
	})
	apiRouter.GET("/user/", func(c *gin.Context) {
		userId := c.Query("user_id")
		token := c.Query("token")
		data := controller.UserInfo(userId, token)
		c.JSON(http.StatusOK, data)
	})
	apiRouter.POST("/user/register/", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		data := controller.Register(username, password)
		c.JSON(http.StatusOK, data)
	})
	apiRouter.POST("/user/login/", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")
		data := controller.Login(username, password)
		c.JSON(http.StatusOK, data)
	})
	apiRouter.POST("/publish/action/", func(c *gin.Context) {
		response := controller.Publish(c)
		c.JSON(http.StatusOK, response)
	})
	apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.GET("/favorite/list/", controller.FavoriteList)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", controller.RelationAction)
	apiRouter.GET("/relation/follow/list/", controller.FollowList)
	apiRouter.GET("/relation/follower/list/", controller.FollowerList)

	// apiRouter.Static("/public/", "public/")
}

func main() {
	r := gin.Default()
	initRouter(r)
	repository.Init()
	r.Run() // listen and serve on 0.0.0.0:8080
}
