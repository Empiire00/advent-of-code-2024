package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("./day-03/input.txt")
	if err != nil {
		log.Fatal("Failed to open input.txt: %v", err)
		return
	}

	matches, err := getMatches(string(input))
	if err != nil {
		log.Printf("No matches found: %v", err)
		return
	}

	multplicationResults, err := getMultiplicationResults(matches)
	if err != nil {
		log.Printf("Could not get multiplication results: %v", err)
		return
	}

	sum := reduce(multplicationResults, func(a, b int) int {
		return a + b
	}, 0)

	fmt.Printf("Part 1: %v\n", sum)

	// part 2
	relevantInstructions := make([]string, 0)
	currentlyRelevant := true
	matches2, err := getMatches2(string(input))
	if err != nil {
		log.Printf("Could not match (Part 2): %v", err)
	}
	for _, v := range matches2 {
		switch {
		case strings.Contains(v, "don't"):
			currentlyRelevant = false
		case strings.Contains(v, "do"):
			currentlyRelevant = true
		default:
			{
				if currentlyRelevant {
					relevantInstructions = append(relevantInstructions, v)
				}
			}
		}
	}
	multplicationResults2, err := getMultiplicationResults(relevantInstructions)
	if err != nil {
		log.Printf("Could not get multiplication results: %v", err)
		return
	}
	sum2 := reduce(multplicationResults2, func(a, b int) int {
		return a + b
	}, 0)

	fmt.Printf("Part 2: %v\n", sum2)
}

func getMatches(input string) ([]string, error) {
	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := regex.FindAllString(input, -1)

	if len(matches) > 0 {
		return matches, nil
	}
	return nil, errors.New("No matches found")
}

func getMatches2(input string) ([]string, error) {
	regex := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	matches := regex.FindAllString(input, -1)

	if len(matches) > 0 {
		return matches, nil
	}
	return nil, errors.New("No matches found")
}

func getMultiplicationResults(multiplicationStrings []string) ([]int, error) {
	result := make([]int, 0)

	for _, v := range multiplicationStrings {
		regex := regexp.MustCompile(`\((\d+),(\d+)\)`)
		matches := regex.FindStringSubmatch(v)

		if len(matches) != 3 {
			return nil, fmt.Errorf("Could not find matches in: %v", v)
		}
		s1, err1 := strconv.Atoi(matches[1])
		s2, err2 := strconv.Atoi(matches[2])
		if err1 != nil || err2 != nil {
			return nil, fmt.Errorf("Failed to parse: %v", v)
		}
		multiplicationResult := s1 * s2
		result = append(result, multiplicationResult)
	}
	return result, nil
}

func reduce[T, M any](values []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, value := range values {
		acc = f(acc, value)
	}
	return acc
}
