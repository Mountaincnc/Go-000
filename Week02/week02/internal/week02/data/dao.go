package data

import (
	"database/sql"
	"fmt"
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
	querySql := fmt.Sprintf("select id, name from student where id = %v", id)
	if err := db.QueryRow(querySql).Scan(&info); err != nil {
		return nil, xerrors.Wrapf(sql.ErrNoRows,  fmt.Sprintf("querySql %s err: %v", querySql, err))
	}

	return info, nil
}
