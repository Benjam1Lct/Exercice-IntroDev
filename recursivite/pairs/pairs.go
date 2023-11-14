package pairs

/*
La fonction pairs doit indiquer combien il y a de nombres pairs dans un tableau donné en paramètre.
On peut remarquer que le nombre de nombres pairs dans un tableau t est la somme du nombre de nombres pairs dans
la première moitié de t et du nombre de nombres pairs dans la deuxième moitié de t.

Vous ne devez pas utiliser de boucles for dans cet exercice.

# Entrée
- t, le tableau dans lequel compter les nombres pairs

# Sortie
- num, le nombre de nombres pairs dans t

# Info
2022-2023, test 2, exercice 3
*/

func pairs(t []int) (num int) {
	if len(t) == 0 {
		return 0
	} else if len(t) == 1 {
		if t[0]%2 == 0 {
			return 1
		} else {
			return 0
		}
	} else {
		taille1 := t[:len(t)/2]
		taille2 := t[len(t)/2:]
		return pairs(taille1) + pairs(taille2)
	}
}
