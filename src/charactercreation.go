package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func characterCreation() *Character {
	reader := bufio.NewReader(os.Stdin)

	// --- Nom ---
	var nom string
	for {
		fmt.Print("Entrez le nom de votre personnage : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		valid := true
		for _, r := range input {
			if !unicode.IsLetter(r) {
				valid = false
				break
			}
		}

		if !valid || len(input) == 0 {
			fmt.Println("❌ Le nom doit contenir uniquement des lettres. Réessayez.")
			continue
		}

		nom = strings.Title(strings.ToLower(input))
		break
	}

	// --- Classe ---
	var classe string
	for {
		fmt.Println("\nChoisissez une classe :")
		fmt.Println(Red + "1. Assassin (100 PV max)" + Reset)
		fmt.Println(Blue + "2. Mage (80 PV max)" + Reset)
		fmt.Println(Purple + "3. Nécromancien (120 PV max)" + Reset)
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			classe = "Assassin"
		case 2:
			classe = "Mage"
		case 3:
			classe = "Nécromancien"
		default:
			fmt.Println("❌ Choix invalide, réessayez.")
			continue
		}
		break
	}

	// Crée le perso
	c := initCharacter(nom, classe)

	// Ajuste ses PV selon la classe
	switch classe {
	case "Assassin":
		c.PVMax = 100
	case "Mage":
		c.PVMax = 80
	case "Nécromancien":
		c.PVMax = 120
	}
	c.PVActuels = c.PVMax

	return c
}
