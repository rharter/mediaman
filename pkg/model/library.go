package model

import (
	"time"
)

type Library struct {
	ID         int64     `meddler:"id,pk"              json:"id"`
	Name       string    `meddler:"name"               json:"name"`
	Path       string    `meddler:"path"               json:"path"`

	Created    time.Time `meddler:"created,utctime"    json:"created"`
	Updated    time.Time `meddler:"updated,utctime"    json:"updated"`
	LastScan   time.Time `meddler:"last_scan,utctime"  json:"last_scan"`
}

func NewLibrary(name, path string) *Library {
	library := Library{}
	library.Name = name
	library.Path = path
	return &library
}