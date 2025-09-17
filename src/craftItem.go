package main

import "fmt"

var recipe = map[string]map[string]int{
	"Casque du Chasseur": {"Plume de Corbeau": 1, "Cuir de Sanglier": 1},
	"Armure du Chasseur": {"Fourrure de Loup": 2, "Peau de Goblin": 1},
	"Bottes du Chasseur":  {"Fourrure de Loup": 1, "Cuir de Sanglier": 1},
}

func craftItem(c *Character, item string, price int) {
	// Vérification de l'or
	if c.Gold < price {
		fmt.Printf(Red+"❌ Pas assez d'or pour %s. (%d or requis, vous avez %d)\n"+Reset, item, price, c.Gold)
		waitForEnter()
		return
	}

	// Vérification de l'inventaire (au cas où, mais on n’ajoute plus l’objet dedans)
	if len(c.Inventaire) >= c.InventoryMax {
		fmt.Println(Red + "❌ Inventaire plein, impossible de fabriquer cet objet." + Reset)
		waitForEnter()
		return
	}

	// Vérification des matériaux nécessaires
	materials, ok := recipe[item]
	if ok {
		for mat, qty := range materials {
			if countItem(c, mat) < qty {
				fmt.Printf(Red+"❌ Il vous manque %d %s pour fabriquer %s.\n"+Reset, qty, mat, item)
				waitForEnter()
				return
			}
		}

		// Retire les matériaux de l'inventaire
		for mat, qty := range materials {
			for i := 0; i < qty; i++ {
				removeInventory(c, mat)
			}
		}
	}

	// Déduction de l'or
	c.Gold -= price

	// ⚡ Au lieu d’ajouter dans l’inventaire → équipe directement
	equipItem(c, item)

	fmt.Printf(Green+"✅ Vous avez fabriqué et équipé %s pour %d or.\n"+Reset, item, price)
	waitForEnter()
}
