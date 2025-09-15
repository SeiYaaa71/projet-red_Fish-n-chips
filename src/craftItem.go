package main

import "fmt"

var recipe = map[string]map[string]int{
	"Chapeau de l’aventurier": {"Plume de Corbeau": 1, "Cuir de Sanglier": 1},
	"Tunique de l’aventurier": {"Fourrure de Loup": 2, "Peau de Troll": 1},
	"Bottes de l’aventurier":  {"Fourrure de Loup": 1, "Cuir de Sanglier": 1},
}

func craftItem(c *Character, item string, price int) {
	// Vérification de l'or
	if c.Gold < price {
		fmt.Printf(Red+"❌ Pas assez d'or pour %s. (%d or requis, vous avez %d)\n"+Reset, item, price, c.Gold)
		waitForEnter()
		return
	}

	// Vérification de l'inventaire disponible
	if len(c.Inventaire) >= c.InventoryMax {
		fmt.Println(Red + "❌ Inventaire plein, impossible de fabriquer cet objet." + Reset)
		waitForEnter()
		return
	}

	// Vérification des matériaux
	materials, ok := recipe[item]
	if ok {
		for mat, qty := range materials {
			if countItem(c, mat) < qty {
				fmt.Printf(Red+"❌ Il vous manque %d %s pour fabriquer %s.\n"+Reset, qty, mat, item)
				waitForEnter()
				return
			}
		}

		// Suppression des matériaux
		for mat, qty := range materials {
			for i := 0; i < qty; i++ {
				removeInventory(c, mat)
			}
		}
	}

	// Déduction de l'or et ajout de l'objet
	c.Gold -= price
	c.Inventaire = append(c.Inventaire, item)
	fmt.Printf(Green+"✅ Vous avez fabriqué %s pour %d or.\n"+Reset, item, price)
	waitForEnter()
}
