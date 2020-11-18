//- Create a map with key type string and value type [] string.
//- Key must contain names in the form surname_name
//- Value must contain the person's favorite hobbies
//- Demonstrate all of these values and their indexes.
package main

import "fmt"

func main() {
	mepe := map[string][]string{
		"dasilva_guriaestranhadostrangerthings": []string{"desaparecer", "ser assustadora", "raspar o cabelo"},
		"senna_ayrton":                          []string{"andar de jetski", "ser milionário", "falar com sotaque de paulista mano"},
		"pike_rob":                              []string{"criar linguagens de programação", "usar uns ternos muito loucos"},
	}

	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
