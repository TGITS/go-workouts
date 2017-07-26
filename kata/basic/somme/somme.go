package main

import "fmt"

func main() {
	fmt.Printf("Programme de calcul de la somme des entiers de 1 à N\n")
	fmt.Printf("Entrer une valeur pour N (valeur supérieure à 1)\n")
	var N int
	fmt.Scan(&N)
	for N < 1 {
		fmt.Printf("Entrer une valeur pour N (valeur supérieure à 1)\n")
		fmt.Scan(&N)
	}
	somme := 0
	for i := 1; i <= N; i++ {
		somme = somme + i
		fmt.Printf("Somme intermédiare des entiers de 1 à %d : %d\n", i, somme)
	}
	fmt.Printf("Somme des entiers de 1 à %d : %d\n", N, somme)
}
