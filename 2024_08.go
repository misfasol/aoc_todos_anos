package main

import (
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func Ex8() {
	pos := make(map[rune][]Pos)
	anti1 := make(map[Pos]bool)
	anti2 := make(map[Pos]bool)
	qtdAntenas := 0
	arq, _ := os.ReadFile("2024_08.txt")
	linhas := strings.Fields(string(arq))
	for i, linha := range linhas {
		for j, r := range linha {
			if r != '.' {
				pos[r] = append(pos[r], Pos{x: i, y: j})
				qtdAntenas++
			}
		}
	}
	for _, l := range pos {
		for _, v1 := range l {
			for _, v2 := range l {
				if v1 != v2 {
					x := v1.x - v2.x + v1.x
					y := v1.y - v2.y + v1.y
					if x >= 0 && x < len(linhas) && y >= 0 && y < len(linhas[0]) {
						anti1[Pos{x: x, y: y}] = true
					}
				}
			}
		}
	}
	for _, l := range pos {
		for _, v1 := range l {
			for _, v2 := range l {
				if v1 != v2 {
					dx := v1.x - v2.x
					x := dx + v1.x
					dy := v1.y - v2.y
					y := dy + v1.y
					for x >= 0 && x < len(linhas) && y >= 0 && y < len(linhas[0]) {
						anti2[Pos{x: x, y: y}] = true
						x += dx
						y += dy
					}
				}
			}
		}
	}

	final1 := len(anti1)
	final2 := len(anti2) + qtdAntenas
	for _, l := range pos {
		for _, p := range l {
			_, existe := anti2[p]
			if existe {
				final2--
			}
		}
	}
	fmt.Printf("final1: %v\nfinal2: %v\n", final1, final2)
}
