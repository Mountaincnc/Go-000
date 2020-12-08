package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	xerrors "github.com/pkg/errors"
)

type UserInfo struct {
	Id       int64 `sql:"id"`
	Username string `sql:"username"`
}

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/userDB?charset=utf8")
	if err != nil {
		panic(err)
	}
}

func User(id int64) (*UserInfo, error) {
	var info *UserInfo
	if err := db.QueryRow(`select id, name from student where id = ?`, id).Scan(&info); err != nil {
		return nil, xerrors.Wrapf(err, "get user failed id = %d", id)
	}

	return info, nil
}
