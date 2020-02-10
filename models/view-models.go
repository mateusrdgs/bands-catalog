package models

// Band struct serves as interface to bands
type Band struct {
	Name             string `json:"band_name" form:"band_name"`
	YearOfFoundation int    `json:"year_of_foundation" form:"year_of_foundation"`
	Biography        string `json:"biography" form:"biography"`
	Country          string `json:"country" form:"country"`
	Genre            string `json:"genre" form:"genre"`
}
