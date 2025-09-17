package main

import "fmt"

func applyEffects(c *Character, m *Monster) {
    // Effets sur le joueur
    newEffets := []Effect{}
    for _, e := range c.Effets {
        if e.Tours > 0 {
            switch e.Nom {
            case "Poison":
                c.PVActuels -= e.Degats
                fmt.Printf(Red+"☠ %s subit %d dégâts de poison ! (PV : %d/%d)\n"+Reset,
                    c.Nom, e.Degats, c.PVActuels, c.PVMax)
            }
            e.Tours--
            if e.Tours > 0 {
                newEffets = append(newEffets, e)
            }
        }
    }
    c.Effets = newEffets

    // Effets sur le monstre
    newEffetsM := []Effect{}
    for _, e := range m.Effets {
        if e.Tours > 0 {
            switch e.Nom {
            case "Poison":
                m.PVActuels -= e.Degats
                fmt.Printf(Red+"☠ %s subit %d dégâts de poison ! (PV : %d/%d)\n"+Reset,
                    m.Nom, e.Degats, m.PVActuels, m.PVMax)
            }
            e.Tours--
            if e.Tours > 0 {
                newEffetsM = append(newEffetsM, e)
            }
        }
    }
    m.Effets = newEffetsM
}

// Permet de lancer une potion de poison sur soi-même
func usePoison(c *Character) {
	c.Effets = append(c.Effets, Effect{
		Nom:   "Poison",
		Tours:  3,
		Degats: 5,
	})
	fmt.Println(Red + "💀 Vous êtes empoisonné !" + Reset)
}

// Permet de lancer une potion de poison sur le monstre
func throwPoisonOnMonster(m *Monster) {
	m.Effets = append(m.Effets, Effect{
		Nom:   "Poison",
		Tours:  3,
		Degats: 5,
	})
	fmt.Println(Red + "💀 Le monstre est empoisonné !" + Reset)
}
