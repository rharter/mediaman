package database

import (
	"database/sql"
	"time"

	. "github.com/rharter/mediaman/pkg/model"
	"github.com/russross/meddler"
)

// Name of the Element table in the database
const elementTable = "elements"

// SQL Query to retrieve a element by it's unique database key
const elementFindIdStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, remote_id, created, updated
FROM elements
WHERE id = ?
`

// SQL Query to retrieve a element by it's title
const elementFindTitleStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, remote_id, created, updated
FROM elements
WHERE title = ?
`

// SQL Query to retrieve a element by parent id
const elementFindParentStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, remote_id, created, updated
FROM elements
WHERE parent_id = ?
`

// SQL Query to retrieve a element by filename
const elementFindFileStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, remote_id, created, updated
FROM elements
WHERE file = ?
`

// SQL Query to retrieve all elements
const elementStmt = `
SELECT id, file, parent_id, title, description, thumbnail, background, poster, banner, remote_id, created, updated
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

// Returns an element with a given name.
func GetElementByTitle(t string) (*Element, error) {
	element := Element{}
	err := meddler.QueryRow(db, &element, elementFindTitleStmt, t)
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

// Deletes an existing Element, and all of it's descendents.
func DeleteElementCascade(id int64) error {
	err := DeleteElement(id)
	if err != nil {
		return err
	}
	es, err := GetElementsForParent(id)
	if err != nil && err != sql.ErrNoRows {
		return err
	}
	for _, e := range es {
		err = DeleteElementCascade(e.Id)
		if err != nil {
			return err
		}
	}
	return nil
}

// Returns a list of all Elements
func ListElements() ([]*Element, error) {
	var elements []*Element
	err := meddler.QueryAll(db, &elements, elementStmt)
	return elements, err
}
