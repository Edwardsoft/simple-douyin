package util

import (
	"fmt"
	"github.com/golang-jwt/jwt"
)

const (
	TOKEN_KEY = "http://simple-douyin.com"
)

func GenToken(userId int64, expSec int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    expSec,
		"userId": userId,
	})
	tokenString, err := token.SignedString([]byte(TOKEN_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(TOKEN_KEY), nil
	})
	if err != nil {
		return -1, err
	}
	claims := token.Claims.(jwt.MapClaims)
	err = claims.Valid()
	if err != nil {
		return -1, fmt.Errorf("invalid token")
	}

	return int64(claims["userId"].(float64)), nil
}

/**
// GetVideoCover 生成视频缩略图并保存（作为封面）
func GetVideoCover(videoPath, coverPath string, frameNum int) (coverName string) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	err = imaging.Save(img, coverPath+".jpeg")
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
	}

	// 成功则返回生成的缩略图名
	names := strings.Split(coverPath, "\")
	coverName = names[len(names)-1] + ".jpeg"
	return
}

*/
