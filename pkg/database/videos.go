package database

import (
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Video table in the database
const videoTable = "videos"

// SQL Query to retrieve a video by it's unique database key
const videoFindIdStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated
FROM videos
WHERE id = ?
`

// SQL Query to retrieve a video by parent id
const videoFindParentStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated
FROM videos
WHERE parent_id = ?
`

// SQL Query to retrieve a video by filename
const videoFindFileStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated
FROM videos
WHERE file = ?
`

// SQL Query to retrieve all videos
const videoStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated
FROM videos
`

// Returns a video with a given Id.
func GetVideo(id int64) (*Video, error) {
	video := Video{}
	err := meddler.QueryRow(db, &video, videoFindIdStmt, id)
	return &video, err
}

// Returns all videos belonging to parent id
func GetVideosForParent(id int64) ([]*Video, error) {
	var videos []*Video
	err := meddler.QueryAll(db, &videos, videoFindParentStmt, id)
	return videos, err
}

// Returns a video with a given filename.
func GetVideoByFile(f string) (*Video, error) {
	video := Video{}
	err := meddler.QueryRow(db, &video, videoFindFileStmt, f)
	return &video, err
}

// Saves a Video.
func SaveVideo(video *Video) error {
	if video.Id == 0 {
		video.Created = time.Now().UTC()
	}
	video.Updated = time.Now().UTC()
	return meddler.Save(db, videoTable, video)
}

// Deletes an existing Video.
func DeleteVideo(id int64) error {
	db.Exec("DELETE FROM videos WHERE id = ?", id)
	return nil
}

// Returns a list of all Videos
func ListVideos() ([]*Video, error) {
	var videos []*Video
	err := meddler.QueryAll(db, &videos, videoStmt)
	return videos, err
}
