package main

import (
	"fmt"

	"github.com/jakoberzar/advent-of-code-2024/internal/grids"
	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

const SIMPLE_INPUT = "./../../inputs/day-08/simple.txt"
const FULL_INPUT = "./../../inputs/day-08/full.txt"

type Loc struct {
	x, y int
}

func (l Loc) Diff(other Loc) (diffX, diffY int) {
	diffX = other.x - l.x
	diffY = other.y - l.y
	return
}

type Group struct {
	name      rune
	locations []Loc
}

func main() {
	input := readInput(FULL_INPUT)
	fmt.Printf("Result of star 1 is %d\n", star1(input))
	fmt.Printf("Result of star 2 is %d\n", star2(input))
}

func star1(input []string) (count int) {
	groups := getAntennaGroups(input)
	height := len(input)
	width := len(input[0])
	antinodeGrid := grids.CreateBoolGrid(height, width)
	for _, group := range groups {
		for i, loc1 := range group.locations {
			for _, loc2 := range group.locations[i+1:] {
				antinode1 := findAntinode(loc1, loc2)
				antinode2 := findAntinode(loc2, loc1)
				if antinodeGrid.IsInsideBounds(antinode1.x, antinode1.y) {
					antinodeGrid.Set(antinode1.x, antinode1.y)
				}
				if antinodeGrid.IsInsideBounds(antinode2.x, antinode2.y) {
					antinodeGrid.Set(antinode2.x, antinode2.y)
				}
			}
		}
	}
	return antinodeGrid.CountTrue()
}

func findAntinode(loc1, loc2 Loc) Loc {
	diffX, diffY := loc1.Diff(loc2)
	return getAntinode(loc2, 1, diffX, diffY)
}

func getAntinode(loc Loc, ith, diffX, diffY int) Loc {
	return Loc{x: loc.x + ith*diffX, y: loc.y + ith*diffY}
}

func setAntinodes(loc1, loc2 Loc, grid *grids.BoolGrid) {
	diffX, diffY := loc1.Diff(loc2)
	ith := 1
	antinode := getAntinode(loc2, ith, diffX, diffY)
	for grid.IsInsideBounds(antinode.x, antinode.y) {
		grid.Set(antinode.x, antinode.y)
		ith += 1
		antinode = getAntinode(loc2, ith, diffX, diffY)
	}
}

func star2(input []string) (count int) {
	groups := getAntennaGroups(input)
	height := len(input)
	width := len(input[0])
	antinodeGrid := grids.CreateBoolGrid(height, width)
	for _, group := range groups {
		for i, loc1 := range group.locations {
			for _, loc2 := range group.locations[i+1:] {
				// Add antennas themselves
				antinodeGrid.Set(loc1.x, loc1.y)
				antinodeGrid.Set(loc2.x, loc2.y)

				setAntinodes(loc1, loc2, &antinodeGrid)
				setAntinodes(loc2, loc1, &antinodeGrid)

			}
		}
	}
	return antinodeGrid.CountTrue()
}

func readInput(path string) []string {
	return utils.ReadLinesOrDie(path)
}

func getAntennaGroups(input []string) map[rune]Group {
	groups := make(map[rune]Group)
	for y, line := range input {
		for x, ch := range line {
			if ch != '.' {
				group, in := groups[ch]
				if !in {
					group = Group{name: ch}
				}
				group.locations = append(group.locations, Loc{x: x, y: y})
				groups[ch] = group
			}
		}
	}
	return groups
}
