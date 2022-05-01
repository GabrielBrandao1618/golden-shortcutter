package database

import (
	"github.com/GabrielBrandao1618/golden-shortcutter/model/shortedLink"
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

type linkDbModel struct {
	gorm.Model
	Ref  string
	Name string
}

func Migrate() {
	db := getDatabase()

	db.AutoMigrate(&linkDbModel{})
}

func GetUrlByCustomName(name string) string {
	db := getDatabase()

	var result linkDbModel
	db.First(&result, "name = ?", name)
	return result.Ref
}

func checkIfLinkAlreadyExists(ref string) (exists bool, hashCode string) {
	db := getDatabase()
	var result linkDbModel

	db.First(&result, "ref = ?", ref)

	return result.Name != "", result.Name
}

func CreateUrl(ref string, name string) shortedLink.ShortedLink {
	db := getDatabase()
	var linkToInsert shortedLink.ShortedLink
	linkToInsert.Ref = ref

	linkAlreadyExists, existingCustomName := checkIfLinkAlreadyExists(ref)

	if !linkAlreadyExists {
		linkToInsert = shortedLink.ShortedLink{Ref: ref, Name: name}
		db.Create(&linkDbModel{Ref: linkToInsert.Ref, Name: linkToInsert.Name})
	} else {
		linkToInsert.Name = existingCustomName
	}
	return linkToInsert

}
