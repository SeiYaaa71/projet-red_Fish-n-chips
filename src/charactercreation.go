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
		fmt.Println("1. Humain (100 PV max)")
		fmt.Println("2. Elfe (80 PV max)")
		fmt.Println("3. Nain (120 PV max)")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			classe = "Humain"
		case 2:
			classe = "Elfe"
		case 3:
			classe = "Nain"
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
	case "Humain":
		c.PVMax = 100
	case "Elfe":
		c.PVMax = 80
	case "Nain":
		c.PVMax = 120
	}
	c.PVActuels = c.PVMax / 2

	return c
}
