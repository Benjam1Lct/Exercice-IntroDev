package suitemoinssimple

/*
On considère la suite dont le terme u(n) est défini par
- u(0)=1
- si n > 0 et 100 est divisible par u(n-1) alors u(n) vaut u(n-1) - 1
- sinon u(n) est le reste de la division de u(n-1) + 11 par 100
La fonction terme calcule les termes de cette suite.

# Entrée
- n : un entier indiquant le numéro du terme à calculer

# Sortie
- un : le terme u(n) de la suite

# Remarque
Les boucles ne sont pas autorisées pour résoudre cet exercice
À toute fin utile, on signale que 100 n'est pas divisible par 0
Votre fonction devra pouvoir calculer U(1000) en moins de 2 secondes

# Info
2022-2023, test 4, exercice 6
*/

func terme(n int) (un int) {
	memo := make(map[int]int)

	var computeTerm func(int) int
	computeTerm = func(n int) int {
		if val, ok := memo[n]; ok {
			return val
		}

		var un int
		if n == 0 {
			un = 1
		} else {
			precedent := computeTerm(n - 1)
			if precedent > 0 && 100%precedent == 0 {
				un = precedent - 1
			} else {
				un = (precedent + 11) % 100
			}
		}

		memo[n] = un
		return un
	}

	return computeTerm(n)
}
