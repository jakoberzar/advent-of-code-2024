package main

import (
	"testing"
)

func TestStar1Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	want := 18
	got := star1(input)
	if got != want {
		t.Errorf("Simple result of star 1 is %d instead of %d", got, want)
	}
}

func TestStar1Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	want := 2297
	got := star1(input)
	if got != want {
		t.Errorf("Full result of star 1 is %d instead of %d", got, want)
	}
}
func TestStar2Simple(t *testing.T) {
	input := readInput(SIMPLE_INPUT)
	want := 9
	got := star2(input)
	if got != want {
		t.Errorf("Simple result of star 2 is %d instead of %d", got, want)
	}
}
func TestStar2Full(t *testing.T) {
	input := readInput(FULL_INPUT)
	want := 1745
	got := star2(input)
	if got != want {
		t.Errorf("Full result of star 2 is %d instead of %d", got, want)
	}
}
