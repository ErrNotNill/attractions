package handler

import (
	"attractions/repo"
	"encoding/json"
	"net/http"
)

func GetAttraction(w http.ResponseWriter, r *http.Request) {
	//todo get attraction rating by id
	//mediumRating :=
}

func GetAttractions(w http.ResponseWriter, r *http.Request) {
	attr := repo.Attraction{}
	attractions := attr.GetAllAttractions()
	w.Write(attractions)
}

func GetAttractionsByTraveller(w http.ResponseWriter, r *http.Request) {
	attractions := make([]repo.Attraction, 0)
	travelerName := r.FormValue("traveler")
	traveler := repo.Traveler{Name: travelerName}
	Attractions, err := repo.GetAttractionsByTraveler(traveler)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for _, v := range Attractions {
		attractions = append(attractions, v)
	}
	data, err := json.MarshalIndent(&attractions, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Write(data)
}
