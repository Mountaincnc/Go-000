package biz

import (
	"week02/internal/week02/data"
)

func User(id int64) (*data.UserInfo, error) {
	return data.User(id)
}