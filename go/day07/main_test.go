package main

import (
	"testing"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

func TestStar1Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star1(input), 3749, t)
}

func TestStar1Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	utils.ExpectEqInt(star1(input), 1611660863222, t)
}

func TestStar2Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star2(input), 11387, t)
}

func TestStar2Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	utils.ExpectEqInt(star2(input), 62098619, t)
}
