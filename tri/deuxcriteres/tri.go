package tri

import "fmt"

/*
La fonction tri doit trier un tableau d'entiers de la manière suivante :
on trouvera d'abord tous les nombres strictement négatifs du tableau, dans l'ordre décroissant,
puis tous les nombres positifs ou nuls, dans l'ordre croissant.

# Entrée
- t, le tableau à trier (en place)

# Info
2022-2023, test 4, exercice 8
*/

func tri(t []int) {
	if t == nil {
		return
	} else {
		negatifs, positifs := determineNegative(t)
		for i := 0; i < len(negatifs)-1; i++ {
			for j := 0; j < len(negatifs)-i-1; j++ {
				if negatifs[j] < negatifs[j+1] {
					negatifs[j], negatifs[j+1] = negatifs[j+1], negatifs[j]
				}
			}
		}
		fmt.Println(negatifs, positifs)
		for i := 0; i < len(positifs)-1; i++ {
			for j := 0; j < len(positifs)-i-1; j++ {
				if positifs[j] < positifs[j+1] {
					positifs[j], positifs[j+1] = positifs[j+1], positifs[j]
				}
			}
		}

	}

}

func determineNegative(t []int) (positifs, negatifs []int) {
	for i := 0; i < len(t); i++ {
		if t[i] < 0 {
			negatifs = append(negatifs, t[i])
		} else {
			positifs = append(positifs, t[i])
		}
	}
	return negatifs, positifs
}
