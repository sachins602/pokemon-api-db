package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mtslzr/pokeapi-go"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	// Load environment variables from file.
	if err := godotenv.Load(); err != nil {
		log.Fatalf("failed to load environment variables: %v", err)
	}

	// Connect to PlanetScale database using DSN environment variable.
	db, err := gorm.Open(mysql.Open(os.Getenv("DSN")), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		log.Fatalf("failed to connect to PlanetScale: %v", err)
	}

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
		time.Sleep(time.Millisecond * 100)
	}

}

type PokemonStruct struct {
	Id        int    `gorm:"primary_key column:id"`
	Name      string `gorm:"column:name"`
	Attack    int    `gorm:"column:attack"`
	SpriteUrl string `gorm:"column:spriteUrl"`
}
