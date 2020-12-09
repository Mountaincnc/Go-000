package biz

import (
	"database/sql"
	"errors"
	"week02/internal/week02/data"
)

func User(id int64) (*data.UserInfo, error) {
	user, err := data.User(id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}
	return user, err
}