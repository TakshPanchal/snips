package models

import (
	"database/sql"
	"errors"
)

type SnippetModel struct {
	DB *sql.DB
}

type Snippet struct {
	Id               int
	Title, Content   string
	Created, Expires string
}

func (s *SnippetModel) Insert(snip Snippet) (int, error) {
	query := `INSERT INTO snips (title, content, created, expires)
	VALUES($1, $2, now(), now() + INTERVAL '1 DAY' * $3) RETURNING id;`

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	id := 0

	// TODO: Change the hardcoded value here
	err = stmt.QueryRow(snip.Title, snip.Content, '7').Scan(&id)
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get queries a snippet for id from the database
func (s *SnippetModel) Get(id int) (*Snippet, error) {
	query := `SELECT id, title, content, created, expires FROM snips
		WHERE expires > NOW() AND id = $1`

	snip := &Snippet{}

	err := s.DB.QueryRow(query, id).Scan(&snip.Id, &snip.Title, &snip.Content,
		&snip.Created, &snip.Expires)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		}

		return nil, err
	}

	return snip, nil
}
