package util

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func MysqlInit() *sql.DB {

	if Db != nil {
		err := Db.Ping()
		if err != nil {
			return nil
		}
		return Db
	}
	var err error
	Db, err = sql.Open("mysql", "root:123456@/netstore_user")
	if err != nil {
		return nil
	}
	err = Db.Ping()
	if err != nil {
		return nil
	}
	return Db

}
