package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"path/filepath"
	"simple-douyin/service"
	"simple-douyin/util"
	"strconv"
	"time"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

type publishResponse struct {
	Response
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) *publishResponse {
	token := c.PostForm("token")
	print(token)
	// 查询用户Id，鉴权
	userId, err := util.ParseToken(token)
	if err != nil {
		return &publishResponse{
			Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		}
	}
	// 处理视频
	data, err := c.FormFile("data")
	if err != nil {
		return &publishResponse{
			Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
	}
	title := c.PostForm("title")
	rand.Seed(time.Now().UnixNano())
	filename := strconv.Itoa(rand.Intn(1000000)) + "_" + filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", userId, filename)
	saveFile := filepath.Join("./videos/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		return &publishResponse{
			Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		}
	}
	rootdir := "http://172.18.94.225:8080/videos/"
	service.AddVideo(userId, title, rootdir+finalName)
	return &publishResponse{
		Response{
			StatusCode: 0,
			StatusMsg:  finalName + " uploaded successfully",
		},
	}
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
