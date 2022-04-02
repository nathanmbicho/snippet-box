package mysql

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/nathanmbicho/snippetbox/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type UserModel struct {
	DB *sql.DB
}

//Insert - add new user to the database users tbl
func (m *UserModel) Insert(name, email, password string) error {
	//bcrypt hash the plain password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created) VALUES(?, ?, ?, UTC_TIMESTAMP())`

	/**
	use Exec() method to insert users details & hashed password in users table.
	use mysql Errors to check error number and if it related to our email unique key by checking message sent and if true return ErrDuplicateEmail else
	original error
	*/
	_, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 && strings.Contains(mysqlErr.Message, "users_uc_email") {
				return models.ErrDuplicateEmail
			}
		}
	}
	return err
}

//Authenticate - verify users credentials if exists with email and password provided and return user ID if found
func (m *UserModel) Authenticate(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	//retrieve id and hashed_password from the email given and check if matching email exists then throw ErrInValidCredentials
	row := m.DB.QueryRow("SELECT id, hashed_password FROM users WHERE email = ?", email)
	err := row.Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	//hashed & password if matches else return ErrInValidCredentials
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}
	//Return user id if validated
	return id, nil
}

//Get - fetch specific user details using their ID
func (m *UserModel) Get(id int) (*models.User, error) {
	return nil, nil
}
