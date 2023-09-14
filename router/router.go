package router

import (
	"attractions/handler"
	"net/http"
)

func Router() {
	http.HandleFunc("/", handler.Template)
	http.HandleFunc("/get_city", handler.GetCities)
	http.HandleFunc("/get_travelers", handler.GetTravelers)
	http.HandleFunc("/get_sight", handler.GetSight)
	http.HandleFunc("/add_sight", handler.AddSightByTraveler)
}
