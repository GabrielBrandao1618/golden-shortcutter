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
	shortedLink.ShortedLink
	Visits uint64
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

func checkIfLinkAlreadyExists(ref string) string {
	db := getDatabase()
	var result linkDbModel

	db.First(&result, "ref = ?", ref)

	return result.Name
}
func checkIfCustomNameIsAlreadyInUse(customName string) bool {
	db := getDatabase()
	var result linkDbModel

	db.First(&result, "name = ?", customName)

	return result.Ref != ""
}

type createUrlResult struct {
	Sucess      bool `json:"sucess"`
	Msg         string `json:"msg"`
	ShortedLink shortedLink.ShortedLink `json:"shortedLink"`
}

func CreateUrl(ref string, name string) createUrlResult {
	db := getDatabase()
	var result createUrlResult
	result.ShortedLink.Ref = ref

	if checkIfCustomNameIsAlreadyInUse(name) {
		result.Sucess = false
		result.Msg = "Custom name already in use"
		return result
	}

	existingCustomName := checkIfLinkAlreadyExists(ref)

	if existingCustomName == "" {
		result.ShortedLink = shortedLink.ShortedLink{Ref: ref, Name: name}
		db.Create(&linkDbModel{ShortedLink: shortedLink.ShortedLink{Ref: result.ShortedLink.Ref, Name: result.ShortedLink.Name}})
		result.Msg = "Link created successfully!"
		result.Sucess = true
		return result
	}

	result.ShortedLink.Name = existingCustomName
	result.Msg = "Link already created!"
	result.Sucess = true
	return result

}

func IncrementVisitsCount(customName string){
	var link linkDbModel
	db := getDatabase()

	db.First(&link, "name = ?", customName)
	link.Visits = link.Visits+1
	db.Save(&link)
}
