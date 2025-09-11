package main

func main() {
	c1 := initCharacter(
		"TonNom",
		"Elfe",
		1,
		100,
		40,
		[]string{"Potion"},
	)
	menu(&c1)
	isDead(&c1)
	displayInfo(c1)
}
