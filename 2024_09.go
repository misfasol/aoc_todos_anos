package main

import (
	"fmt"
	"os"
)

func Ex9() {
	arq, _ := os.ReadFile("2024_09.txt")
	var blocos []int
	var id int
	var espaco bool = true
	for _, r := range arq {
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
	// for _, v := range blocos {
	// 	if v == -1 {
	// 		fmt.Print(".")
	// 	} else {
	// 		fmt.Print(v)
	// 	}
	// }
	// fmt.Println()
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
		// fmt.Printf("%2v %2v %2v %2v ", menor, maior, blocos[menor], blocos[maior])
		blocos[menor] = blocos[maior]
		blocos = blocos[:len(blocos)-1]
		maior--
	}
	// for _, v := range blocos {
	// 	if v == -1 {
	// 		fmt.Print(".")
	// 	} else {
	// 		fmt.Print(v)
	// 	}
	// }
	// fmt.Println()
	final1 := 0
	for i, v := range blocos {
		final1 += i * v
	}
	fmt.Printf("final1: %v\n", final1)
}

// 1928
