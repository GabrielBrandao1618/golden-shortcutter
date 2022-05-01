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
func checkIfCustomNameIsAlreadyInUse(customName string) bool {
	db := getDatabase()
	var result linkDbModel

	db.First(&result, "name = ?", customName)

	return result.Ref != ""
}

type createUrlResult struct {
	sucess      bool
	msg         string
	shortedLink shortedLink.ShortedLink
}

func CreateUrl(ref string, name string) createUrlResult {
	db := getDatabase()
	var result createUrlResult
	result.shortedLink.Ref = ref

	if checkIfCustomNameIsAlreadyInUse(name) {
		result.sucess = false
		result.msg = "Custom name already in use"
		return result
	}

	linkAlreadyExists, existingCustomName := checkIfLinkAlreadyExists(ref)

	if !linkAlreadyExists {
		result.shortedLink = shortedLink.ShortedLink{Ref: ref, Name: name}
		db.Create(&linkDbModel{Ref: result.shortedLink.Ref, Name: result.shortedLink.Name})
		result.msg = "Link created successfully!"
		result.sucess = true
		return result
	}

	result.shortedLink.Name = existingCustomName
	return result

}
