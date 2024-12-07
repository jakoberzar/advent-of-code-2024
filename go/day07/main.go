package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

const SIMPLE_INPUT = "./../../inputs/day-07/simple.txt"
const FULL_INPUT = "./../../inputs/day-07/full.txt"

type Equation struct {
	result int
	values []int
}

func main() {
	input := readInput(FULL_INPUT)
	fmt.Printf("Result of star 1 is %d\n", star1(input))
	// fmt.Printf("Result of star 2 is %d\n", star2(input))
}

func star1(input []Equation) (sum int) {
	for _, equation := range input {
		if canCompute(equation.values, equation.result, false) {
			sum += equation.result
		}
	}
	return
}

func canCompute(numbers []int, result int, allow_concat bool) bool {
	if len(numbers) < 1 {
		log.Fatal("Invalid call to compute with no numbers")
		return false // Shouldn't be reachable
	} else if len(numbers) == 1 {
		return numbers[0] == result
	}

	last := numbers[len(numbers)-1]
	headSlice := numbers[:len(numbers)-1]
	if canCompute(headSlice, result-last, allow_concat) {
		return true
	}
	if result%last == 0 && canCompute(headSlice, result/last, allow_concat) {
		return true
	}
	if allow_concat {
		largerBase10 := findMinLargeBase10(last)
		if result%largerBase10 == last {
			return canCompute(headSlice, result/largerBase10, allow_concat)
		}

	}
	return false
}

func findMinLargeBase10(n int) (res int) {
	res = 10
	for res <= n {
		res *= 10
	}
	return
}

func star2(input []Equation) (sum int) {
	for _, equation := range input {
		if canCompute(equation.values, equation.result, true) {
			sum += equation.result
		}
	}
	return
}

func readInput(path string) []Equation {
	lines := utils.ReadLinesOrDie(path)
	equations := make([]Equation, 0, len(lines))
	for _, line := range lines {
		split := strings.Split(line, ":")
		result, _ := strconv.Atoi(split[0])
		numStrings := strings.Fields(strings.TrimSpace(line))
		numbers := make([]int, 0, len(numStrings))
		for _, numStr := range numStrings {
			num, _ := strconv.Atoi(numStr)
			numbers = append(numbers, num)
		}
		equations = append(equations, Equation{result, numbers})
	}
	return equations
}
