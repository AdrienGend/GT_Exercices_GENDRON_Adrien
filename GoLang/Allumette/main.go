package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Nombre d'allumettes : ")
	fmt.Scan(&n)

	if n < 4 {
		fmt.Println("Le nombre d'allumettes doit être au moins 4.")
		return
	}

	player := 1
	for n > 0 {
		fmt.Printf("Joueur %d, combien d'allumettes prenez-vous (1 à 3) ? ", player)
		var take int
		fmt.Scan(&take)

		if take < 1 || take > 3 {
			fmt.Println("Vous devez prendre entre 1 et 3 allumettes.")
			continue
		}

		n -= take
		if n <= 0 {
			fmt.Printf("Joueur %d a perdu.\n", player)
			return
		}

		player = 3 - player // Alterner entre les joueurs 1 et 2 (3 - 1 = 2, 3 - 2 = 1)
	}

	fmt.Println("Erreur : le nombre d'allumettes est négatif.")
}
