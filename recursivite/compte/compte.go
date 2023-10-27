package compte

/*
La fonction compte doit indiquer combien de fois une valeur v apparaît dans un
tableau tab.

Vous ne devez pas utiliser de boucles, la fonction compte sera donc récursive.

Vous pouvez vous baser sur la remarque suivante : le nombre de fois que v apparaît
dans tab est la somme du nombre de fois où v apparaît dans la première moitié
de tab et du nombre de fois où v apparaît dans la deuxième moité de tab.

# Entrées
- tab : un tableau d'entiers
- v : la valeur à chercher

# Sortie
- num : le nombre de fois que v apparaît dans tab

# Info
2021-2022, test3, exercice 3
*/

func compte(tab []int, v int) (num int) {
	// Cas de base : tableau vide
    if len(tab) == 0 {
        return 0
    }

    // Cas de base : tableau avec un seul élément
    if len(tab) == 1 {
        if tab[0] == v {
            return 1
        } else {
            return 0
        }
    }

    // Diviser le tableau en deux moitiés
    middle := len(tab) / 2
    leftHalf := tab[:middle]
    rightHalf := tab[middle:]

    // Compter les occurrences de v dans les deux moitiés
    countLeft := compte(leftHalf, v)
    countRight := compte(rightHalf, v)

    // Retourner la somme des comptages des deux moitiés
    return countLeft + countRight
}
