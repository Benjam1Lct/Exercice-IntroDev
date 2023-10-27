package doublons2

	/*
	On dispose d'un tableau d'entiers de longueur n et on suppose qu'il contient
	exactement une fois chaque nombre compris entre 1 et n. On voudrait vérifier
	cela. C'est le travail de la fonction doublons.

	Coder la fonction doublons de manière à ne parcourir le tableau tab qu'une seule
	fois rapportera plus de points.

	# Entrée
	- tab : un tableau d'entiers

	# Sortie
	- ok : un booléen qui vaut true si tab contient bien exactement une
	fois chaque entier entre 1 et len(tab) et false sinon

	# Info
	2021-2022, test 1, exercice 8
	*/

	func doublons(tab []int) (ok bool) {
		n := len(tab)
		
		// Créez un tableau de booléens pour suivre les éléments rencontrés
		seen := make([]bool, n)
		
		// Parcourez le tableau tab
		for _, num := range tab {
			// Vérifiez si le nombre est valide (dans la plage [1, n])
			if num < 1 || num > n {
				return false
			}
			
			// Si l'élément a déjà été rencontré, renvoyez false
			if seen[num-1] {
				return false
			}
			
			// Marquez l'élément comme rencontré
			seen[num-1] = true
		}
		
		// Vérifiez si tous les éléments de 1 à n ont été rencontrés
		for i := 0; i < n; i++ {
			if !seen[i] {
				return false
			}
		}
		
		// Si toutes les conditions sont satisfaites, renvoyez true
		return true
	}
	
