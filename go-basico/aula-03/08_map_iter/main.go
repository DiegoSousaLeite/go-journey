package main

import "fmt"

func main() {
    
	mapa := map[string]int{"Alice": 25, "Bob": 30}

	//Range para iterar sobre o mapa
	//Range retorna dois valores: chave e valor
	//Range serve para iterar sobre arrays, slices, strings, mapas e canais
	for key, value := range mapa {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
}
