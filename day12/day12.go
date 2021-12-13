package day12

import (
	"aoc2021/util"
	"strings"
)

func visit(nbs map[string][]string, visited map[string]bool, next string, count *int, usedSingle bool) {
	if next == "end" {
		*count += 1
		return
	}

	for _, nb := range nbs[next] {
		canV, usingSingle := canVisit(nb, visited, usedSingle)
		if canV {
			preVisited := visited[nb]

			visited[nb] = true
			visit(nbs, visited, nb, count, usingSingle)
			visited[nb] = preVisited
		}
	}
}

func canVisit(name string, visited map[string]bool, usedSingle bool) (res bool, usingSingle bool) {
	if isLarge(name) {
		return true, usedSingle
	}

	if name == "start" {
		return false, usedSingle
	}

	if !visited[name] {
		return true, usedSingle
	}

	if usedSingle {
		return false, usedSingle
	}

	return true, true
}

func isLarge(name string) bool {
	return strings.ToUpper(name) == name
}

func Part1() {
	count := 0
	graph := make(map[string][]string)

	for _, line := range util.GetLines("input") {
		parseLine(line, graph)
	}

	visited := map[string]bool{
		"start": true,
	}
	visit(graph, visited, "start", &count, false)
	println(count)
}

func parseLine(line string, graph map[string][]string) {
	parts := strings.Split(line, "-")
	p1 := parts[0]
	p2 := parts[1]
	graph[p1] = append(graph[p1], p2)
	graph[p2] = append(graph[p2], p1)
}
