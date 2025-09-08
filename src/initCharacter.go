package main

type Inventory struct {
    name     string
    quantity int
}

type Character struct {
    name      string
    class     string
    level     int
    max_hp    int
    hp        int
    inventory []Inventory
}

func initCharacter(name, class string, level, max_hp, hp int, inventory []Inventory) Character {
    return Character{
        name:      name,
        class:     class,
        level:     level,
        max_hp:    max_hp,
        hp:        hp,
        inventory: inventory,
    }
}
