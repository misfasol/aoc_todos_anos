package main

import (
	"fmt"
	"os"
	"strings"
)

type Plot struct {
	x, y, id int
	run      rune
}

func dia12PrintarMapa(mapa [][]Plot) {
	for _, l := range mapa {
		for _, p := range l {
			fmt.Printf("%c", p.run)
		}
		fmt.Println()
	}
}

func Ex12() {
	inputBytes, err := os.ReadFile("2024_12.txt")
	if err != nil {
		panic(err)
	}
	inputStr := strings.Fields(string(inputBytes))
	var conteudo [][]Plot
	for x, l := range inputStr {
		var linha []Plot
		for y, r := range l {
			linha = append(linha, Plot{x: x, y: y, id: -1, run: r})
		}
		conteudo = append(conteudo, linha)
	}
	final1 := 0
	for x, linha := range conteudo {
		for y, p := range linha {
			if p.id == -1 {
				conteudo[x][y].id = x*len(conteudo) + y
				a, p := acharReg(x, y, conteudo)
				final1 += a * p
			}
		}
	}
	fmt.Printf("final1: %v\n", final1)
}

func acharReg(x, y int, mapa [][]Plot) (int, int) {
	plot := mapa[x][y]
	area := 1
	per := 4
	if x > 0 {
		plotCima := mapa[x-1][y]
		if plot.run == plotCima.run {
			per -= 1
			if plotCima.id == -1 {
				// fmt.Println(plot, plotCima)
				mapa[x-1][y].id = plot.id
				ao, po := acharReg(x-1, y, mapa)
				area += ao
				per += po
			}
		}
	}
	if y < len(mapa[0])-1 {
		plotDir := mapa[x][y+1]
		if plot.run == plotDir.run {
			per -= 1
			if plotDir.id == -1 {
				// fmt.Println(plot, plotDir)
				mapa[x][y+1].id = plot.id
				ao, po := acharReg(x, y+1, mapa)
				area += ao
				per += po
			}
		}
	}
	if x < len(mapa)-1 {
		plotBaixo := mapa[x+1][y]
		if plot.run == plotBaixo.run {
			per -= 1
			if plotBaixo.id == -1 {
				// fmt.Println(plot, plotBaixo)
				mapa[x+1][y].id = plot.id
				ao, po := acharReg(x+1, y, mapa)
				area += ao
				per += po
			}
		}
	}
	if y > 0 {
		plotEsq := mapa[x][y-1]
		if plot.run == plotEsq.run {
			per -= 1
			if plotEsq.id == -1 {
				// fmt.Println(plot, plotEsq)
				mapa[x][y-1].id = plot.id
				ao, po := acharReg(x, y-1, mapa)
				area += ao
				per += po
			}
		}
	}
	// fmt.Printf("ret a: %v | p: %v\n", area, per)
	return area, per
}
