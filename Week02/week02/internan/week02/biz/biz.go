package biz

import (
	"week02/internan/week02/data"
)

func User(id int64) (*data.UserInfo, error) {
	return data.User(id)
}