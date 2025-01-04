package repository

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/shelroc/tnote/src/models"
	"github.com/shelroc/tnote/src/utils"
)

var db *sql.DB

func Init() {
	var err error
	db, err = sql.Open("sqlite3", utils.DbPath)
	utils.HandleError(err)

	_, err = db.Exec(CreateTableQuery)
	utils.HandleError(err)
}

func Open() *sql.DB {
	return db
}

func GetNotes() []models.Note {
	rows, err := db.Query(SelectAllQuery)
	utils.HandleError(err)

	var notes []models.Note
	for rows.Next() {
		var note models.Note
		err = rows.Scan(
			&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt,
		)
		utils.HandleError(err)
		notes = append(notes, note)
	}

	return notes
}

func GetNote(id int) models.Note {
	var note models.Note
	err := db.QueryRow(SelectNoteQuery, id).Scan(
		&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}

	return note
}

func AddNote(note models.Note) models.Note {
	result, err := db.Exec(InsertQuery, note.Title, note.Content)
	utils.HandleError(err)

	id, _ := result.LastInsertId()
	note.Id = int(id)

	return note
}

func EditNote(id int, note models.Note) models.Note {
	_, err := db.Exec(UpdateQuery, note.Title, note.Content)
	utils.HandleError(err)

	return note
}

func DeleteNote(id int) {
	_, err := db.Exec(DeleteQuery, id)
	utils.HandleError(err)
}
