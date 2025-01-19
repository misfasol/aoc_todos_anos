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

func Ex12() {
	inputBytes, err := os.ReadFile("asd.txt")
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
	fmt.Println()
	for _, l := range conteudo {
		fmt.Print(" ")
		for _, r := range l {
			fmt.Printf("%c", r.run)
		}
		fmt.Println()
	}
	final1 := 0
	final2 := 0
	for x, linha := range conteudo {
		for y, p := range linha {
			if p.id == -1 {
				conteudo[x][y].id = x*len(conteudo) + y
				pontas := novoSet()
				a, p := acharReg(x, y, conteudo, pontas)
				fmt.Printf("verts [%c]: %v\n", conteudo[x][y].run, pontas)
				final1 += a * p
				final2 += a * pontas.qtdItens()
			}
		}
	}
	fmt.Printf("final1: %v\n", final1)
	fmt.Printf("final2: %v\n", final2)
}

func acharReg(x, y int, mapa [][]Plot, pontas Set) (int, int) {
	pontas.add(Pos{x: x, y: y})
	plot := mapa[x][y]
	area := 1
	per := 4
	for _, v := range ADJ4 {
		np := somarPos(Pos{x: x, y: y}, v)
		if np.x < 0 || np.y < 0 || np.x >= len(mapa) || np.y >= len(mapa[0]) {
			continue
		}
		plotLado := mapa[np.x][np.y]
		if plot.run == plotLado.run {
			per--
			if plotLado.id == -1 {
				mapa[np.x][np.y].id = plot.id
				ao, po := acharReg(np.x, np.y, mapa, pontas)
				area += ao
				per += po
			}
		}
	}
	foraBound := func(p Pos, n, m int) bool {
		if p.x < 0 || p.y < 0 || p.x >= n || p.y >= m {
			return true
		}
		return false
	}
	pos4menos := []Pos{{x: -1, y: -1}, {x: -1, y: 0}, {x: 0, y: -1}, {x: 0, y: 0}}
	pos4mais := []Pos{{x: 0, y: 0}, {x: 0, y: 1}, {x: 1, y: 0}, {x: 1, y: 1}}
	for _, v := range pos4menos {
		np := somarPos(Pos{x: x, y: y}, v)
		qtdFora := 0
		iguais := []bool{false, false, false, false}
		qtdIguais := 0
		for k, p := range pos4mais {
			ncp := somarPos(np, p)
			fmt.Println(ncp)
			if !foraBound(ncp, len(mapa), len(mapa[0])) {
				if plot.id == mapa[ncp.x][ncp.y].id {
					iguais[k] = true
					qtdIguais++
				}
			} else {
				qtdFora++
			}
		}
		if qtdFora == 3 {
			pontas.add(np)
		} else if qtdFora == 0 {
			if qtdIguais == 1 || qtdIguais == 3 {
				pontas.add(np)
			} else if qtdIguais == 2 {
				pos1, pos2 := func() (int, int) {
					a := []int{}
					for k, v := range iguais {
						if v {
							a = append(a, k)
						}
					}
					return a[0], a[1]
				}()
				if pos1%2 != pos2%2 {
					pontas.add(np)
				}
			}
		}
	}
	return area, per
}
