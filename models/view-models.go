package models

// Band struct work as an interface to bands
type Band struct {
	UUID             string `json:"uuid"`
	Name             string `json:"name"`
	YearOfFoundation int    `json:"year_of_foundation"`
	Biography        string `json:"biography"`
	Country          string `json:"country"`
	Genre            string `json:"genre"`
}

// Artist struct work as an interface to artists
type Artist struct {
	UUID       string `json:"uuid"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Country    string `json:"country"`
	Genre      string `json:"genre"`
	Biography  string `json:"biography"`
	Instrument string `json:"instrument"`
}

// Album struct work as an interface to albums
type Album struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	ReleaseDate string `json:"release_date"`
	Label       string `json:"label"`
	BandUUID    string `json:"band_uuid"`
}

// Song struct work as an interface to songs
type Song struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Number   int    `json:"number"`
	Duration int    `json:"duration"`
	Lyrics   string `json:"lyrics"`
}
