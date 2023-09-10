package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func InitDB(url string) (err error) {
	//var dataSourceName = "username:userlogin@tcp(ipaddress:port)/dbaseName"
	Db, err = sqlx.Connect("mysql", url)
	if err != nil {
		fmt.Println("not connected")
		return
	}
	err = Db.Ping()
	return
}
