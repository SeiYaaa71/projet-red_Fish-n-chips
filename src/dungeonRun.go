package main

import (
	"fmt"
	"math/rand"
)

// Lancement du Donjon
func startDungeonRun(c *Character) {
	clearScreen()
	fmt.Println(Bold + Purple + "\n🏰 Vous entrez dans le Donjon !" + Reset)
	fmt.Printf("Vous êtes à l’étage %d.\n\n", currentFloor)
	waitForEnter()

	// Tant que le joueur est en vie, il peut continuer
	for !c.isDead() {
		clearScreen()
		fmt.Printf(Bold+Blue+"=== ÉTAGE %d ===\n"+Reset, currentFloor)

		// Récupérer les multiplicateurs de l’étage
		floorData := getFloorData(currentFloor)

		// Générer un monstre adapté à cet étage
		monster := initMonsterForFloor(floorData)

		// Combat contre le monstre
		trainingFightWithFloor(c, monster, currentFloor)

		// Vérifie si le joueur a abandonné
		if c.AbandonCombat {
			fmt.Println(Red + "\n🏳 Vous avez quitté le Donjon." + Reset)
			c.AbandonCombat = false // reset pour la prochaine fois
			break
		}

		// Vérifie si le joueur est mort
		if c.isDead() {
			fmt.Println(Red + "\n💀 Vous êtes mort dans le Donjon !" + Reset)
			break
		}

		// 🎁 Bonus coffre tous les 3 étages
		if currentFloor%3 == 0 {
			coffreBonus(c)
		}

		// Demander au joueur s’il veut continuer
		var choix string
		fmt.Print(Yellow + "\nVoulez-vous continuer au prochain étage ? (o/n) : " + Reset)
		fmt.Scanln(&choix)

		if choix == "o" || choix == "O" {
			currentFloor++ // On monte d’un étage
			fmt.Printf(Green+"\n⬆ Vous progressez vers l’étage %d !"+Reset+"\n", currentFloor)
			waitForEnter()
			continue
		} else {
			fmt.Println(Green + "\n🏳 Vous quittez le Donjon avec vos gains !" + Reset)
			break
		}
	}
}

// 🎁 Système de coffre bonus (tous les 3 étages)
func coffreBonus(c *Character) {
	clearScreen()
	fmt.Println(Yellow + Bold + "\n🎁 Vous trouvez un coffre mystérieux !" + Reset)

	rewardType := rand.Intn(3) // 0 = potion, 1 = gold, 2 = exp

	switch rewardType {
	case 0: // Potions
		nb := rand.Intn(2) + 1 // 1 à 2 potions
		c.Potions += nb
		fmt.Printf(Blue+"🧪 Vous obtenez %d potion(s) supplémentaire(s) !"+Reset+"\n", nb)
	case 1: // Gold
		gold := rand.Intn(30) + 20 // 20 à 50 gold
		c.Gold += gold
		fmt.Printf(Yellow+"💰 Vous trouvez %d or dans le coffre !"+Reset+"\n", gold)
	case 2: // Expérience
		exp := rand.Intn(15) + 10 // 10 à 25 exp
		c.ExpActuelle += exp
		fmt.Printf(Cyan+"⭐ Vous gagnez %d points d’expérience bonus !"+Reset+"\n", exp)
	}

	waitForEnter()
}
