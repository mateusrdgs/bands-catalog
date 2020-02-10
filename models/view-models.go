package models

// Band struct work as an interface to bands
type Band struct {
	Name             string `json:"name" form:"name"`
	YearOfFoundation int    `json:"year_of_foundation" form:"year_of_foundation"`
	Biography        string `json:"biography" form:"biography"`
	Country          string `json:"country" form:"country"`
	Genre            string `json:"genre" form:"genre"`
}

// Artist struct work as an interface to artists
type Artist struct {
	Name       string `json:"name" form:"name"`
	Age        int    `json:"age" form:"age"`
	Country    string `json:"country" form:"country"`
	Genre      string `json:"genre" form:"genre"`
	Biography  string `json:"biography" form:"biography"`
	Instrument string `json:"instrument" form:"instrument"`
}

// Album struct work as an interface to albums
type Album struct {
	Name        string `json:"name" form:"name"`
	Type        string `json:"type" form:"type"`
	ReleaseDate string `json:"release_date" form:"release_date"`
	Label       string `json:"label" form:"label"`
}

// Song struct work as an interface to songs
type Song struct {
	Name     string `json:"name" form:"name"`
	Number   int    `json:"number" form:"number"`
	Duration int    `json:"duration" form:"duration"`
	Lyrics   string `json:"lyrics" form:"lyrics"`
}
