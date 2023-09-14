package handler

import (
	"attractions/repo"
	"encoding/json"
	"net/http"
)

func GetCities(w http.ResponseWriter, r *http.Request) {
	city := repo.GetCity()
	for _, v := range city {
		//fmt.Fprint(w, v.CityID)
		//fmt.Fprint(w, v.Name)
		json.Marshal(v)
		json.NewEncoder(w).Encode(v)
		//json.NewEncoder(w).Encode(v.CityID)
		//json.NewEncoder(w).Encode(v.Name)
	}
	//json.MarshalIndent(city, "", ",")
	//fmt.Fprint(w, city)
}
func GetTravelers(w http.ResponseWriter, r *http.Request) {
	travelers := repo.GetTraveler()
	for _, v := range travelers {
		json.Marshal(v)
		json.NewEncoder(w).Encode(v)
	}
}
func GetSight(w http.ResponseWriter, r *http.Request) {
	sights := repo.GetSight()
	for _, v := range sights {
		json.Marshal(v)
		json.NewEncoder(w).Encode(v)
	}
}
func AddSightByTraveler(w http.ResponseWriter, r *http.Request) {
	sights := repo.GetSight()
	for _, v := range sights {
		json.Marshal(v)
		json.NewEncoder(w).Encode(v)
	}
}

/*func GetAttraction(w http.ResponseWriter, r *http.Request) {
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
}*/
