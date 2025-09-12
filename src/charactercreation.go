package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Fonction pour créer un personnage personnalisé
func characterCreation() Character {
	reader := bufio.NewReader(os.Stdin)

	// --- Nom ---
	var nom string
	for {
		fmt.Print("Entrez le nom de votre personnage : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Vérifie que le nom contient uniquement des lettres
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

		// Mise en forme : première lettre majuscule, reste en minuscule
		nom = strings.Title(strings.ToLower(input))
		break
	}

	// --- Classe ---
	var classe string
	var pvMax int
	for {
		fmt.Println("\nChoisissez une classe :")
		fmt.Println("1. Requin (100 PV max)")
		fmt.Println("2. Poisson Clown (80 PV max)")
		fmt.Println("3. Orque (120 PV max)")
		fmt.Print("Votre choix : ")

		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			classe = "Requin"
			pvMax = 100
		case 2:
			classe = "Poisson Clown"
			pvMax = 80
		case 3:
			classe = "Requin"
			pvMax = 120
		default:
			fmt.Println("❌ Choix invalide, réessayez.")
			continue
		}
		break
	}

	// --- Initialisation du personnage ---
	return initCharacter(
		nom,
		classe,
		1,                          // Niveau
		pvMax,                      // PV max
		pvMax/2,                    // PV actuels = 50% PV max
		[]string{"Potion"},         // Inventaire
		[]string{"Coup de Poing"},  // Sort de base
	)
}
