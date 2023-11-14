package classer

import "fmt"

/*
La fonction classer doit trier un tableau d'étudiants (tels que définis) ci-dessous
de celui qui a la meilleur moyenne (la plus élevée) à celui qui a la plus mauvaise
moyenne. En cas d'égalité de moyenne entre deux étudiants, celui qui a le nom de
famille qui arrive en premier dans l'ordre alphabétique doit être classer avant
l'autre (on peut utiliser <, >, <=, >=, == pour comparer les chaînes de caractères
par ordre alphabétique). Si deux étudiants ont la même moyenne et le même nom, on
met en premier celui qui a le prénom qui est en premier dans l'ordre alphabétique.

# Entrée
- promo : le tableau d'étudiants à trier

# Sortie
- classement : un tableau contenant les mêmes étudiants mais trié

# Info
2021-2022, test3, exercice 5
*/

type etudiant struct {
	nom, prenom string
	moyenne     float64
}

func classer(promo []etudiant) (classement []etudiant) {
	var n int = len(promo)
	var max int
	for i := 0; i < n; i++ {
		max = i
		for j := i + 1; j < n; j++ {
			if promo[j].moyenne > promo[max].moyenne {
				max = j
			} else if promo[j].moyenne == promo[max].moyenne {
				if promo[j].nom < promo[max].nom {
					max = j
				} else if promo[j].nom == promo[max].nom {
					if promo[j].prenom < promo[max].prenom {
						max = j
					}
				}
			}
		}
		promo[i], promo[max] = promo[max], promo[i]
	}
	classement = promo
	fmt.Println(classement)
	return classement
}
