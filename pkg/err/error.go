package _error

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

type ErrorReporter struct {
	Message    string
	StatusCode int
}

func (e ErrorReporter) Error() string {
	return e.Message
}

func NewErrorReport(code int, msg string) error {
	return &ErrorReporter{
		Message:    msg,
		StatusCode: code,
	}
}

func IsDuplicateEntryError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return true
	}
	return false
}

var (
	ErrConflict       = errors.New("resource_already_exist")
	ErrNoOrganization = errors.New("no_organization_associated")
	ErrNoUpdated      = errors.New("no_resource_to_update")
	ErrNoDeleted      = errors.New("no_resource_to_delete")
	ErrNotFound       = errors.New("resource_not_found")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("access_denied_to_this_resource")
	ErrInternal       = errors.New("internal_server_error")
)
