package main

import (
	"fmt"
)

type Grenal struct {
	golsInter  int
	golsGremio int
	vitInter   int
	vitGremio  int
	empate     int
	totalJogos int
}

func (g *Grenal) jogou() {
	var inter, gremio int
	fmt.Scanln(&inter, &gremio)
	g.golsInter += inter
	g.golsGremio += gremio
	g.totalJogos++
	if inter > gremio {
		g.vitInter++
	} else if gremio > inter {
		g.vitGremio++
	} else {
		g.empate++
	}
}

func (g Grenal) resultados() {
	fmt.Printf("%d grenais\n", g.totalJogos)
	fmt.Printf("Inter:%d\n", g.vitInter)
	fmt.Printf("Gremio:%d\n", g.vitGremio)
	fmt.Printf("Empates:%d\n", g.empate)
	if g.vitInter > g.vitGremio {
		fmt.Println("Inter venceu mais")
	} else if g.vitGremio > g.vitInter {
		fmt.Println("Gremio venceu mais")
	} else {
		fmt.Println("Nao houve vencedor")
	}
}

func main() {
	var g Grenal
	for {
		g.jogou()
		var novoGrenal int
		fmt.Println("Novo grenal (1-sim 2-nao)")
		fmt.Scanln(&novoGrenal)
		if novoGrenal == 2 {
			break
		}
	}
	g.resultados()
}
