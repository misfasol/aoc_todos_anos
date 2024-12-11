package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func Ex11() {
	nome := "2024_11.txt"
	parte_11_1(nome)
	parte_11_2(nome)
}

type PeVe struct {
	num, vez int
}

var cache map[PeVe]int = make(map[PeVe]int)

func parte_11_2(nome string) {
	contStr, err := os.ReadFile(nome)
	if err != nil {
		panic(err)
	}
	contStrSep := strings.Fields(string(contStr))
	var nums []int
	for _, v := range contStrSep {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	final2 := 0
	for _, v := range nums {
		final2 += avaliar(v, 75)
	}
	fmt.Printf("final2: %v\n", final2)
}

// recebe um numero e qtd d vezes para avaliar, retorna o tamanho da lista final
func avaliar(n, vezes int) int {
	cac, existe := cache[PeVe{num: n, vez: vezes}]
	if existe {
		return cac
	}
	if vezes == 0 {
		return 1
	}
	valor := 0
	qtd := qtdNums(n)
	if n == 0 {
		valor = avaliar(1, vezes-1)
	} else if qtd%2 == 0 {
		ad := int(math.Pow10(qtd / 2))
		valor = avaliar(n/ad, vezes-1) + avaliar(n%ad, vezes-1)
	} else {
		valor = avaliar(n*2024, vezes-1)
	}

	cache[PeVe{num: n, vez: vezes}] = valor
	return valor
}

func parte_11_1(nome string) {
	contStr, err := os.ReadFile(nome)
	if err != nil {
		panic(err)
	}
	contStrSep := strings.Fields(string(contStr))
	var nums []int
	for _, v := range contStrSep {
		n, _ := strconv.Atoi(v)
		nums = append(nums, n)
	}
	var final []int
	final = append(final, nums...)
	for range 25 {
		var novo []int
		for _, v := range final {
			novo = append(novo, umAvaliar(v)...)
		}
		final = []int{}
		final = append(final, novo...)
	}
	fmt.Printf("final1: %v\n", len(final))
}

func umAvaliar(n int) []int {
	if n == 0 {
		return []int{1}
	}
	qtdd := 0
	i := n
	for i != 0 {
		i /= 10
		qtdd++
	}
	if qtdd%2 == 0 {
		ad := int(math.Pow10((qtdd / 2)))
		return []int{n / ad, n % ad}
	}
	return []int{n * 2024}
}

func qtdNums(n int) int {
	qtd := 0
	for n != 0 {
		n /= 10
		qtd++
	}
	return qtd
}
