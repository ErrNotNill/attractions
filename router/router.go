package router

import (
	"attractions/handler"
	"net/http"
)

func Router() {
	http.HandleFunc("/", handler.Template)
	http.HandleFunc("/set_rating", handler.SetRating)
	http.HandleFunc("/attr_by_traveler", handler.GetAttractionsByTraveller)
	http.HandleFunc("/add_attract", handler.AddAttraction)
	http.HandleFunc("/attracts", handler.GetAttractions)
	http.HandleFunc("/add_traveler", handler.AddTraveler)
	http.HandleFunc("/add_city", handler.AddCity)
}
