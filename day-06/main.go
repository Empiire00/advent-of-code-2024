package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type (
	coordinate struct {
		x, y int
	}
	obstacleMap  map[coordinate]bool
	directionMap map[direction]bool
	visitMap     map[coordinate]directionMap
	direction    int
	gridSize     struct {
		x, y int
	}
)

const (
	up direction = iota
	right
	down
	left
)

type guard struct {
	direction direction
	position  coordinate
}

func (v directionMap) Character() string {
	hasVertical := false
	hasHorizontal := false
	if _, hasUp := v[up]; hasUp {
		hasVertical = true
	} else if _, hasDown := v[down]; hasDown {
		hasVertical = true
	}
	if _, hasRight := v[right]; hasRight {
		hasHorizontal = true
	} else if _, hasLeft := v[left]; hasLeft {
		hasHorizontal = true
	}

	switch {
	case hasHorizontal && hasVertical:
		return "+"
	case hasHorizontal:
		return "-"
	case hasVertical:
		return "|"
	default:
		panic("cannot call Character() on empty map")
	}
}

func (g *guard) turnRight() {
	g.direction = (g.direction + 1) % 4
	// fmt.Printf("Turning Right. New direction: %v\n", g.direction)
}

func (g *guard) getFacingCoordinate() coordinate {
	x := g.position.x
	y := g.position.y
	switch g.direction {
	case right:
		x++
	case left:
		x--
	case up:
		y--
	case down:
		y++
	}
	return coordinate{x: x, y: y}
}

func (g *guard) moveForward() {
	// fmt.Printf("Moving forward to: %v\n", g.getFacingCoordinate())
	g.position = g.getFacingCoordinate()
}

func (g *guard) moveTo(position coordinate) {
	// fmt.Printf("Moving to: %v\n", position)
}

func main() {
	guardian, obstacles, gridSize, err := parseFile("./day-06/input.txt")
	if err != nil {
		log.Fatalf("Failed to parse input.txt: %v", err)
		return
	}
	visited, _ := getVisited(gridSize, obstacles, guardian)
	amountOfVisited := len(visited)
	printGrid(gridSize, visited, obstacles, guardian)
	fmt.Printf("Part 1: %d\n", amountOfVisited)

	// part2
	positions := make(map[coordinate]int, 0)
	// try every visited position as an obstacle spot
	for position := range visited {
		// do not put it at the starting position
		if position == guardian.position {
			continue
		}
		// copy obstacles in new map
		newObstacles := make(obstacleMap)
		for i, element := range obstacles {
			newObstacles[i] = element
		}
		newObstacles[position] = true

		_, didExit := getVisited(gridSize, newObstacles, guardian)
		if !didExit {
			positions[position]++
		}
	}
	fmt.Printf("Part 2: %v\n", len(positions))
}

// function returns visited spots and whether the guard left the area
func getVisited(gridSize gridSize, obstacles obstacleMap, guardian guard) (visitMap, bool) {
	visited := make(visitMap, 0)
	leftArea := false
	for {
		// printGrid(gridSize, visited, obstacleMap, guard)
		// create map at position, if it does not exist already
		if _, exists := visited[guardian.position]; !exists {
			visited[guardian.position] = make(map[direction]bool)
		}
		// check, whether guard has been here yet -> cycle
		if _, exists := visited[guardian.position][guardian.direction]; exists {
			break
		}
		visited[guardian.position][guardian.direction] = true
		facingPosition := guardian.getFacingCoordinate()
		if !(facingPosition.x < gridSize.x && facingPosition.x >= 0 && facingPosition.y < gridSize.y && facingPosition.y >= 0) {
			leftArea = true
			break
		}
		if _, included := obstacles[facingPosition]; included {
			guardian.turnRight()
		} else {
			guardian.moveForward()
		}
	}
	return visited, leftArea
}

func parseFile(path string) (guard, obstacleMap, gridSize, error) {
	file, err := os.Open(path)
	if err != nil {
		return guard{}, nil, gridSize{}, err
	}
	defer file.Close()

	var guardian guard
	obstacle := make(obstacleMap, 0)
	y := 0
	maxX := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		maxX = len(line)
		for pos, char := range line {
			// pos is the x coordinate
			switch char {
			case '.':
				continue
			case '#':
				obstacle[coordinate{x: pos, y: y}] = true
			case '^':
				guardian.position = coordinate{pos, y}
			}
		}
		y += 1
	}
	return guardian, obstacle, gridSize{x: maxX, y: y}, nil
}

func printGrid(size gridSize, visited visitMap, obstacles obstacleMap, guard guard) {
	for y := 0; y < size.y; y++ {
		fmt.Print("")
		for x := 0; x < size.x; x++ {
			position := coordinate{x: x, y: y}
			if position == guard.position {
				fmt.Print("O")
				continue
			}
			if value, ok := visited[position]; ok {
				fmt.Print(value.Character())
				continue
			}
			if _, ok := obstacles[position]; ok {
				fmt.Print("#")
				continue
			}

			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Println()
}
