package handler

import (
	"fmt"
	"html/template"
	"net/http"
)

func Template(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error parsing")
	}
	err = ts.Execute(w, r)
	if err != nil {
		fmt.Println("Error executing template")
		return
	}
}
