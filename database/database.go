package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/GabrielBrandao13/golden-shortcutter/model/shortedLink"
)

func getDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})

	if err != nil {
		panic("Erro ao conectar-se ao banco de dados")
	}
	return db
}

type urlDatabaseModel struct {
	gorm.Model
	ref      string
	hashCode string
}

func Migrate() {
	db := getDatabase()

	db.AutoMigrate(&urlDatabaseModel{})
}

func GetUrlByHashCode(hashCode string) string {
	db := getDatabase()

	var result urlDatabaseModel
	db.First(&result, "hashCode = ?", hashCode)
	return result.ref
}

func CreateUrl(link shortedLink.ShortedLink) {
	db := getDatabase()

	db.Create(&urlDatabaseModel{ref: link.Ref, hashCode: link.HashCode})
}
