package repo

import (
	"attractions/db"
	"fmt"
	"time"
)

func SetRatingForAttraction(attraction string) {
	//`insert into Attractions value (rating) where Attraction name = $`,attraction
}

func (c City) AddCity(title string) {
	city := &City{}
	city.Title = title
	result, err := db.Db.Exec(`insert into City (
      Title) values (?)`,
		city.Title)
	if err != nil {
		fmt.Println("cant insert data to dbase")
		panic(err)
	}
	fmt.Println("rows inserted")
	fmt.Println(result.RowsAffected())
}

func (t Traveler) AddTraveler(name string, date Date) {
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
}

func (a Attraction) AddAttraction() {
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
}

type Date struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
	Second int
}
