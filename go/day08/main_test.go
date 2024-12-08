package main

import (
	"testing"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

func TestStar1Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star1(input), 14, t)
}

func TestStar1Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	utils.ExpectEqInt(star1(input), 332, t)
}

func TestStar2Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star2(input), 34, t)
}

func TestStar2Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	utils.ExpectEqInt(star2(input), 1174, t)
}
