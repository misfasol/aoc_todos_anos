package main

func main() {
	Ex12()
}

type Vec2 struct {
	x, y int
}

type Pos Vec2

func somarPos(a, b Pos) Pos {
	return Pos{x: a.x + b.x, y: a.y + b.y}
}

var ADJ4 = []Pos{Pos{x: -1, y: 0}, Pos{x: 0, y: 1}, Pos{x: 1, y: 0}, Pos{x: 0, y: -1}}

type Set struct {
	hash map[interface{}]nothing
}

type nothing struct{}

func novoSet() Set {
	return Set{hash: make(map[interface{}]nothing)}
}

func (s *Set) add(valor interface{}) {
	s.hash[valor] = nothing{}
}

func (s *Set) qtdItens() int {
	v := 0
	for range s.hash {
		v++
	}
	return v
}
