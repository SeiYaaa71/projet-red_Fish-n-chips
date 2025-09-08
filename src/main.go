package main

import (
	"fmt"
)

func main() {
	inv := []Inventory{
		{name: "Potion", quantity: 3},
	}
	c1 := initCharacter("Fish", "Elfe", 1, 100, 40, inv)
	fmt.Println(c1)
}
