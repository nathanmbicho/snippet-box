package mysql

import (
	"database/sql"
	"github.com/nathanmbicho/snippetbox/pkg/models"
)

// SnippetModel to wrap sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert new snippet into db
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// stmt sql execute statement
	stmt := "INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))"

	//execute statement and return sql.Result
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	//get lastInsertedId
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	//return id - convert it from int64 to int
	return int(id), nil
}

// Get specific snipped by request id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest - most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
