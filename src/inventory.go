package main

import "fmt"

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

// Ajouter un item dans l'inventaire si possible
func addToInventory(c *Character, item string) bool {
	if len(c.Inventaire) >= c.InventoryMax {
		fmt.Println(Red + "❌ Inventaire plein, impossible d’ajouter : " + item + Reset)
		return false
	}
	c.Inventaire = append(c.Inventaire, item)
	fmt.Println(Green + "✅ " + item + " ajouté à votre inventaire." + Reset)
	return true
}

// Supprimer un item de l'inventaire
func removeInventory(c *Character, item string) bool {
	for i, v := range c.Inventaire {
		if v == item {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			return true
		}
	}
	return false
}

func hasItem(c *Character, item string, count int) bool {
	quantité := 0
	for _, invItem := range c.Inventaire {
		if invItem == item {
			quantité++
			if quantité >= count {
				return true
			}
		}
	}
	return false
}

// Retire `count` exemplaires d’un item de l’inventaire
func removeItem(c *Character, item string, count int) {
	nouveauInv := []string{}
	àRetirer := count

	for _, invItem := range c.Inventaire {
		if invItem == item && àRetirer > 0 {
			àRetirer--
			continue // on "consomme" l’item, donc on ne le garde pas
		}
		nouveauInv = append(nouveauInv, invItem)
	}

	c.Inventaire = nouveauInv
}
