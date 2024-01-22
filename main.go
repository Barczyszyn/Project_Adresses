package main

import (
	"adresses/app"
	"log"
	"os"
)

//Search for IP example: ./adresses.exe ip --host amazon.com.br
//Search for NS example: ./adresses.exe ns --host amazon.com.br

func main() {
	application := app.Generate()
	if erro := application.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
