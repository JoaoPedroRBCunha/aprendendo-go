package main

import (
	"fmt"
)

func diasDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}
}

func diasDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		// fallthrough - utilizado para que seja executado/apresentado/armazenado seja o próximo item do switch
		fallthrough // nesse caso o diaDaSemana armazenado seria "Segunda"
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número inválido"
	}
	return diaDaSemana
}

func main() {
	dia := diasDaSemana(200)
	fmt.Println(dia)

	dia2 := diasDaSemana2(5)
	fmt.Println(dia2)
}
