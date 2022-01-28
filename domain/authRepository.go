package domain

import (
	"database/sql"
	"log"

	"example.com/hexagonal-auth/errs"
	"github.com/jmoiron/sqlx"
)

type AuthRepository interface {
	FindBy(string, string) (*Login, *errs.AppError)
}

type AuthRepositoryDb struct {
	client *sqlx.DB
}

func (d AuthRepositoryDb) FindBy(username, password string) (*Login, *errs.AppError) {
	var login Login
	sqlVerify := `SELECT u.username, u.customer_id, u.role, group_concat(a.account_id) as account_numbers FROM users u
					LEFT JOIN accounts a ON a.customer_id = u.customer_id
					WHERE username = ? and password = ?
					GROUP BY u.customer_id`
	err := d.client.Get(&login, sqlVerify, username, password)
	if err != nil {
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errs.NewAuthenticationError("invalid credentials")
			} else {
				log.Println("error while verifying login request from database: " + err.Error())
				return nil, errs.NewUnexpectedError("unexpected database error")
			}
		}
	}
	return &login, nil
}

func NewAuthRepository(dbClient *sqlx.DB) AuthRepositoryDb {
	return AuthRepositoryDb{dbClient}
}
