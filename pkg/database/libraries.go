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
SELECT id, type, name, root_id, created, updated, last_scan
FROM libraries
WHERE id = ?
`

// SQL Query to retrieve all libraries
const libraryStmt = `
SELECT id, type, name, root_id, created, updated, last_scan
FROM libraries
`

// Returns a Library with the given ID.
func GetLibrary(id int64) (*Library, error) {
	library := Library{}
	err := meddler.QueryRow(db, &library, libraryFindIdStmt, id)
	if err != nil {
		return &library, err
	}

	library.Root, err = GetElement(library.RootId)

	return &library, err
}

// Saves a Library.
func SaveLibrary(library *Library) error {
	err := SaveElement(library.Root)
	if err != nil {
		return err
	}
	library.RootId = library.Root.Id

	if library.Id == 0 {
		library.Created = time.Now().UTC()
	}
	library.Updated = time.Now().UTC()
	return meddler.Save(db, libraryTable, library)
}

// Deletes an existing Library.
func DeleteLibrary(id int64) error {
	lib, err := GetLibrary(id)
	if err != nil {
		return err
	}

	err = DeleteElement(lib.RootId)
	if err != nil {
		return err
	}

	db.Exec("DELETE FROM libraries WHERE id = ?", id)
	return nil
}

// Returns a list of all Libraries
func ListLibraries() ([]*Library, error) {
	var libraries []*Library
	err := meddler.QueryAll(db, &libraries, libraryStmt)

	for _, lib := range libraries {
		lib.Root, err = GetElement(lib.RootId)
		if err != nil {
			return libraries, err
		}
	}

	return libraries, err
}
