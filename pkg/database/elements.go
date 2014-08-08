package database

import (
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Element table in the database
const elementTable = "elements"

// SQL Query to retrieve a element by it's unique database key
const elementFindIdStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated
FROM elements
WHERE id = ?
`

// SQL Query to retrieve a element by parent id
const elementFindParentStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated
FROM elements
WHERE parent_id = ?
`

// SQL Query to retrieve a element by filename
const elementFindFileStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated
FROM elements
WHERE file = ?
`

// SQL Query to retrieve all elements
const elementStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, created, updated
FROM elements
`

// Returns a element with a given Id.
func GetElement(id int64) (*Element, error) {
	element := Element{}
	err := meddler.QueryRow(db, &element, elementFindIdStmt, id)
	return &element, err
}

// Returns all elements belonging to parent id
func GetElementsForParent(id int64) ([]*Element, error) {
	var elements []*Element
	err := meddler.QueryAll(db, &elements, elementFindParentStmt, id)
	return elements, err
}

// Returns a element with a given filename.
func GetElementByFile(f string) (*Element, error) {
	element := Element{}
	err := meddler.QueryRow(db, &element, elementFindFileStmt, f)
	return &element, err
}

// Saves a Element.
func SaveElement(element *Element) error {
	if element.Id == 0 {
		element.Created = time.Now().UTC()
	}
	element.Updated = time.Now().UTC()
	return meddler.Save(db, elementTable, element)
}

// Deletes an existing Element.
func DeleteElement(id int64) error {
	db.Exec("DELETE FROM elements WHERE id = ?", id)
	return nil
}

// Returns a list of all Elements
func ListElements() ([]*Element, error) {
	var elements []*Element
	err := meddler.QueryAll(db, &elements, elementStmt)
	return elements, err
}
