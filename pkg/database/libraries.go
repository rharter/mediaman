package database

import (
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Library table in the database
const libraryTable = "libraries"

// SQL Query to retrieve a library by it's unique database key
const libraryFindIdStmt = `
SELECT id, type, name, path, created, updated, last_scan
FROM libraries
WHERE id = ?
`

// SQL Query to retrieve all libraries
const libraryStmt = `
SELECT id, type, name, path, created, updated, last_scan
FROM libraries
`

// Returns a Library with the given ID.
func GetLibrary(id int64) (*Library, error) {
	library := Library{}
	err := meddler.QueryRow(db, &library, libraryFindIdStmt, id)
	return &library, err
}

// Saves a Library.
func SaveLibrary(library *Library) error {
	if library.Id == 0 {
		library.Created = time.Now().UTC()
	}
	library.Updated = time.Now().UTC()
	return meddler.Save(db, libraryTable, library)
}

// Deletes an existing Library.
func DeleteLibrary(id int64) error {
	db.Exec("DELETE FROM libraries WHERE id = ?", id)
	return nil
}

// Returns a list of all Libraries
func ListLibraries() ([]*Library, error) {
	var libraries []*Library
	err := meddler.QueryAll(db, &libraries, libraryStmt)
	return libraries, err
}
