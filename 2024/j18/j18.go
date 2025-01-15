package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed j18_inputTest.txt
var inputTest string

//go:embed j18_inputDay.txt
var inputDay string

type Pos struct {
	x, y int
}

type Node struct {
	f, g, h int
	Pos
	parent *Node
}

func parseCorrupted(input string) []Pos {
	var positions []Pos
	var lines = strings.Split(input, "\n")
	for _, line := range lines {
		var position = strings.Split(line, ",")
		if len(position) < 2 {
			continue
		}
		var y, _ = strconv.Atoi(position[0])
		var x, _ = strconv.Atoi(position[1])

		positions = append(positions, Pos{x, y})
	}
	return positions
}

func manhattanDistance(point, goal Pos) int {
	return abs(point.x-goal.x) + abs(point.y-goal.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findLeastF(open map[Pos]Node) Node {
	var minNode Node
	var minF = math.MaxInt
	for _, node := range open {
		if node.f < minF {
			minF = node.f
			minNode = node
		}
	}
	return minNode
}

func findSuccessors(node Node, sizeGrid int, corrupted []Pos) []Node {
	var successors []Node
	var directions = []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, dir := range directions {
		newX, newY := node.x+dir.x, node.y+dir.y
		if newX >= 0 && newX <= sizeGrid && newY >= 0 && newY <= sizeGrid &&
			!posInSlice(Pos{newX, newY}, corrupted) {
			successors = append(successors, Node{0, 0, 0, Pos{newX, newY}, &node})
		}
	}
	return successors
}

func inMapWithLowerF(successor Node, maps map[Pos]Node) bool {
	for _, val := range maps {
		if val.x == successor.x && val.y == successor.y && val.f < successor.f {
			return true
		}
	}
	return false
}

func reconstructPath(node Node) (int, []Pos) {
	var path []Pos
	for current := &node; current != nil; current = current.parent {
		path = append([]Pos{current.Pos}, path...)
	}
	return len(path) - 1, path
}

func aStar(corrupted []Pos, sizeGrid int) (int, []Pos) {
	var open = make(map[Pos]Node)
	var close = make(map[Pos]Node)
	var start = Node{0, 0, 0, Pos{0, 0}, nil}
	var goal = Pos{sizeGrid, sizeGrid}

	open[Pos{0, 0}] = start

	for len(open) > 0 {
		var node = findLeastF(open)
		delete(open, node.Pos)

		if node.Pos == goal {
			return reconstructPath(node)
		}

		var successors = findSuccessors(node, sizeGrid, corrupted)
		for _, successor := range successors {
			successor.g = node.g + 1
			successor.h = manhattanDistance(successor.Pos, goal)
			successor.f = successor.g + successor.h

			if _, exists := close[successor.Pos]; exists {
				continue
			}

			if inMapWithLowerF(successor, open) {
				continue
			}

			open[successor.Pos] = successor
		}

		close[node.Pos] = node
	}

	return -1, []Pos{}
}

func posInSlice(position Pos, positions []Pos) bool {
	for _, pos := range positions {
		if pos.x == position.x && pos.y == position.y {
			return true
		}
	}
	return false
}

func solution1(input string, nb_bytes int, sizeGrid int) int {
	var corrupted = parseCorrupted(input)
	corrupted = corrupted[:nb_bytes]
	var result, _ = aStar(corrupted, sizeGrid)
	return result
}

func solution2(input string, nb_bytes int, sizeGrid int) Pos {
	var corrupted = parseCorrupted(input)
	var _, chemin = aStar(corrupted[:nb_bytes], sizeGrid)
	for i := nb_bytes; i < len(corrupted); i++ {
		var currentCorrupted = corrupted[:i+1]
		if posInSlice(corrupted[i], chemin) {
			var pathLength int
			if pathLength, chemin = aStar(currentCorrupted, sizeGrid); pathLength == -1 {
				return Pos{corrupted[i].y, corrupted[i].x}
			}
		}
	}
	return Pos{-1, -1}
}

func main() {
	start := time.Now()
	fmt.Println("Solution 1:", solution1(inputTest, 12, 6))
	t1 := time.Now()
	fmt.Println("Solution 2:", solution2(inputDay, 1024, 70))
	end := time.Now()
	fmt.Println("Time solution 1 :", t1.Sub(start), "\nTime solution 2 :", end.Sub(t1))
}
