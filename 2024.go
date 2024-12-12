package main

func main() {
	Ex12()
}

type Vec2 struct {
	x, y int
}

type Pos Vec2

type Set struct {
	hash map[interface{}]bool
}

func novoSet() Set {
	return Set{hash: make(map[interface{}]bool)}
}

func (s *Set) add(valor interface{}) {
	s.hash[valor] = true
}

// func (s *Set) existe(valor interface{}) bool {
// 	_, existe := s.hash[valor]
// 	return existe
// }

func (s *Set) qtdItens() int {
	v := 0
	for range s.hash {
		v++
	}
	return v
}
