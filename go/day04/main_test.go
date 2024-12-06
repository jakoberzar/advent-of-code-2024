package main

import (
	"testing"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

func TestStar1Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star1(input), 18, t)
}

func TestStar1Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	utils.ExpectEqInt(star1(input), 2297, t)
}
func TestStar2Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star2(input), 9, t)
}
func TestStar2Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	utils.ExpectEqInt(star2(input), 1745, t)
}
