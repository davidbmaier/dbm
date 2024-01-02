package database

import (
	"github.com/davidbmaier/dbm/types"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var env map[string]string

func InitDatabaseConnection(envMap map[string]string) *gorm.DB {
	env = envMap

	// only create a new connection if there is no existing one yet
	if db == nil {
		var err error
		db, err = gorm.Open(postgres.Open(env["DATABASE_STRING"]), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db.AutoMigrate(&types.Artist{})
		db.AutoMigrate(&types.Work{})
	}

	log.Info("Database connection established")
	return db
}

func RetrieveArtists(search string, page int) types.ArtistsResponse {
	pageSize := 20
	offset := (page - 1) * pageSize

	var artists []types.Artist
	var totalArtistsCount int64
	db.Order("ID").Where("name ILIKE ?", "%"+search+"%").Limit(pageSize).Offset(offset).Find(&artists)
	db.Model(&types.Artist{}).Where("name ILIKE ?", "%"+search+"%").Count(&totalArtistsCount)

	return types.ArtistsResponse{
		Artists: []types.Artist(artists),
		Total:   totalArtistsCount,
	}
}

func RetrieveArtist(artistID int) types.Artist {
	var artist types.Artist
	db.First(&artist, artistID)

	return artist
}

func RetrieveWorks(search string, artistID int, page int) types.WorksResponse {
	pageSize := 20
	offset := (page - 1) * pageSize

	var works []types.Work
	var totalWorksCount int64

	// get all works if artistID is 0 (unset)
	if artistID != 0 {
		db.Preload("Artist").Order("ID").Where("artist_id = ?", artistID).Where("name ILIKE ?", "%"+search+"%").Limit(pageSize).Offset(offset).Find(&works)
		db.Model(&types.Work{}).Preload("Artist").Where("artist_id = ?", artistID).Where("name ILIKE ?", "%"+search+"%").Count(&totalWorksCount)
	} else {
		db.Preload("Artist").Order("ID").Where("name ILIKE ?", "%"+search+"%").Limit(pageSize).Offset(offset).Find(&works)
		db.Model(&types.Work{}).Preload("Artist").Where("name ILIKE ?", "%"+search+"%").Count(&totalWorksCount)
	}

	return types.WorksResponse{
		Works: []types.Work(works),
		Total: totalWorksCount,
	}
}

func RetrieveWork(workID int) types.Work {
	var work types.Work
	db.Preload("Artist").First(&work, workID)

	return work
}
