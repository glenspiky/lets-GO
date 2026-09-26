package psql

import (
	"database/sql"

	"github.com/glenspiky/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// this will insert a new snippet into the db
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets(title,content, created,expires)
	VALUES($1,$2,CURRENT_TIMESTAMP,CURRENT_TIMESTAMP + ($3*INTERVAL '1 day'))
	RETURNING id`

	var id int

	err := m.DB.QueryRow(stmt, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	//Get the id of our newly inserted record
	return int(id), nil
}
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil

}
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
