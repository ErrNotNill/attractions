package repo

type Sight struct {
	SightID  int    `json:"sight_id,omitempty"`
	Title    string `json:"title,omitempty"`
	Distance string `json:"distance,omitempty"`
}
type City struct {
	CityID int    `json:"cityID"`
	Name   string `json:"title,omitempty"`
}
type Traveler struct {
	TravelerID int    `json:"traveler_id,omitempty"`
	Name       string `json:"name,omitempty"`
}
type Visit struct {
	VisitID    int      `json:"visit_id,omitempty"`
	SightID    Sight    `json:"sight_id"`
	TravelerID Traveler `json:"traveler_id"`
	VisitDate  Date     `json:"visit_date"`
	Rating     int      `json:"rating,omitempty"`
}
type Rating struct {
	RatingID    int      `json:"rating_id,omitempty"`
	SightID     Sight    `json:"sight_id"`
	TravelerID  Traveler `json:"traveler_id"`
	RatingValue int      `json:"rating_value,omitempty"`
}
