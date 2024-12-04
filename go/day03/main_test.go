package main

import (
	"testing"
)

func TestStar1Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	want := 161
	got := star1(input)
	if got != want {
		t.Errorf("Simple result of star 1 is %d instead of %d", got, want)
	}
}

func TestStar1Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	want := 183788984
	got := star1(input)
	if got != want {
		t.Errorf("Full result of star 1 is %d instead of %d", got, want)
	}
}
func TestStar2Simple(t *testing.T) {
	input := readInput(SIMPLE2_INPUT)
	want := 48
	got := star2(input)
	if got != want {
		t.Errorf("Simple result of star 2 is %d instead of %d", got, want)
	}
}
func TestStar2Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	want := 62098619
	got := star2(input)
	if got != want {
		t.Errorf("Full result of star 2 is %d instead of %d", got, want)
	}
}
