package auth

import (
	"fmt"
	"net/http"
)

func CreateLogin(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("username")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	for qry.Next() {
		if err := qry.Scan(&count); err != nil {
			fmt.Println(`error scan`)
			//log.Fatal(err)
		}
	}
}
func CreatePassword(w http.ResponseWriter, r *http.Request) {
	pass := r.FormValue("password")

}
