package types

type User struct {
	ID             uint
	Username       string `gorm:"uniqueIndex"`
	HashedPassword string
	Admin          bool
}
