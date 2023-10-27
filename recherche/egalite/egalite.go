package egalite

/*
On considère des ensembles de nombres représentés par des tableaux : on considère
que ces tableaux ne contiennent qu'une seule fois chaque nombre (puisqu'ils
représentent des ensembles) et les nombres ne sont pas nécessairement dans
l'ordre dans les tableaux.

On veut savoir si deux ensembles sont égaux ou pas, c'est-a-dire savoir si deux
tableaux contiennent les mêmes nombres ou pas. C'est à cette question que doit
répondre la fonction egalite.

# Entrées
- t1 : un tableau d'entiers (sans doublons) représentant un ensemble
- t2 : un tableau d'entiers (sans doublons) représentant un ensemble

# Sortie
- egaux : un booléen qui vaut true si t1 et t2 représentent le même ensemble et
          qui vaut false sinon

# Info
2021-2022, test2, exercice 3
*/

func egalite(t1, t2 []int) (egaux bool) {
	if len(t1)!=len(t2) {
		return false
	} else {
		var valeur bool
		for i:=0; i<len(t1); i++ {
			for j:=0; j<len(t2); j++ {
				if t1[i] == t2[j] {
					valeur = true
					break
				}
			}
			if !valeur {
				return false
			}
		}
		return true
	}
}
