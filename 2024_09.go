package main

import (
	"fmt"
	"os"
)

func Ex9() {
	parte_09_1("asd.txt")
	parte_09_2("asd.txt")
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

func parte_09_2(nome string) {
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
	fmt.Printf("final2: %v\n", final1)
}
