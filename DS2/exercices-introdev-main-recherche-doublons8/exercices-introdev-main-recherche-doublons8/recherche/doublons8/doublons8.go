package doublons8

import "fmt"

/*
On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient exactement p fois chaque nombre compris entre k et k + (n/p) - 1 (pas forcément dans l'ordre) pour un certain k et un certain p non connus à l'avance (p différent de 0). On voudrait vérifier que c'est bien le cas. C'est le travail de la fonction doublons.

# Entrée
- tab : un tableau d'entiers

# Sortie
- ok : un booléen qui doit valoir true si tab contient exactement p fois chaque entier de k à k + len(tab)/p - 1 et qui vaut false sinon

# Info
2023-2024, test 2, exercice 8
*/

func doublons(tab []int) (ok bool) {
	for i := 0; i < len(tab); i++ {
		for j := i + 1; j < len(tab); j++ {
			if tab[i] > tab[j] {
				tab[i], tab[j] = tab[j], tab[i]

			}
		}
	}
	fmt.Println(tab)

	max_present := 0
	for k := 0; k < len(tab); k++ {
		num_present := 1
		for i := 0; i < len(tab); i++ {
			if tab[k] == tab[i] && k != i {
				num_present += 1
			}
		}
		if num_present > max_present {
			max_present = num_present
		}

	}

	fmt.Println(max_present)

	var already_check []int = []int{}
	for i := 0; i < len(tab); i++ {
		verify := true
		for k := 0; k < len(already_check); k++ {
			if already_check[k] == i {
				verify = false
			}
		}
		if verify == true {
			counter := 1
			for j := i + 1; j < len(tab); j++ {
				if tab[i] == tab[j] {
					counter = counter + 1
				}
			}
			if counter != max_present {
				return false
			}
		}

	}

	return true
}
