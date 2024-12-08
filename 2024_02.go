package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Ex2() {
	dados, err := os.ReadFile("2024_02.txt")
	if err != nil {
		panic(err)
	}
	conteudo := string(dados)
	linhas := strings.FieldsFunc(conteudo, func(r rune) bool {
		return r == '\n'
	})
	var numeros [][]int
	for _, i := range linhas {
		nstr := strings.Fields(i)
		var nint []int
		for _, v := range nstr {
			conv, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			nint = append(nint, conv)
		}
		numeros = append(numeros, nint)
	}
	checar := func(lista []int) bool {
		nova := make([]int, len(lista))
		copy(nova, lista)
		slices.Reverse(nova)
		if !slices.IsSorted(lista) && !slices.IsSorted(nova) {
			return false
		}
		for k := range len(lista) - 1 {
			if math.Abs(float64(lista[k+1]-lista[k])) < 1 || math.Abs(float64(lista[k+1]-lista[k])) > 3 {
				return false
			}
		}
		return true
	}
	contador := 0
	dois := 0
	for _, v := range numeros {
		// criar variacoes
		for i := 0; i < len(v); i++ {
			vari := slices.Concat(v[:i], v[i+1:])
			if checar(vari) {
				dois++
				break
			}
		}
		if checar(v) {
			contador++
		}
	}
	fmt.Printf("cont: %v\n", contador)
	fmt.Printf("dois: %v\n", dois)
}
