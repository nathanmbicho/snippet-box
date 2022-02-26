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
	return 0, nil
}

// Get specific snipped by request id
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	return nil, nil
}

// Latest - most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
