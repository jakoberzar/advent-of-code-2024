package main

import (
	"fmt"

	"github.com/jakoberzar/advent-of-code-2024/internal/utils"
)

const SIMPLE_INPUT = "./../../inputs/day-06/simple.txt"
const SIMPLE2_INPUT = "./../../inputs/day-06/simple2.txt"
const FULL_INPUT = "./../../inputs/day-06/full.txt"

type Grid = [][]byte

func main() {
	s := readInput(FULL_INPUT)
	fmt.Printf("Result of star 1 is %d\n", star1(s))
	s = readInput(FULL_INPUT)
	fmt.Printf("Result of star 2 is %d\n", star2(s))
}

func star1(grid Grid) (count int) {
	height := len(grid)
	width := len(grid[0])
	walkGrid(grid)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			char := grid[y][x]
			// fmt.Printf("%c", char)
			if char == 'X' || char == '^' {
				count += 1
			}
		}
		// fmt.Println()
	}
	return
}

func findGuard(grid Grid) (x, y int) {
	for y = 0; y < len(grid); y++ {
		for x = 0; x < len(grid[y]); x++ {
			if grid[y][x] == '^' {
				return
			}
		}
	}
	return -1, -1
}

func getDirections(grid Grid) []utils.Direction {
	height := len(grid)
	width := len(grid[0])
	return []utils.Direction{
		utils.Up{
			MaxY: height,
			MaxX: width,
			Len:  2,
		},
		utils.Right{
			MaxY: height,
			MaxX: width,
			Len:  2,
		},
		utils.Down{
			MaxY: height,
			MaxX: width,
			Len:  2,
		},
		utils.Left{
			MaxY: height,
			MaxX: width,
			Len:  2,
		},
	}

}

func walkGrid(grid Grid) {
	directions := getDirections(grid)
	dirIdx := 0
	curX, curY := findGuard(grid)
	for directions[dirIdx].Eligible(curX, curY) {
		nextX, nextY := directions[dirIdx].Next(curX, curY)
		if grid[nextY][nextX] == '#' {
			dirIdx = (dirIdx + 1) % 4
			continue
		}
		grid[nextY][nextX] = '^'
		grid[curY][curX] = 'X'
		curX, curY = nextX, nextY
	}
}

type ObstacleFound struct {
	y, x, turn, direction int
}

func star2(grid Grid) (count int) {
	directions := getDirections(grid)
	dirIdx := 0
	curX, curY := findGuard(grid)
	turn := 0
	obstacles := make([]ObstacleFound, 0)
	// obstaclesAdded := make(map[int]map[int]bool)
	// for y := 0; y < len(grid); y++ {
	// 	obstaclesAdded[y] = make(map[int]bool)
	// }
	for directions[dirIdx].Eligible(curX, curY) {
		nextX, nextY := directions[dirIdx].Next(curX, curY)
		if grid[nextY][nextX] == '#' {
			obstacles = append(obstacles, ObstacleFound{y: nextY, x: nextX, turn: turn, direction: dirIdx})
			dirIdx = (dirIdx + 1) % 4
			turn += 1
			continue
		}

		if grid[nextY][nextX] == '.' {
			// Try putting an obstacle here
			if obstacleMakesALoop(curY, curX, nextY, nextX, turn, dirIdx, grid, obstacles, directions) {
				// obstaclesAdded[nextY][nextX] = true
				count += 1
			}
		}

		// Move
		grid[nextY][nextX] = '^'
		grid[curY][curX] = 'X'
		curX, curY = nextX, nextY
		turn += 1
	}
	return
}

func obstacleMakesALoop(curY, curX, obstacleY, obstacleX, turn, dirIdx int, grid Grid, obstacles []ObstacleFound, directions []utils.Direction) bool {
	// Hit the obstacle here
	obstacles = append(obstacles, ObstacleFound{y: obstacleY, x: obstacleX, turn: turn, direction: dirIdx})
	dirIdx = (dirIdx + 1) % 4
	turn += 1
	// Next turn
	for directions[dirIdx].Eligible(curX, curY) {
		nextX, nextY := directions[dirIdx].Next(curX, curY)
		if grid[nextY][nextX] == '#' || nextY == obstacleY && nextX == obstacleX {
			for _, obstacle := range obstacles {
				if obstacle.y == nextY && obstacle.x == nextX && obstacle.direction == dirIdx {
					// Already came this way before
					return true
				}
			}
			// This obstacle was not encountered from this direction previously
			obstacles = append(obstacles, ObstacleFound{y: nextY, x: nextX, turn: turn, direction: dirIdx})
			dirIdx = (dirIdx + 1) % 4
			turn += 1
			continue
		}
		curX, curY = nextX, nextY
		turn += 1
	}
	// Came out of the map
	return false
}

func readInput(path string) Grid {
	lines := utils.ReadLinesOrDie(path)
	grid := make(Grid, 0, len(lines))
	for _, line := range lines {
		grid = append(grid, []byte(line))
	}
	return grid
}
