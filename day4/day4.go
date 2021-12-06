package day4

import (
	"aoc2021/util"
	"fmt"
	"strconv"
	"strings"
)

type Pos struct {
	Row int
	Col int
}

type Board struct {
	Map        map[int]Pos
	RowRemains []int
	ColRemains []int
	Won        bool
}

func (b *Board) Mark(val int) (score int, won bool) {
	if b.Won {
		return 0, false
	}

	pos, ok := b.Map[val]
	if !ok {
		return 0, false
	}

	delete(b.Map, val)
	b.Map[-val-1] = pos

	b.RowRemains[pos.Row] -= 1
	if b.RowRemains[pos.Row] == 0 {
		b.Won = true
		score := b.GetScore(val)
		return score, true
	}

	b.ColRemains[pos.Col] -= 1
	if b.ColRemains[pos.Col] == 0 {
		b.Won = true
		score := b.GetScore(val)
		return score, true
	}

	return 0, false
}

func (b *Board) GetScore(lastHit int) int {
	sum := 0
	for k := range b.Map {
		if k >= 0 {
			sum += k
		}
	}

	return sum * lastHit
}

func parseBoard(lines []string) *Board {
	board := &Board{
		RowRemains: []int{5, 5, 5, 5, 5},
		ColRemains: []int{5, 5, 5, 5, 5},
		Map:        make(map[int]Pos),
		Won:        false,
	}

	for row, line := range lines {
		for col, val := range strings.Fields(line) {
			intVal, _ := strconv.Atoi(val)
			board.Map[intVal] = Pos{Row: row, Col: col}
		}
	}

	return board
}

func Part1() {
	lines := util.GetLines("day4/input")
	nums := parseLine1(lines[0])
	boards := []*Board{}

	for i := 2; i < len(lines); i += 6 {
		b := parseBoard(lines[i : i+5])
		boards = append(boards, b)
	}

	bCount := len(boards)
	for _, num := range nums {
		for _, b := range boards {
			score, won := b.Mark(num)
			if won {
				bCount -= 1
				if bCount == 0 {
					fmt.Println(score)
					return
				}
			}
		}
	}

	fmt.Println("Nobody won!")
}

func parseLine1(line string) []int {
	res := []int{}
	for _, str := range strings.Split(line, ",") {
		intVal, _ := strconv.Atoi(str)
		res = append(res, intVal)
	}

	return res
}
