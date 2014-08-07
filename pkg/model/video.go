package model

import (
	"time"
)

type Video struct {
	Id          int64  `meddler:"id,pk" json:"id"`
	File        string `meddler:"file" json:"file"`
	ParentId    int64  `meddler:"parent_id" json:"parent_id"`
	Title       string `meddler:"title" json:"title"`
	Description string `meddler:"description" json:"description"`
	Thumbnail   string `meddler:"thumbnail" json:"thumbnail"`
	Background  string `meddler:"background" json:"background"`
	Poster      string `meddler:"poster" json:"poster"`
	Banner      string `meddler:"banner" json:"banner"`

	Created time.Time `meddler:"created,localtime" json:"created"`
	Updated time.Time `meddler:"updated,localtime" json:"updated"`
}

func NewVideo(f string, p int64) *Video {
	video := Video{}
	video.File = f
	video.ParentId = p
	return &video
}
