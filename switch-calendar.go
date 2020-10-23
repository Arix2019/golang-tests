package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// global variavel
var currentTime = time.Now() //obtem a data corrente

func main() {
	year := currentTime.Format("2006") //ano corrente
	conv, err := strconv.Atoi(year)    //converte a data str p/ int
	if err != nil {
		log.Fatal(err)
	}

	for mes := 1; mes <= 12; mes++ {
		fmt.Printf("Mês:%v\n", mes)
		switch {
		case mes == 12:
			for dia := 1; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 11:
			for dia := 1; dia <= 30; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 10:
			for dia := 1; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 9:
			for dia := 1; dia <= 30; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 8:
			for dia := 1; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}
		case mes == 2:
			// testa se o ano é bissexto ou não
			if conv%4 == 0 {
				for dia := 1; dia <= 29; dia++ {
					fmt.Print(" ", dia)
				}

			} else {
				for dia := 1; dia <= 28; dia++ {
					fmt.Print(" ", dia)
				}

			}
		case mes%2 == 0:
			for dia := 1; dia <= 30; dia++ {
				fmt.Print(" ", dia)
			}
		case mes%2 != 0:
			for dia := 1; dia <= 31; dia++ {
				fmt.Print(" ", dia)
			}
		}
		fmt.Println(" ")
	}
}
