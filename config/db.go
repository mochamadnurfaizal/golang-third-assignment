package config

import (
	"fmt"
	"golang-third-assignment/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbGorm *gorm.DB
var err error

const host = "127.0.0.1"
const port = 5432
const user = "postgres"
const password = "postgres"
const dbname = "third-assignment"

func ConnectGorm() {

	psqlInfo := fmt.Sprintf(`
	host=%s
	port=%d
	user=%s`+`
	password=%s
	dbname=%s
	sslmode=disable`, host, port, user, password, dbname)

	DbGorm, err = gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DbGorm.AutoMigrate(models.Environtment{})
	fmt.Println("Sukses Konek DB")

}

func GetDB() *gorm.DB {
	return DbGorm
}
