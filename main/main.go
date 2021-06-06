package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)
func main()  {
	db, err := sql.Open("mysql", "root:root@(127.0.0.1:3306)/test");
	if err != nil {
		return "", errors.Wrap(err, "db open error")
	}
	var userName string
	row := db.QueryRow("select user_name from user where id = 1");
	err1 := row.Scan(&userName)

	if err1 != sql.ErrNoRows {
		return "", errors.Wrap(err1, "db open error")
	}
	return userName, nil
}
