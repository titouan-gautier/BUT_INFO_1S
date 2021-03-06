package recherche

/*
La fonction recherche doit indiquer la valeur et la position du plus petit
entier strictement positif dans un tableau.

# Entrée
- tab : le tableau dans lequel on cherche

# Sorties
- trouve : un booléen qui vaut true s'il existe au moins un entier strictement
           positif dans tab et false sinon,
- pos : la position dans tab du plus petit entier strictement positif s'il existe
- val : la valeur du plus petit entier strictement positif de tab s'il existe
(pos et val auront la valeur par défaut d'un entier s'il n'existe pas d'entier
strictement positif dans tab)
*/

func recherche(tab []int) (trouve bool, pos, val int) {
	for i:= 0 ; i< len(tab) ; i++ {
		if tab[i] > 0 {
			trouve = true
			val = tab[i]
			for j:= 0 ; j < len(tab) ; j++ {
				if tab [j] < val && tab[j] > 0 {
					val = tab[j]
					pos = j
				}
			}
		}
	}
	return trouve, pos, val
}
