package main

import "fmt"

type Inventory struct {
	name string
	quantity int
}

type Character struct {
	name string
	class string
	level int
	max_hp int
	hp int
	inventory []inventory
}
