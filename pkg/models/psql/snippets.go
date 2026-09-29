package psql

import (
	"database/sql"
	"errors"

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
	stmt := `SELECT id, title,content,created,expires FROM snippets WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	row := m.DB.QueryRow(stmt, id)

	//initialize a pointer to a new zeroed snippt struct
	s := &models.Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil

}

func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id,title,content,created,expires FROM snippets WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	// We defer rows.Close() to ensure the sql.Rows resultset is
	// always properly closed before the Latest() method returns. This defer
	// statement should come *after* you check for an error from the Query()
	// method. Otherwise, if Query() returns an error, you'll get a panic
	// trying to close a nil resultset.

	defer rows.Close()

	snippets := []*models.Snippet{}

	for rows.Next() {
		s := &models.Snippet{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// append it to the snippet of slice
		snippets = append(snippets, s)
	}
	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration. It's important to
	// call this - don't assume that a successful iteration was completed
	// over the whole resultset.

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
