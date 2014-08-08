package model

import (
	"time"
)

type Library struct {
	Id   int64  `meddler:"id,pk"              json:"id"`
	Type string `meddler:"type"               json:"type"`
	Name string `meddler:"name"               json:"name"`

	RootId int64    `meddler:"root_id"  json:"-"`
	Root   *Element `meddler:"-"        json:"root"`

	Created  time.Time `meddler:"created,utctime"    json:"created"`
	Updated  time.Time `meddler:"updated,utctime"    json:"updated"`
	LastScan time.Time `meddler:"last_scan,utctime"  json:"last_scan"`
}

func NewLibrary(n, t, p string) *Library {
	library := Library{}
	library.Name = n
	library.Type = t
	library.Root = &Element{File: p, Type: "directory"}
	return &library
}
