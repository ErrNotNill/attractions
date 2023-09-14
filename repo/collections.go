package repo

import (
	"attractions/db"
	"fmt"
)

func SetRatingForAttraction(attraction string) {
	//`insert into Attractions value (rating) where Attraction name = $`,attraction
}

func AddRating(rating int) {
	//SightID, TravelerID, RatingValue
	sights := GetSight()
	var sightID int
	for _, v := range sights {
		sightID = v.SightID
	}
	var travelerID int
	travelers := GetTraveler()
	for _, v := range travelers {
		travelerID = v.TravelerID
	}
	result, err := db.Db.Exec(`insert into Rating (SightID, TravelerID, RatingValue
      ) values (?,?,?)`,
		&sightID, &travelerID, &rating)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}

func AddCity() {
	city := &City{}
	//TravelerID,SightID,CityID
	result, err := db.Db.Exec(`insert into City (Name
     ) values (?)`,
		&city.Name)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}

func AddSightInCity() {
	sight := &Sight{}
	cities := GetCity()
	for _, city := range cities {
		fmt.Println(city.CityID)
	}
	result, err := db.Db.Exec(`insert into Sight (CityID
      ) values (?)`,
		sight.SightID)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}

func AddSightByTraveler(travelerId int) {
	TravelerID := GetTraveler()
	var travelers []int
	for _, v := range TravelerID {
		travelers = append(travelers, v.TravelerID)
	}
	fmt.Println(travelers)

	cityID := GetCity()
	fmt.Println("cityID in AddSightByTraveler: ", cityID)
	//`insert into Sight Title, Distance, CityID
	/*sight := &Sight{}
		city.Title = title
		result, err := db.Db.Exec(`insert into City (
	      Title) values (?)`,
			city.Title)
		if err != nil {
			fmt.Println("cant insert data to dbase")
			panic(err)
		}
		fmt.Println("rows inserted")
		fmt.Println(result.RowsAffected())*/
}

/*func AddRating(rating int) {
	if rating < 0 || rating > 5 {
		fmt.Println("Rating cannot be negative or greater than 5")
		return
	}
	SightID := GetSight
	TravelerID := GetTraveler

	result, err := db.Db.Exec(`insert into Rating (
      SightID, TravelerID, RatingValue) values (?,?,?)`,
		traveler.Name, traveler.TimeStarted, rating)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}

}*/

/*func (t Traveler) AddTraveler(name string, date Date) {
	traveler := &Traveler{}
	timeNow := time.Now()
	timeStart := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		timeNow.Year(), timeNow.Month(), timeNow.Day(),
		timeNow.Hour(), timeNow.Minute(), timeNow.Second())
	timeEnd := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		date.Year, date.Month, date.Day,
		date.Hour, date.Minute, date.Second)
	if timeEnd < timeStart {
		fmt.Println("timeEnd", timeEnd)
		fmt.Println("timeEnd < timeStart")
		return
	}
	traveler.TimeStarted = timeStart
	traveler.TimeEnded = timeEnd
	traveler.Name = name

	result, err := db.Db.Exec(`insert into Traveler (
      Name,TimeStarted,TimeEnded) values (?,?,?)`,
		traveler.Name, traveler.TimeStarted, traveler.TimeEnded)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}*/

/*func (a Attraction) AddAttraction() {
	result, err := db.Db.Exec(`insert into Attraction (
                        traveler_id,city_id,Title,RangeFromCenter,Rating) values (1,?,?,?,?)`,
		a.City.ID, a.Title, a.RangeFromCenter, a.Rating)
	fmt.Println("traveler.ID, city.ID", a.Traveler.ID, a.City.ID)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}*/

type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}
