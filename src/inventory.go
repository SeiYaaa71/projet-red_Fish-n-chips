package main

// Retire le premier exemplaire d'un item
func removeInventory(c *Character, item string) {
	for i, v := range c.Inventaire {
		if v == item {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			return
		}
	}
}

// Compte combien de fois un item est présent
func countItem(c *Character, item string) int {
	count := 0
	for _, v := range c.Inventaire {
		if v == item {
			count++
		}
	}
	return count
}
