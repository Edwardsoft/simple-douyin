package service

import (
	"fmt"
	"simple-douyin/repository"
	"time"
)

type VideoInfoFlow struct {
	NextTime  int64       `json:"next_time"`
	VideoList []VideoInfo `json:"video_list"`
}

type VideoInfo struct {
	Id            int64    `json:"id"`
	Author        UserInfo `json:"author"`
	PlayUrl       string   `json:"play_url"`
	CoverUrl      string   `json:"cover_url"`
	FavoriteCount int      `json:"favorite_count"`
	CommentCount  int      `json:"comment_count"`
	IsFavorite    bool     `json:"is_favorite"`
	Title         string   `json:"title"`
}

func Feed(latestTime time.Time, token string) (*VideoInfoFlow, error) {
	videos, err := repository.SelectVideoListByLatest(latestTime)
	if err != nil {
		return nil, fmt.Errorf("获取视频列表失败:%v", err.Error())
	}
	print("嘻嘻")
	videoInfoFlow := VideoInfoFlow{}
	if len(videos) == 0 {
		return &videoInfoFlow, nil
	}
	videoInfoFlow.NextTime = videos[len(videos)-1].CreateTime.Unix() * 1000
	videoInfoList := make([]VideoInfo, len(videos))
	for i := 0; i < len(videos); i++ {
		userInfo, err := GetUserInfo(videos[i].UserId, token)
		if err != nil {
			return nil, err
		}
		videoInfoList[i].Id = videos[i].Id
		videoInfoList[i].Author = *userInfo
		videoInfoList[i].PlayUrl = videos[i].PlayUrl
		videoInfoList[i].CoverUrl = videos[i].CoverUrl
		videoInfoList[i].Title = videos[i].Title
		videoInfoList[i].FavoriteCount = 0
		videoInfoList[i].CommentCount = 0
		videoInfoList[i].IsFavorite = false
	}
	videoInfoFlow.VideoList = videoInfoList
	return &videoInfoFlow, nil
}
