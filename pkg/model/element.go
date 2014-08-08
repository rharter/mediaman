package model

import (
	"time"
)

type Element struct {
	Id          int64  `meddler:"id,pk" json:"id"`
	Type        string `meddler:"type" json:"type"`
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

func NewElement(f string, p int64, t string) *Element {
	element := Element{}
	element.File = f
	element.ParentId = p
	element.Type = t
	return &element
}
