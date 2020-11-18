//- Using the previous exercise, remove an entry from the map and demonstrate the entire map using range.
package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	mepe := map[string][]string{
		"dasilva_guriaestranhadostrangerthings": []string{"desaparecer", "ser assustadora", "raspar o cabelo"},
		"senna_ayrton":                          []string{"andar de jetski", "ser milionário", "falar com sotaque de paulista mano"},
		"pike_rob":                              []string{"criar linguagens de programação", "usar uns ternos muito loucos"},
	}
	mepe["loureiro_kiko"] = []string{"usar os trequinho no punho", "tacar fogo na guitarra"}
	delete(mepe, "dasilva_guriaestranhadostrangerthings")
	for k, v := range mepe {
		fmt.Println(k)
		for i, h := range v {
			fmt.Println("\t", i, " - ", h)
		}
	}
}
