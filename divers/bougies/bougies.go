package bougies

/*
La fonction bougies doit indiquer, pour une date donnée, combien de bougies il y avait 
sur le dernier gateau d'anniversaire d'une personne née le 17 avril 1986. 
Si jamais la date en question est un 17 avril, on donnera le nombre de bougies sur le 
gateau à cette date. Avant son premier anniversaire, une personne n'a jamais eu de 
gateau d'anniversaire et, donc, a toujours eu 0 bougies sur ses gateaux d'anniversaire.

# Entrées
- année : un entier représentant l'année
- mois : un entier représentant le numéro d'un mois dans l'année
- jour : un entier représentant le numéro d'un jour dans le mois

# Sortie
- numBougies : le nombre de bougies sur le dernier gateau d'anniversaire d'une personne 
née le 17 avril 1986 si on est à la date définie par jour, mois et année

# Remarque
On considèrera que la date donnée en entrée est toujours valide (pas de jour 35 ou de 
	mois 52)

# Exemples
bougies(23, 5, 1990) = 4
bougies(17, 4, 1990) = 4
bougies(25, 2, 1990) = 3

# Info
2022-2023, test 1, exercice 5
*/

func bougies(jour, mois, annee uint) (numBougies uint) {
	if mois == 4 {
		if annee >= 1986 {
			if jour >= 17 {
				return annee-1986
			} else {
				return 0
			}
		} else if annee < 1986 {
			return 0
		}
	} else if mois < 4 {
		if annee == 1986 {
			return 0
		} else if annee > 1986 {
			return annee-1986-1
		}
	} else if mois > 4 {
		if annee >= 1986 {
			return annee-1986
		}
	}
	return 0
}
