package main

import (
	"fmt"
	"log"
	"os"

	//"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/mtslzr/pokeapi-go"
)

var db *gorm.DB

func main() {

	// Load environment variables from file.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	db, err = gorm.Open("mysql", DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database", Dbdriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", Dbdriver)
	}

	// Connect to PlanetScale database using DSN environment variable.
	// db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
	// 	DisableForeignKeyConstraintWhenMigrating: true,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to connect to PlanetScale: %v", err)
	// }

	for index := 1; index < 500; index++ {
		i := fmt.Sprintf("%d", index)
		pokemonData, err := pokeapi.Pokemon(i)
		if err != nil {
			fmt.Println(err)
		}
		id := index
		name := pokemonData.Name
		attack := pokemonData.Stats[4].BaseStat
		spriteUrl := fmt.Sprintf("https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/other/home/%d.png", index)

		if err != nil {
			fmt.Println(err)
		}

		pokemon := &PokemonStruct{
			Id:        id,
			Name:      name,
			Attack:    attack,
			SpriteUrl: spriteUrl,
		}

		db.Table("Pokemon").Create(&pokemon)
		fmt.Println("Pokemon added to database")

		fmt.Println(name, attack, spriteUrl)
		//time.Sleep(time.Millisecond * 100)
	}

}

type PokemonStruct struct {
	Id        int    `gorm:"primary_key column:id"`
	Name      string `gorm:"column:name"`
	Attack    int    `gorm:"column:attack"`
	SpriteUrl string `gorm:"column:spriteUrl"`
}
