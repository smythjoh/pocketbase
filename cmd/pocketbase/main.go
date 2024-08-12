package main

import (
	"log"

	"github.com/pocketbase/pocketbase"

	_ "github.com/abhi2303237/pocketbase/migrations"
)

// func main() {
// 	app := pocketbase.New()

// 	if err := app.Start(); err != nil {
// 		panic(err)
// 	}
// }

func main() {
	app := pocketbase.New()

	if err := app.Start(); err != nil {
		log.Print("error")
		log.Fatal(err)
	}
}
