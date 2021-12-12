package day11

import (
	"aoc2021/util"
)

type Coord struct {
	I, J int
}

func Part1() {
	lines := util.GetLines("input")
	grid := [][]int{}
	for i, line := range lines {
		grid = append(grid, []int{})
		for _, e := range line {
			grid[i] = append(grid[i], int(e-'0'))
		}
	}

	step := 0
	for ; !isInSync(grid); step++ {
		advance(grid)
	}

	println(step)
}

func isInSync(grid [][]int) bool {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] <= 9 {
				return false
			}
		}
	}
	return true
}

func advance(grid [][]int) (flash int) {
	q := []Coord{}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 9 {
				q = append(q, Coord{i, j})
			}

			if grid[i][j] > 9 {
				grid[i][j] = 0
			}

			grid[i][j] += 1
		}
	}

	for len(q) > 0 {
		first := q[0]
		flash += 1
		q = q[1:]
		nbs := getNbs(grid, first)
		for _, nb := range nbs {
			if grid[nb.I][nb.J] == 9 {
				q = append(q, nb)
			}

			grid[nb.I][nb.J] += 1
		}
	}

	return flash
}

func getNbs(grid [][]int, pt Coord) []Coord {
	nbs := []Coord{}
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			newI := i + pt.I
			newJ := j + pt.J
			newPt := Coord{newI, newJ}
			if newI >= 0 && newI < len(grid) &&
				newJ >= 0 && newJ < len(grid[0]) &&
				newPt != pt {
				nbs = append(nbs, Coord{pt.I + i, pt.J + j})
			}
		}
	}

	return nbs
}
