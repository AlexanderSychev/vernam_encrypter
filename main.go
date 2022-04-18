package main

import (
	"log"
	"os"
	"vernam_encrypter/cmd"
)

func main() {
	app := cmd.NewApp()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
