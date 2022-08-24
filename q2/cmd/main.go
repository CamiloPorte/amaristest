package main

import (
	"fmt"
	"log"
	"q2/client"
)

func main() {
	logger := log.Default()
	PA := client.NewPokeApi("https://pokeapi.co/api/v2/pokemon-form/", *logger)
	answ, err := PA.Get(1)
	if err != nil {
		logger.Println("error in the call:", err)
	}
	fmt.Println("the name is: ", answ)
}
