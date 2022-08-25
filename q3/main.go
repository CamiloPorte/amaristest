package main

import (
	"fmt"
)

func main() {
	var u, v string
	fmt.Scanln(&u)
	fmt.Scanln(&v)
	StringsFriends(u, v)

}

func StringsFriends(u string, v string) {
	if u == v {
		fmt.Println("cadenas amigas")
		return
	}
	for i := 0; i < len(u); i++ {
		aux := v[i:] + v[:i]
		if u == aux {
			fmt.Println("cadenas amigas")
			return
		}
	}
	fmt.Println("cadenas no amigas")
}
