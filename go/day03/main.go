package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

const SIMPLE_INPUT = "./../../inputs/day-03/simple.txt"
const SIMPLE2_INPUT = "./../../inputs/day-03/simple2.txt"
const FULL_INPUT = "./../../inputs/day-03/full.txt"

func main() {
	s := readInput(FULL_INPUT)
	fmt.Printf("Result of star 1 is %d\n", star1(s))
	fmt.Printf("Result of star 2 is %d\n", star2(s))
}

func star1(input string) (sum int) {
	r := regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)
	for _, line := range r.FindAllStringSubmatch(input, -1) {
		first, _ := strconv.Atoi(line[1])
		second, _ := strconv.Atoi(line[2])
		sum += first * second
	}
	return
}

func star2(input string) (sum int) {
	r := regexp.MustCompile(`(?:mul\(([0-9]+),([0-9]+)\))|(do\(\))|(don't\(\))`)
	enabled := true
	for _, line := range r.FindAllStringSubmatch(input, -1) {
		switch line[0] {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				first, _ := strconv.Atoi(line[1])
				second, _ := strconv.Atoi(line[2])
				sum += first * second
			}
		}
	}
	return
}

func readInput(path string) string {
	return utils.ReadFileOrDie(path)
}
