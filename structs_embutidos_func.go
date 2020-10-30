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

func xhd1(c string, p float64, q int, a bool, ca int, g bool, t bool) {

	hd1 := capacity{
		harddisk: harddisk{
			company: c,
			price:   p,
			qtd:     q,
			avaible: a,
		},
		cap:  ca,
		giga: g,
		tera: t,
	}
	fmt.Println("\tFabricante:", hd1.company,
		"\n\tPreço:R$", hd1.price,
		"\n\tQuantidade:", hd1.qtd,
		"\n\tDisponível:", hd1.avaible,
		"\n\tCapacidade:", hd1.cap,
		"\n\tGigabyte:", hd1.giga,
		"\n\tTerabyte:", hd1.tera,
	)
}

var hd2 = capacity{
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

func fHd1() {

}

func fHd2() {
	fmt.Println("\tFabricante:", hd2.company,
		"\n\tPreço:R$", hd2.price,
		"\n\tQuantidade:", hd2.qtd,
		"\n\tDisponível:", hd2.avaible,
		"\n\tCapacidade:", hd2.cap,
		"\n\tGigabyte:", hd2.giga,
		"\n\tTerabyte:", hd2.tera,
	)

}

func main() {
	xhd1("Samsung", 850.55, 356, true, 2, false, true)
	fHd1()
	fmt.Println("")
	fHd2()

}