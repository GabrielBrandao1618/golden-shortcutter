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
	Ref      string
	HashCode string
}

func Migrate() {
	db := getDatabase()

	db.AutoMigrate(&urlDatabaseModel{})
}

func GetUrlByHashCode(hashCode string) string {
	db := getDatabase()

	var result urlDatabaseModel
	db.First(&result, "hash_code = ?", hashCode)
	return result.Ref
}

func CreateUrl(link shortedLink.ShortedLink) {
	db := getDatabase()

	db.Create(&urlDatabaseModel{Ref: link.Ref, HashCode: link.HashCode})
}
