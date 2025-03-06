package mysql

import (
	"database/sql"
	"errors"
	"strings"
)

var uniqueErrorKeywords = map[string]struct{}{ //nolint:gochecknoglobals // nothing
	"duplicate key":     {},
	"unique constraint": {},
	"UNIQUE constraint": {},
	"Duplicate entry":   {},
}

func CheckUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}

	errMsg := err.Error()
	for keyword := range uniqueErrorKeywords {
		if strings.Contains(errMsg, keyword) {
			return true
		}
	}

	return false
}

func IsNoRowsError(err error) bool {
	return errors.Is(err, sql.ErrNoRows)
}
