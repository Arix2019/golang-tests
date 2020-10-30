package main

import "fmt"

func main() {
	tel := map[string]int{
		"Joana":    551155685417,
		"Michelle": 551166874112,
		"Antonia":  551177847425,
		"Sakura":   551188547874,
	}
	fmt.Println("Tel:", tel["Sakura"])
	
	//adicionando
	tel["Renata"] = 551199874114
	fmt.Println(tel)
}
