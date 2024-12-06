package main

import (
	"fmt"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

const SIMPLE_INPUT = "./../../inputs/day-04/simple.txt"
const FULL_INPUT = "./../../inputs/day-04/full.txt"

func main() {
	lines := readInput(FULL_INPUT)
	fmt.Printf("Result of star 1 is %d\n", star1(lines))
	fmt.Printf("Result of star 2 is %d\n", star2(lines))
}

func star1(input []string) (count int) {
	maxY := len(input)
	maxX := len(input[0])
	len := 4
	var directions = []utils.Direction{
		utils.Up{MaxY: maxY, MaxX: maxX, Len: len},
		utils.Left{MaxY: maxY, MaxX: maxX, Len: len},
		utils.Down{MaxY: maxY, MaxX: maxX, Len: len},
		utils.Right{MaxY: maxY, MaxX: maxX, Len: len},
		utils.LeftUp{MaxY: maxY, MaxX: maxX, Len: len},
		utils.RightUp{MaxY: maxY, MaxX: maxX, Len: len},
		utils.LeftDown{MaxY: maxY, MaxX: maxX, Len: len},
		utils.RightDown{MaxY: maxY, MaxX: maxX, Len: len},
	}
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if input[y][x] == 'X' {
				for _, direction := range directions {
					if !direction.Eligible(x, y) {
						continue
					}
					mX, mY := direction.Next(x, y)
					if input[mY][mX] != 'M' {
						continue
					}
					aX, aY := direction.Next(mX, mY)
					if input[aY][aX] != 'A' {
						continue
					}
					sX, sY := direction.Next(aX, aY)
					if input[sY][sX] != 'S' {
						continue
					}
					count += 1
				}
			}
		}
	}
	return
}

func star2(input []string) (count int) {
	maxY := len(input)
	maxX := len(input[0])
	len := 2
	leftUp := utils.LeftUp{MaxY: maxY, MaxX: maxX, Len: len}
	rightUp := utils.RightUp{MaxY: maxY, MaxX: maxX, Len: len}
	leftDown := utils.LeftDown{MaxY: maxY, MaxX: maxX, Len: len}
	rightDown := utils.RightDown{MaxY: maxY, MaxX: maxX, Len: len}
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if input[y][x] == 'A' {
				if !leftUp.Eligible(x, y) || !rightDown.Eligible(x, y) {
					continue
				}
				luX, luY := leftUp.Next(x, y)
				luChar := input[luY][luX]
				ruX, ruY := rightUp.Next(x, y)
				ruChar := input[ruY][ruX]
				ldX, ldY := leftDown.Next(x, y)
				ldChar := input[ldY][ldX]
				rdX, rdY := rightDown.Next(x, y)
				rdChar := input[rdY][rdX]
				if !(luChar == 'M' && rdChar == 'S' || luChar == 'S' && rdChar == 'M') {
					continue
				}
				if !(ruChar == 'M' && ldChar == 'S' || ruChar == 'S' && ldChar == 'M') {
					continue
				}
				count += 1
			}
		}
	}
	return
}

func readInput(path string) (lines []string) {
	return utils.ReadLinesOrDie(path)
}
