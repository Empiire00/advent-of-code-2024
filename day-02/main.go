package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	levelsList, err := parseFile("./day-02/input.txt")
	if err != nil {
		log.Fatalf("Failed to parse input.txt: %v", err)
		return
	}
	// fmt.Printf("List of levels: %v\n", levelsList)
	sumOfSafe := 0
	for _, levels := range levelsList {
		if isLevelSafe(levels) {
			sumOfSafe++
		}
	}
	fmt.Printf("Part 1: %v\n", sumOfSafe)
	// part2

	sumOfDampenedSafe := 0
	for _, levels := range levelsList {
		for _, dampenedLevels := range getDampenedLevels(levels) {
			if isLevelSafe(dampenedLevels) {
				sumOfDampenedSafe++
				break
			}
		}
	}
	fmt.Printf("Part 2: %v\n", sumOfDampenedSafe)
}

func isLevelSafe(levels []int) bool {
	// if all values are either ascending or descending, the list is sorted (and possibly also reversed)
	isSorted := true
	sortedDescList := slices.Sorted(slices.Values(levels))
	slices.Reverse(sortedDescList)
	isSorted = slices.Equal(sortedDescList, levels) || slices.IsSorted(levels)
	if !isSorted {
		return false
	}
	for index := 1; index < len(levels); index++ {
		difference := math.Abs((float64(levels[index-1] - levels[index])))
		if difference > 3 || difference == 0 {
			return false
		}
	}
	return true
}

func getDampenedLevels(levels []int) [][]int {
	dampenedLevels := make([][]int, 0)
	// fmt.Printf("Levels: %v\n", levels)
	for index := 0; index < len(levels); index++ {
		// fmt.Println(index)
		dampenedLevel := removeFromSlice(levels, index)
		// fmt.Printf("%v, %v\n", levels[:index], levels[index+1:])
		dampenedLevels = append(dampenedLevels, dampenedLevel)
	}
	// fmt.Printf("Dampened: %v\n", dampenedLevels)
	return dampenedLevels
}

func removeFromSlice(slice []int, index int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	return append(sliceCopy[:index], sliceCopy[index+1:]...)
}

func parseFile(path string) ([][]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := make([]int, 0)
		for _, char := range strings.Split(line, " ") {
			num, err := strconv.Atoi(string(char))
			if err != nil {
				return nil, err
			}
			levels = append(levels, num)
		}
		result = append(result, levels)
	}
	return result, nil
}
