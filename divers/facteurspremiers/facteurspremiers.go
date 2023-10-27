package facteurspremiers

/*
La fonction facteursPremiers doit retourner un tableau contenant 
la liste de tous
les facteurs premiers d'un entier n, doublons compris

# Entrée
- n : un nombre entier positif

# Sortie
- facteurs : un tableau contenant tous les facteurs premiers de n, 
si n vaut 0 il
faut retourner un tableau à zéro éléments.

# Exemple
premiers(24) = [2 2 2 3] (l'ordre n'a pas d'importance)
*/
func facteursPremiers(n uint) (facteurs []uint) {
	if n<2 {
		return []uint{}
	}

	for n%2 == 0 {
        facteurs = append(facteurs, 2)
        n /= 2
    }
    
    // Maintenant, n est impair, donc nous pouvons itérer à travers les nombres impairs
    for i := uint(3); i*i <= n; i += uint(2) {
        for n%i == 0 {
            facteurs = append(facteurs, i)
            n /= i
        }
    }
    
    // Si n est maintenant plus grand que 1, cela signifie qu'il est lui-même un facteur premier
    if n > 1 {
        facteurs = append(facteurs, n)
    }
    
    return facteurs
	
}
