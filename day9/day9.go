package day9

import (
	"aoc2021/util"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseLine(line string) []int {
	arr := []int{}
	for _, num := range strings.Split(line, "") {
		val, _ := strconv.Atoi(num)
		arr = append(arr, val)
	}

	return arr
}

func Part1() {
	lines := util.GetLines("day9/input")
	n := len(lines)
	grid := [][]int{}
	for i := 0; i < n; i++ {
		grid = append(grid, parseLine(lines[i]))
	}

	sizes := []int{}
	for index, low := range findLows(grid) {
		sizes = append(sizes, findBasin(grid, low, -index-1))
	}
	sort.Ints(sizes)

	product := 1
	for i := 0; i < 3; i++ {
		product *= sizes[len(sizes)-1-i]
	}

	fmt.Println(product)
}

func findLows(grid [][]int) []Point {
	n := len(grid)
	m := len(grid[0])
	res := []Point{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			val := grid[i][j]
			up := getNb(grid, i-1, j)
			down := getNb(grid, i+1, j)
			left := getNb(grid, i, j-1)
			right := getNb(grid, i, j+1)

			if val < up && val < down && val < left && val < right {
				res = append(res, Point{X: i, Y: j})
			}
		}
	}

	return res
}

func getNb(grid [][]int, i, j int) int {
	if i >= 0 && i < len(grid) && j >= 0 && j < len(grid[0]) {
		return grid[i][j]
	}

	return math.MaxInt
}

type Point struct {
	X, Y int
}

var dirXs = []int{0, 0, 1, -1}
var dirYs = []int{1, -1, 0, 0}

func findBasin(grid [][]int, start Point, basinIndex int) int {
	toProcess := []Point{start}
	size := 0

	for len(toProcess) > 0 {
		cur := toProcess[0]
		toProcess = toProcess[1:]

		val := grid[cur.X][cur.Y]
		if val < 0 {
			continue
		}

		nbs := []Point{}
		isCurBasin := true
		for i := 0; i < 4; i++ {
			nbX := cur.X + dirXs[i]
			nbY := cur.Y + dirYs[i]
			nbVal := getNb(grid, nbX, nbY)

			switch {
			case nbVal == basinIndex:
			case nbVal == math.MaxInt:
			case nbVal < 0:
			case nbVal == 9:
			case nbVal < val:
				isCurBasin = false
			default:
				nbs = append(nbs, Point{X: nbX, Y: nbY})
			}
		}

		if !isCurBasin {
			continue
		} else {
			size += 1
			grid[cur.X][cur.Y] = basinIndex
			toProcess = append(toProcess, nbs...)
		}
	}

	return size
}
