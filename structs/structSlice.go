package main

import (
	"fmt"
)

type filme struct {
	id       int
	title    string
	director string
	year     int
}

var filmes = []filme{
	filme{
		id:       0,
		title:    "Matrix",
		director: "Irmãos Wachowski",
		year:     1999,
	},
	filme{
		id:       1,
		title:    "Mad Max",
		director: "George Miller",
		year:     1979,
	},
	filme{
		id:       2,
		title:    "Nascido em 4 de Julho",
		director: "Oliver Stone",
		year:     1989,
	},
	filme{
		id:       3,
		title:    "Duro de Matar",
		director: "John McTiernan",
		year:     1988,
	},
}

func main() {
	fmt.Println(filmes)

}
