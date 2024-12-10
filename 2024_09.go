package main

import (
	"fmt"
	"os"
)

func Ex9() {
	nome := "asd.txt"
	parte_09_1(nome)
	parte_09_2(nome)
}

func parte_09_1(nome string) {
	arq, _ := os.ReadFile(nome)

	var blocos []int
	var id int
	var espaco bool = true
	for _, r := range arq {
		if r < 48 || r >= 58 {
			continue
		}
		for i := r - 48; i > 0; i-- {
			if espaco {
				blocos = append(blocos, id)
			} else {
				blocos = append(blocos, -1)
			}
		}
		if espaco {
			id++
		}
		espaco = !espaco
	}
	menor := 0
	maior := len(blocos) - 1
	for {
		for blocos[menor] != -1 {
			menor++
			if menor >= len(blocos) || menor >= maior {
				break
			}
		}
		if menor >= maior {
			break
		}
		if blocos[maior] == -1 {
			maior--
			blocos = blocos[:len(blocos)-1]
			continue
		}
		blocos[menor] = blocos[maior]
		blocos = blocos[:len(blocos)-1]
		maior--
	}
	if blocos[len(blocos)-1] == -1 {
		blocos = blocos[:len(blocos)-1]
	}
	var final1 uint64 = 0
	for i, v := range blocos {
		final1 += uint64(i * v)
	}
	fmt.Printf("final1: %v\n", final1)
}

type Bloco struct {
	pos, tam, id int
}

func printarBlocos(blocos []int) {
	for _, v := range blocos {
		if v >= 0 {
			fmt.Printf("%v", v)
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Println()
}

func parte_09_2(nome string) {
	arq, _ := os.ReadFile(nome)

	var blocos []int
	var id int = 0
	var espaco bool = true
	var livres []Bloco
	var usados []Bloco
	for _, r := range arq {
		if r < 48 || r >= 58 {
			continue
		}
		blocoCont := Bloco{0, 0, 0}
		blocoCont.pos = len(blocos)
		blocoCont.id = id
		for i := r - 48; i > 0; i-- {
			blocoCont.tam++
			if espaco {
				blocos = append(blocos, id)
			} else {
				blocos = append(blocos, -1)
			}
		}
		if espaco {
			usados = append(usados, blocoCont)
			id++
		} else {
			blocoCont.id = -1
			livres = append(livres, blocoCont)
		}
		espaco = !espaco
	}
	ultimo := len(usados) - 1
	for ultimo > 0 {
		tam := usados[ultimo].tam
		for k, v := range livres {
			if v.tam >= tam {
				// printarBlocos(blocos)
				for i := range tam {
					blocos[v.pos+i] = usados[ultimo].id
					blocos[usados[ultimo].pos+i] = -1
				}
				livres[k].tam -= tam
				livres[k].pos += tam
				break
			}
		}
		ultimo--
	}
	// printarBlocos(blocos)
	var final2 uint64 = 0
	for k, v := range blocos {
		// fmt.Printf("[%v %v] ", k, v)
		if v != -1 {
			final2 += uint64(k * v)
		}
	}
	fmt.Printf("final2: %v\n", final2)
}
