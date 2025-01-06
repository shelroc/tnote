package repository

import (
	"fmt"

	"github.com/shelroc/tnote/src/utils"
)

var CreateTableQuery string = fmt.Sprintf(`
  CREATE TABLE IF NOT EXISTS %s (
    id integer PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
  )
`, utils.NotesTable)

var InsertQuery string = fmt.Sprintf(
	"INSERT INTO %s (title, content) VALUES (?, ?)",
	utils.NotesTable,
)

var SelectAllQuery string = fmt.Sprintf("SELECT * FROM %s", utils.NotesTable)

var SelectNoteQuery string = fmt.Sprintf(
	"SELECT * from %s WHERE id = ?",
	utils.NotesTable,
)

var UpdateQuery string = fmt.Sprintf(
	"UPDATE %s SET title = ?, content = ? WHERE id = ?",
	utils.NotesTable,
)

var DeleteQuery string = fmt.Sprintf(
	"DELETE FROM %s WHERE id = ?",
	utils.NotesTable,
)
