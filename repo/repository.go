package repo

import (
	"attractions/db"
	"fmt"
)

func GetCity() []City {
	//return GetCity
	//return TravelerID //todo test data for travelers
	qry, err := db.Db.Query(`SELECT CityID, Name
       FROM City`)
	if err != nil {
		fmt.Println(`error query`)
	}
	city := City{}
	var cities []City
	for qry.Next() {
		if err := qry.Scan(&city.CityID, &city.Name); err != nil {
			fmt.Println(`error scan`)
		}
		cities = append(cities, city)
		fmt.Println("city.CityID: ", city.CityID)
		fmt.Println("city.Title: ", city.Name)
	}
	return cities
}

func GetTraveler() []Traveler {
	//return TravelerID //todo test data for travelers
	qry, err := db.Db.Query(`SELECT TravelerID, Name
       FROM Traveler`)
	if err != nil {
		fmt.Println(`error query`)
	}
	traveler := Traveler{}
	var travelers []Traveler
	for qry.Next() {
		if err := qry.Scan(&traveler.TravelerID, &traveler.Name); err != nil {
			fmt.Println(`error scan`)
		}
		travelers = append(travelers, traveler)
		fmt.Println("traveler.ID: ", traveler.TravelerID)
		fmt.Println("traveler.Name: ", traveler.Name)
	}
	return travelers
}

func GetSight() []Sight {
	//return TravelerID //todo test data for travelers
	qry, err := db.Db.Query(`SELECT SightID, Title,Distance,CityID
       FROM Sight LEFT JOIN City ON city.CityID = city.CityID`)
	if err != nil {
		fmt.Println(`error query`)
	}
	sight := Sight{}
	var sights []Sight
	for qry.Next() {
		if err := qry.Scan(&sight.SightID, &sight.Title, &sight.Distance); err != nil {
			fmt.Println(`error scan`)
		}
		sights = append(sights, sight)
		fmt.Println("sight.SightID: ", sight.SightID)
		fmt.Println("sight.Title: ", sight.Title)

	}
	return sights
}

/*func (a Attraction) GetAttractionsByCity() []Attraction {
	qry, err := db.Db.Query(`SELECT
    Attraction.Title
FROM
    Attraction
        LEFT JOIN
    City ON City.ID = Attraction.ID`)
	if err != nil {
		fmt.Println(`error query`)
	}
	attraction := Attraction{}
	for qry.Next() {
		if err := qry.Scan(&attraction.ID, &attraction.Traveler,
			&attraction.City.ID, &attraction.Title, &attraction.RangeFromCenter,
			&attraction.Rating,
		); err != nil {
			fmt.Println(`error scan`)
		}
		fmt.Println("attractions: ", attraction)
	}
	return nil
}

func (t Traveler) getTraveler() Traveler {
	traveler := Traveler{}
	qry, err := db.Db.Query(`SELECT ID
       FROM Traveler WHERE ID = 2`)
	if err != nil {
		fmt.Println(`error query`)
	}
	for qry.Next() {
		if err := qry.Scan(&traveler.ID); err != nil {
			fmt.Println(`error scan`)
		}
	}
	fmt.Println("traveler.get: ", traveler)
	return traveler
}

func (a Attraction) getAllAttractionsByCity() Attraction {
	qry, err := db.Db.Query(`SELECT * FROM Attraction
INNER JOIN
    City ON City.Title = 'Kirov';`)
	if err != nil {
		fmt.Println(`error query`)
	}
	attraction := Attraction{}
	for qry.Next() {
		if err := qry.Scan(&attraction); err != nil {
			fmt.Println(`error scan`)
		}
		fmt.Println("attractions.City.ID: ", attraction.City.ID)
		fmt.Println("attractions by city: ", attraction)
	}
	return attraction
}

func (a Attraction) GetAllAttractions() []byte {
	//traveler_id,city_id,Title,RangeFromCenter,Description,Rating
	//fmt.Println("a.getAllAttractionsByCity()", a.getAllAttractionsByCity())
	qry, err := db.Db.Query(`SELECT ID,traveler_id,city_id,Title,RangeFromCenter,
       Rating
       FROM Attraction`)
	if err != nil {
		fmt.Println(`error query`)
	}
	attraction := Attraction{}
	var attractions []Attraction
	for qry.Next() {
		if err := qry.Scan(&attraction.ID, &attraction.Traveler.ID, &attraction.City.ID,
			&attraction.Title, &attraction.RangeFromCenter, &attraction.Rating,
		); err != nil {
			fmt.Println(`error scan`)
		}
		attractions = append(attractions, attraction)
		fmt.Println("attractions.City.ID: ", attraction.City.ID)
		fmt.Println("attractions: ", attraction)
	}
	js, _ := json.MarshalIndent(&attractions, "", "  ")
	return js
}

func GetRatingFromAttraction(attraction string) int {
	//`SELECT Rating FROM Attractions WHERE ID = ' '`
	attractions := Attraction{}
	var rating int
	qry, err := db.Db.Query(`SELECT AVG(RATING(*)) FROM Attractions WHERE Attraction = $`, attraction)
	if err != nil {
		fmt.Println(`error query`)
	}
	for qry.Next() {
		if err := qry.Scan(&attractions); err != nil {
			fmt.Println(`error scan`)
			//log.Fatal(err)
		}
	}
	fmt.Println("attractions", attractions)
	var count int
	qry, err = db.Db.Query(`SELECT COUNT(*) FROM Attractions WHERE Attraction = $`, attraction)
	if err != nil {
		fmt.Println(`error query`)
	}
	for qry.Next() {
		if err := qry.Scan(&count); err != nil {
			fmt.Println(`error scan`)
			//log.Fatal(err)
		}
	}

	return rating
}

func GetAttractionsByTraveler(traveler Traveler) ([]Attraction, error) {
	Attractions := make([]Attraction, 0)
	qry, err := db.Db.Query(`select * from Attraction where traveler = $`, traveler.Name)
	if err != nil {
		return nil, err
	}
	var attractions Attraction
	for qry.Next() {
		err := qry.Scan(&attractions.Title, &attractions.RangeFromCenter, &attractions.Rating)
		if err != nil {
			return nil, err
		}
		Attractions = append(Attractions, attractions)
	}
	return Attractions, nil
}

func GetAllAttractions() {
	//`select ID, COALESCE(ResponsibleID,0), COALESCE(Title, ''),
	//       COALESCE(Name,''), COALESCE(Phone,''), COALESCE(DateCreate,''),
	//       COALESCE(SourceId,''), COALESCE(SourceDescription,''), COALESCE(AssignedByLead,''), COALESCE(Email,''),
	//       COALESCE(FormName,''
}*/
