package main

import (
	"testing"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

func TestStar1Simple(t *testing.T) {
	rules, updates := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star1(rules, updates), 143, t)
}

func TestStar1Full(t *testing.T) {
	rules, updates := readInput(FULL_INPUT)
	utils.ExpectEqInt(star1(rules, updates), 6041, t)
}

func TestStar2Simple(t *testing.T) {
	rules, updates := readInput(SIMPLE_INPUT)
	utils.ExpectEqInt(star2(rules, updates), 123, t)
}

func TestStar2Full(t *testing.T) {
	rules, updates := readInput(FULL_INPUT)
	utils.ExpectEqInt(star2(rules, updates), 4884, t)
}
