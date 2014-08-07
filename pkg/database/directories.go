package database

import (
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Directory table in the database
const directoryTable = "directories"

// SQL Query to retrieve a directory by it's unique database key
const directoryFindIdStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated, SELECT id, type, name, path, created, updated, last_scan
FROM directories
WHERE id = ?
`

// SQL Query to retrieve a directory by parent directory
const directoryFindParentStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated, SELECT id, type, name, path, created, updated, last_scan
FROM directories
WHERE parent_id = ?
`

// SQL Query to retrieve a directory by filename
const directoryFindFileStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated, SELECT id, type, name, path, created, updated, last_scan
FROM directories
WHERE file = ?
`

// SQL Query to retrieve all directories
const directoryStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated, SELECT id, type, name, path, created, updated, last_scan
FROM directories
`

// Returns a directory with a given Id.
func GetDirectory(id int64) (*Directory, error) {
	directory := Directory{}
	err := meddler.QueryRow(db, &directory, directoryFindIdStmt, id)
	return &directory, err
}

// Returns a list of directories with the specified parent
func GetDirectoriesForParent(id int64) ([]*Directory, error) {
	var directories []*Directory
	err := meddler.QueryAll(db, &directories, directoryFindParentStmt, id)
	return directories, err
}

// Returns a directory with a given filename.
func GetDirectoryByFile(f string) (*Directory, error) {
	directory := Directory{}
	err := meddler.QueryRow(db, &directory, directoryFindFileStmt, f)
	return &directory, err
}

// Saves a Directory.
func SaveDirectory(directory *Directory) error {
	if directory.Id == 0 {
		directory.Created = time.Now().UTC()
	}
	directory.Updated = time.Now().UTC()
	return meddler.Save(db, directoryTable, directory)
}

// Deletes an existing Directory.
func DeleteDirectory(id int64) error {
	db.Exec("DELETE FROM directories WHERE id = ?", id)
	return nil
}

// Returns a list of all Directories
func ListDirectories() ([]*Directory, error) {
	var directories []*Directory
	err := meddler.QueryAll(db, &directories, directoryStmt)
	return directories, err
}
