package main

import (
	"attractions/db"
	"attractions/router"
	"net/http"
)

func main() {

	db.InitDB("mysqld:mysql@tcp(45.141.79.120:3306)/Onviz")
	router.Router()
	http.ListenAndServe(":8080", nil)
}
