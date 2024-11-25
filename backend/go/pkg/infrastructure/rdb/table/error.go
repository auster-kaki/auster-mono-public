package table

import (
	"database/sql"
	"errors"

	"github.com/auster-kaki/auster-mono/pkg/app/repository"

	"github.com/go-sql-driver/mysql"
)

func handleError(err error) error {
	if errors.Is(err, sql.ErrNoRows) {
		return repository.ErrNotFound
	}
	if e := new(mysql.MySQLError); errors.As(err, &e) {
		switch e.Number {
		case 1062:
			return repository.ErrDuplicate
		case 1146:
			return repository.ErrNotExists
		case 1050:
			return repository.ErrAlreadyExists
		}
	}
	return err
}
