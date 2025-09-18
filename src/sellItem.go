package main

import "fmt"

// Fonction pour vendre un objet
func sellItem(c *Character) {
	if len(c.Inventaire) == 0 {
		fmt.Println(Red + "❌ Votre inventaire est vide, rien à vendre." + Reset)
		waitForEnter()
		return
	}

	fmt.Println(Bold + Blue + "\n=== VENTE D’OBJETS ===" + Reset)
	fmt.Println("Sélectionnez l’objet à vendre :")
	for i, item := range c.Inventaire {
		fmt.Printf(Green+"%d."+Reset+" %s\n", i+1, item)
	}

	fmt.Print(Yellow + "\nVotre choix : " + Reset)
	var choix int
	fmt.Scanln(&choix)

	if choix < 1 || choix > len(c.Inventaire) {
		fmt.Println(Red + "❌ Choix invalide." + Reset)
		waitForEnter()
		return
	}

	item := c.Inventaire[choix-1]

	// 💰 Détermine la valeur de vente (ex: moitié du prix d’achat)
	prixVente := getItemSellPrice(item)
	if prixVente == 0 {
		fmt.Println(Red + "❌ Cet objet ne peut pas être vendu." + Reset)
		waitForEnter()
		return
	}

	// Retire l’objet de l’inventaire
	c.Inventaire = append(c.Inventaire[:choix-1], c.Inventaire[choix:]...)

	// Ajoute l’or
	c.Gold += prixVente

	fmt.Printf(Green+"✅ Vous avez vendu %s pour %d or.\n"+Reset, item, prixVente)
	waitForEnter()
}

// Détermine le prix de vente en fonction de l’objet
func getItemSellPrice(item string) int {
	switch item {
	case "Potion de vie":
		return 2
	case "Potion de poison":
		return 3
	case "Fourrure de Loup":
		return 2
	case "Peau de Goblin":
		return 4
	case "Cuir de Sanglier":
		return 2
	case "Plume de Corbeau":
		return 1
	case "Fer":
		return 3
	case "Bois":
		return 1
	case "Épée du Forgeron (+10 dmg)":
		return 10
	default:
		return 0 // invendable
	}
}
