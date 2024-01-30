package main

import (
	"fmt"
	"math/rand"
)

type Cell struct {
	up, right, down, left bool
	visited               bool
}

type Land struct {
	length, height int
	cells          []Cell
	visited_cells  int
}

const (
	up = iota
	right
	down
	left
)

func main() {
	land := Land{}
	land.length = 10
	land.height = 10
	land.visited_cells = 0
	land.cells = make([]Cell, land.height*land.length)
	for i := 0; i < land.height*land.length; i++ {
		land.cells[i] = Cell{up: true, right: true, down: true, left: true, visited: false}
	}
	land.Generate()

	land.Draw()

}

func (l *Land) Generate() {
	inity, initx := 0, 0
	path := make([]int, 0)
	l.cells[inity*l.length+initx].visited = true
	i, j := inity, initx
	l.visited_cells++
	for l.visited_cells < l.height*l.length {
		if l.have_unvisited_neighbour(i, j) {
			unvisited := l.unvisited_neighbour(i, j)
			pos := rand.Int() % len(unvisited)
			switch unvisited[pos] {
			case up:
				l.cells[(i-1)*l.length+j].down = false
				l.cells[(i-1)*l.length+j].visited = true
				l.cells[(i)*l.length+j].up = false
				i -= 1
				l.visited_cells++
				path = append(path, up)
				break

			case right:
				l.cells[(i)*l.length+j+1].left = false
				l.cells[(i)*l.length+j+1].visited = true
				l.cells[(i)*l.length+j].right = false
				j += 1
				l.visited_cells++
				path = append(path, right)
				break

			case down:
				l.cells[(i+1)*l.length+j].up = false
				l.cells[(i+1)*l.length+j].visited = true
				l.cells[(i)*l.length+j].down = false
				i += 1
				l.visited_cells++
				path = append(path, down)
				break

			case left:
				l.cells[(i)*l.length+j-1].right = false
				l.cells[(i)*l.length+j-1].visited = true
				l.cells[(i)*l.length+j].left = false
				j -= 1
				l.visited_cells++
				path = append(path, left)
				break

			}
		} else {
			switch path[len(path)-1] {
			case up:
				i += 1
				break

			case right:
				j -= 1
				break

			case down:
				i -= 1
				break

			case left:
				j += 1
				break
			}

			path = path[:len(path)-1]
		}
	}
}

func (l *Land) unvisited_neighbour(i, j int) []int {
	return_array := make([]int, 0)
	// Cima

	if l.is_acessable(i-1, j) {
		if !l.cells[(i-1)*l.length+j].visited {
			return_array = append(return_array, up)
		}
	}
	// Direita
	if l.is_acessable(i, j+1) {
		if !l.cells[(i)*l.length+j+1].visited {
			return_array = append(return_array, right)
		}

	}

	// Baixo
	if l.is_acessable(i+1, j) {
		if !l.cells[(i+1)*l.length+j].visited {
			return_array = append(return_array, down)
		}
	}

	// Esquerda
	if l.is_acessable(i, j-1) {
		if !l.cells[(i)*l.length+j-1].visited {
			return_array = append(return_array, left)
		}
	}

	return return_array
}

func (l *Land) have_unvisited_neighbour(i, j int) bool {
	// Cima
	if l.is_acessable(i-1, j) && !l.cells[(i-1)*l.length+j].visited {
		return true
	}

	// Direita
	if l.is_acessable(i, j+1) && !l.cells[(i)*l.length+j+1].visited {
		return true
	}

	// Baixo
	if l.is_acessable(i+1, j) && !l.cells[(i+1)*l.length+j].visited {
		return true
	}

	// Esquerda
	if l.is_acessable(i, j-1) && !l.cells[(i)*l.length+j-1].visited {
		return true
	}
	return false

}

func (l *Land) is_acessable(i, j int) bool {
	if i < 0 || i >= l.height || j < 0 || j >= l.length {
		return false

	}
	return true
}

func (l *Land) PrettySquare() {
	fmt.Print("#")
	for j := 0; j < l.length; j++ {
		fmt.Print("##")
	}
	fmt.Println()
	for i := 0; i < l.height; i++ {
		fmt.Print("# ")
		for j := 0; j < l.length; j++ {
			if l.cells[i*l.length+j].right {
				fmt.Print("# ")
			} else {

				fmt.Print("  ")
			}

		}
		fmt.Println()
		fmt.Print("#")
		for j := 0; j < l.length; j++ {
			if l.cells[i*l.length+j].down {
				fmt.Print("##")
			} else {

				fmt.Print(" #")
			}
		}
		fmt.Println()
	}

}

func (l *Land) Draw() {
	fmt.Print("*")
	for j := 0; j < l.length; j++ {
		fmt.Print(" - *")
	}
	fmt.Println()

	for i := 0; i < l.height; i++ {
		fmt.Print("|  ")
		for j := 0; j < l.length; j++ {
			if l.cells[i*l.length+j].right {
				fmt.Print(" |  ")
			} else {

				fmt.Print("    ")
			}

		}
		fmt.Println()
		fmt.Print("*")
		for j := 0; j < l.length; j++ {
			if l.cells[i*l.length+j].down {
				fmt.Print(" - *")
			} else {

				fmt.Print("   *")
			}
		}
		fmt.Println()
	}

}
