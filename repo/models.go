package repo

type Data struct {
	Attraction Attraction
	City       City
	Traveler   Traveler
}

type Attraction struct {
	ID              int      `json:"attraction_id,omitempty"`
	Traveler        Traveler `json:"traveler,omitempty"`
	City            City     `json:"city,omitempty"`
	Title           string   `json:"title,omitempty"`
	RangeFromCenter string   `json:"range_from_center,omitempty"`
	Description     string   `json:"description,omitempty"`
	Rating          string   `json:"rating,omitempty"`
}

type City struct {
	ID     int    `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Rating int    `json:"rating,omitempty"`
}

type Traveler struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	TimeStarted string `json:"timeStarted,omitempty"`
	TimeEnded   string `json:"timeEnded,omitempty"`
}

type TravelerProfile struct {
	Attraction  []*Attraction `json:"attraction,omitempty"`
	TimeStarted string        `json:"time_started,omitempty"`
	TimeEnded   string        `json:"time_end,omitempty"`
}
