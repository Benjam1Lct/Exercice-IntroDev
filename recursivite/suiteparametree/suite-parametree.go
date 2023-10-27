package suiteparam

/*
On considère un ensemble de suites définies par U(n) = 2 * U(n-1) + U(0), chaque U(0) possible définissant une suite de cet ensemble. 
La fonction termeparam doit calculer le n-ième terme de la suite telle que U(0) = n.

Pour cet exercice, les boucles for sont interdites.

# Entrée
- n, un entier

# Sortie
- un, le terme U(n) de la suite telle que U(0) = n

# Info
2022-2023, test 2, exercice 6
*/

func termeparam(n int) (un int) {
	return terme(n, n)
}

func terme(n int, u0 int) (un int) {
	if n == 0 {
		return u0
	} else {
		return 2*terme(n-1, u0) + u0
	}
}
