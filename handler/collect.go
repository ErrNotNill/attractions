package handler

/*func AddCity(w http.ResponseWriter, r *http.Request) {
	city := &repo.City{}
	title := r.FormValue("title")
	city.AddCity(title)
	fmt.Println("title, rating: ", title)
}

func AddTraveler(w http.ResponseWriter, r *http.Request) {
	traveler := &repo.Traveler{}
	year, _ := strconv.Atoi(r.FormValue("year"))
	day, _ := strconv.Atoi(r.FormValue("day"))
	month, _ := strconv.Atoi(r.FormValue("month"))
	hour, _ := strconv.Atoi(r.FormValue("hour"))
	minute, _ := strconv.Atoi(r.FormValue("minute"))
	second, _ := strconv.Atoi(r.FormValue("second"))
	traveler.Name = r.FormValue("traveler")
	fmt.Println("traveler.Name", traveler.Name)
	date := repo.Date{Year: year, Month: month, Day: day, Hour: hour, Minute: minute, Second: second}
	traveler.AddTraveler(traveler.Name, date)
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error read body")
	}
	fmt.Println("BODY: ", string(body))
}

func AddAttraction(w http.ResponseWriter, r *http.Request) {
	attraction := &repo.Attraction{}
	traveler := &repo.Traveler{}
	if r.Method == "POST" {
		attraction.Title = r.FormValue("title")
		traveler.Name = r.FormValue("traveler")
		attraction.City.Title = r.FormValue("city")
		attraction.RangeFromCenter = r.FormValue("range")
		attraction.Rating = r.FormValue("rating")
		attraction.AddAttraction()
		fmt.Println(attraction.Title)
		fmt.Println("attraction", attraction)
	}
}

func SetRating(w http.ResponseWriter, r *http.Request) {
	//traveler := r.FormValue("traveler_name")
	//setRating := r.FormValue("set_rating")
	//SetRatingForAttraction()
}*/
