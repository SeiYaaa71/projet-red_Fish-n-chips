package main

func main() {
	c1 := initCharacter(
		"Marcel",
		"Elfe",
		1,
		100,
		40,
		[]string{"Potion", "Potion", "Potion"},
	)
	menu(&c1)
}
