package mysql_utils

import (
	"fmt"
	"strings"

	"github.com/devBOX03/bookstore_user_api/utils/errors"
	"github.com/go-sql-driver/mysql"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestError {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewBadRequestError("No records matching given id")
		}
		return errors.NewInternalServerError(fmt.Sprintf("error parsing database resopnse due to %s", err.Error()))
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("Invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
