// Structs embutidos
package main

import "fmt"

type harddisk struct {
	company string
	price   float64
	qtd     int
	avaible bool
}

type capacity struct {
	harddisk
	cap  int
	giga bool
	tera bool
}

func main() {

	hd2 := harddisk{
		company: "HP",
		price:   889.65,
		qtd:     156,
		avaible: true,
	}

	hd1 := capacity{
		harddisk: harddisk{
			company: "Samsung",
			price:   565.66,
			qtd:     156,
			avaible: true,
		},
		cap:  1,
		giga: false,
		tera: true,
	}

	fmt.Println(hd2)
	fmt.Println(hd1.company)

}
