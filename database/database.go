package database

import (
	"github.com/GabrielBrandao13/golden-shortcutter/model/shortedLink"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

func checkIfUrlAlreadyExists(ref string) (exists bool, hashCode string) {
	db := getDatabase()
	var result urlDatabaseModel

	db.First(&result, "ref = ?", ref)
	if result.HashCode != "" {
		return true, result.HashCode
	}
	return result.HashCode != "", result.HashCode
}

func CreateUrl(ref string) shortedLink.ShortedLink {
	db := getDatabase()
	var linkToInsert shortedLink.ShortedLink
	linkToInsert.Ref = ref

	urlAlreadyExists, existingHashCode := checkIfUrlAlreadyExists(ref)

	if !urlAlreadyExists {
		linkToInsert = shortedLink.New(ref)
		db.Create(&urlDatabaseModel{Ref: linkToInsert.Ref, HashCode: linkToInsert.HashCode})
	} else {
		linkToInsert.HashCode = existingHashCode
	}
	return linkToInsert

}
