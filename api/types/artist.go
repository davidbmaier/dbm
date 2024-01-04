package types

type Artist struct {
	ID        uint   `json:"id"`
	Name      string `gorm:"uniqueIndex" json:"name"`
	BirthYear *int   `json:"birthYear"`
	DeathYear *int   `json:"deathYear"`
}

type ArtistWithNumberOfWorks struct {
	Artist
	NumberOfWorks int `json:"numberOfWorks"`
}

type ArtistsResponse struct {
	Artists []ArtistWithNumberOfWorks `json:"artists"`
	Total   int64                     `json:"total"`
}
