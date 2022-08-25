package main

import (
	"fmt"
	"q2/client"
)

func main() {
	PA := client.NewPokeApi("https://pokeapi.co/api/v2/pokemon-form/")
	var option int
	fmt.Scanln(&option)
	answ, err := PA.Get(option)
	if err != nil {
		fmt.Println("error in the call:", err)
	}
	fmt.Println("the name is: ", answ)
}
