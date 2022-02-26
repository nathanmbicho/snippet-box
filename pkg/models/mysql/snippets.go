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
	//select sql statement
	stmt := `SELECT id, title, content, created, expires FROM snippets 
			WHERE expires > UTC_TIMESTAMP() AND id = ?`

	//execute statement using QueryRow passing stmt, & id
	row := m.DB.QueryRow(stmt, id)

	// Initialize a pointer to a new zeroed Snippet struct.
	s := &models.Snippet{}

	//use row.Scan to return field sql.Row as our Snippet struct and user model.ErrorNoRecord if no record is found
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows { //if sql.Row error found
		return nil, models.ErrNoRecord
	} else if err != nil { //if any other error
		return nil, err
	}

	return s, nil
}

// Latest - most recently created snippets
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	return nil, nil
}
