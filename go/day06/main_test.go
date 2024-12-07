package main

import (
	"testing"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

func TestStar1Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star1(input), 41, t)
}

func TestStar1Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	utils.ExpectEqInt(star1(input), 4602, t)
}

func TestStar2Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star2(input), 6, t)
}

func TestStar2Simple2(t *testing.T) {
	input := readInput(SIMPLE2_INPUT)
	utils.ExpectEqInt(star2(input), 0, t)
}

func TestStar2Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	utils.ExpectEqInt(star2(input), 1703, t)
}
