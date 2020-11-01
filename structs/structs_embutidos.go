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

	hd2 := capacity{
		harddisk: harddisk{
			company: "HP",
			price:   889.65,
			qtd:     205,
			avaible: true,
		},
		cap:  500,
		giga: true,
		tera: false,
	}

	fmt.Println("\tFabricante:", hd1.company,
		"\n\tPreço:R$", hd1.price,
		"\n\tQuantidade:", hd1.qtd,
		"\n\tDisponível:", hd1.avaible,
		"\n\tCapacidade:", hd1.cap,
		"\n\tGigabyte:", hd1.giga,
		"\n\tTerabyte:", hd1.tera,
	)

	fmt.Println("")

	fmt.Println("\tFabricante:", hd2.company,
		"\n\tPreço:R$", hd2.price,
		"\n\tQuantidade:", hd2.qtd,
		"\n\tDisponível:", hd2.avaible,
		"\n\tCapacidade:", hd2.cap,
		"\n\tGigabyte:", hd2.giga,
		"\n\tTerabyte:", hd2.tera,
	)

}
