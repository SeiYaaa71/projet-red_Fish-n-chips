package main

import "fmt"

func equipItem(c *Character, item string) {
	switch item {
	case "Casque du Chasseur":
		if c.Equipement.Tete != "" {
			c.Inventaire = append(c.Inventaire, c.Equipement.Tete)
			fmt.Println(Yellow + "↺ " + c.Equipement.Tete + " a été replacé dans l’inventaire." + Reset)
		}
		c.Equipement.Tete = item
		fmt.Println(Green + "✅ Vous avez équipé : " + item + Reset)

	case "Armure du Chasseur":
		if c.Equipement.Torse != "" {
			c.Inventaire = append(c.Inventaire, c.Equipement.Torse)
			fmt.Println(Yellow + "↺ " + c.Equipement.Torse + " a été replacé dans l’inventaire." + Reset)
		}
		c.Equipement.Torse = item
		fmt.Println(Green + "✅ Vous avez équipé : " + item + Reset)

	case "Bottes du Chasseur":
		if c.Equipement.Pieds != "" {
			c.Inventaire = append(c.Inventaire, c.Equipement.Pieds)
			fmt.Println(Yellow + "↺ " + c.Equipement.Pieds + " a été replacé dans l’inventaire." + Reset)
		}
		c.Equipement.Pieds = item
		fmt.Println(Green + "✅ Vous avez équipé : " + item + Reset)

	// Directement équiper l’épée
	if c.Equipement.Arme != "" {
	c.Inventaire = append(c.Inventaire, c.Equipement.Arme)
	fmt.Println(Yellow + "↺ " + c.Equipement.Arme + " a été replacé dans l’inventaire." + Reset)
	}
	c.Equipement.Arme = "Épée du Forgeron (+10 dmg)"
	c.Attaque += 10
	fmt.Println(Green + "✅ Vous avez forgé et équipé : Épée du Forgeron (+10 dmg)" + Reset)
	}

	// Retirer l’objet de l’inventaire
	for i, invItem := range c.Inventaire {
		if invItem == item {
			c.Inventaire = append(c.Inventaire[:i], c.Inventaire[i+1:]...)
			break
		}
	}

	// Recalcul des PV max
	recalcPVMax(c)
}

// --------------------
// Recalcule PVMax selon équipement
// --------------------
func recalcPVMax(c *Character) {
	basePV := 100 // tu peux adapter selon la classe
	bonus := 0

	if c.Equipement.Tete == "Casque du Chasseur" {
		bonus += 10
	}
	if c.Equipement.Torse == "Armure du Chasseur" {
		bonus += 25
	}
	if c.Equipement.Pieds == "Bottes du Chasseur" {
		bonus += 15
	}

	c.PVMax = basePV + bonus

	if c.PVActuels > c.PVMax {
		c.PVActuels = c.PVMax
	}
}
