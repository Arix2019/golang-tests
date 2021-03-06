// map com 'key' de int e 'values' de string.

package main

import "fmt"

var clubes = map[int]string{
	1: "Bayern",
	2: "PSG",
	3: "Barcelona",
	4: "Real Madrid",
}

func addClubes(k int, v string) {
	clubes[k] = v
}

func main() {

	fmt.Println(clubes)
	fmt.Println(clubes[3])

	// Adicionados
	addClubes(5, "Manchester City")
	addClubes(6, "Borussia Dortmund")
	addClubes(7, "Porto")

	for cont := 0; cont <= len(clubes); cont++ {
		fmt.Println(clubes[cont])
	}

	fmt.Println(" ")

	//comma OK
	if test, ok := clubes[7]; ok {
		fmt.Println("Clube selecionado:", test)

	} else {
		fmt.Println("unknow value!")

	}

	fmt.Println("\nClube removido da lista:", clubes[5])
	delete(clubes, 5)

	fmt.Println(" ")

	fmt.Println("Lista de Clubes Europeus exibida com: loop for /'key'/'value':")
	for key, value := range clubes {
		fmt.Println(key, value)

	}

}
