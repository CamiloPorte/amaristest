package main

import (
	"fmt"
)

func main() {
	var u, v string
	fmt.Scanln(&u)
	fmt.Scanln(&v)
	if StringsFriends(u, v) {
		fmt.Println("cadenas amigas")
	} else {
		fmt.Println("cadenas no amigas")
	}

}

func StringsFriends(u string, v string) bool {
	if u == v {
		return true
	}
	if len(u) != len(v) {
		return false
	}
	for i := 0; i < len(u); i++ {
		aux := v[i:] + v[:i]
		if u == aux {
			return true
		}
	}
	return false
}
