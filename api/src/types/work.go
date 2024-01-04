package types

type Work struct {
	ID           uint    `json:"id"`
	WorkID       string  `gorm:"uniqueIndex" json:"workID"`
	ArtistID     uint    `json:"artistID"`
	Artist       Artist  `json:"artist"`
	Name         *string `json:"name"`
	CreationInfo *string `json:"creationInfo"`
	Material     *string `json:"material"`
	Size         *string `json:"size"`
	Owner        *string `json:"owner"`
	Source       *string `json:"source"`
}

type WorksResponse struct {
	Works []Work `json:"works"`
	Total int64  `json:"total"`
}
