package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func acharMais(recebido, cx, cy int, mapa [][]int, caminho []Pos, set *Set) int {
	if mapa[cx][cy] == 9 && recebido == 9 {
		set.add(Pos{x: cx, y: cy})
		return 1
	}
	if mapa[cx][cy] != recebido {
		return 0
	}
	valor := 0
	if cx > 0 {
		valor += acharMais(recebido+1, cx-1, cy, mapa, append(caminho, Pos{x: cx - 1, y: cy}), set)
	}
	if cy < len(mapa[0])-1 {
		valor += acharMais(recebido+1, cx, cy+1, mapa, append(caminho, Pos{x: cx, y: cy + 1}), set)
	}
	if cx < len(mapa)-1 {
		valor += acharMais(recebido+1, cx+1, cy, mapa, append(caminho, Pos{x: cx + 1, y: cy}), set)
	}
	if cy > 0 {
		valor += acharMais(recebido+1, cx, cy-1, mapa, append(caminho, Pos{x: cx, y: cy - 1}), set)
	}
	return valor
}

func Ex10() {
	arq, err := os.ReadFile("2024_10.txt")
	if err != nil {
		panic(err)
	}
	conteudoStr := strings.Fields(string(arq))
	var conteudo [][]int
	var inicios []Pos
	for x, l := range conteudoStr {
		var linha []int
		for y, r := range l {
			num, err := strconv.Atoi(string(r))
			if err != nil {
				linha = append(linha, -1)
			} else {
				linha = append(linha, num)
			}
			if r == '0' {
				inicios = append(inicios, Pos{x: x, y: y})
			}
		}
		conteudo = append(conteudo, linha)
	}
	// printarMapa10(conteudo)
	final1 := 0
	final2 := 0
	for _, v := range inicios {
		qtd := novoSet()
		final2 += acharMais(0, v.x, v.y, conteudo, []Pos{v}, &qtd)
		final1 += qtd.qtdItens()
	}
	fmt.Printf("final1: %v\n", final1)
	fmt.Printf("final2: %v\n", final2)
}
