package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
)

func main() {
	file, err := os.Open("./day-01/input.txt")
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}
	defer file.Close()

	// Create a new scanner for the file
	scanner := bufio.NewScanner(file)

	list1 := make([]int, 0)
	list2 := make([]int, 0)
	// Read each line until EOF
	for scanner.Scan() {
		line := scanner.Text()
		regex := regexp.MustCompile(`\s+`)
		matches := regex.Split(line, -1)
		s1, _ := strconv.Atoi(matches[0])
		s2, _ := strconv.Atoi(matches[1])
		list1 = append(list1, s1)
		list2 = append(list2, s2)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error occurred during scanning: %v", err)
	}

	// sort lists
	slices.Sort(list1)
	slices.Sort(list2)

	sumOfDifferences := 0.
	for i, value2 := range list2 {
		value1 := list1[i]
		diff := value2 - value1
		sumOfDifferences += math.Abs(float64(diff))
	}

	fmt.Printf("Part 1: %v\n", int(sumOfDifferences))
	// part 2
	similarityScore := 0
	occurenceMap := getOccurenceMap(list2)
	for _, value := range list1 {
		if amountOfOccurences, exists := occurenceMap[value]; exists {
			similarityScore += amountOfOccurences * value
		}
	}
	fmt.Printf("Part 2: %v\n", similarityScore)
}

func getOccurenceMap(values []int) map[int]int {
	occurenceMap := make(map[int]int)
	for _, value := range values {
		occ, exists := occurenceMap[value]
		if exists {
			occurenceMap[value] = occ + 1
		} else {
			occurenceMap[value] = 1
		}
	}
	return occurenceMap
}
