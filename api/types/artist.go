package types

type Artist struct {
	ID        uint   `json:"id"`
	Name      string `gorm:"uniqueIndex" json:"name"`
	BirthYear *int   `json:"birthYear"`
	DeathYear *int   `json:"deathYear"`
}

type ArtistsResponse struct {
	Artists []Artist `json:"artists"`
	Total   int64    `json:"total"`
}
