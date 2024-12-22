package day21

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"math"
	"strconv"
	"strings"
)

type Day21 struct {
	data []string
}

func newDay21(data []string) *Day21 {
	return &Day21{data: data}
}

var layouts = map[layoutType]layout{
	numeric: {
		'7': *graph.NewVector2(0, 0), '8': *graph.NewVector2(1, 0), '9': *graph.NewVector2(2, 0),
		'4': *graph.NewVector2(0, 1), '5': *graph.NewVector2(1, 1), '6': *graph.NewVector2(2, 1),
		'1': *graph.NewVector2(0, 2), '2': *graph.NewVector2(1, 2), '3': *graph.NewVector2(2, 2),
		'0': *graph.NewVector2(1, 3), 'A': *graph.NewVector2(2, 3),
	},
	directional: {
		'^': *graph.NewVector2(1, 0), 'A': *graph.NewVector2(2, 0),
		'<': *graph.NewVector2(0, 1), 'v': *graph.NewVector2(1, 1), '>': *graph.NewVector2(2, 1),
	},
}

var allPaths = map[layoutType]map[rune]map[rune][]string{
	numeric:     allShortestPaths(numeric),
	directional: allShortestPaths(directional),
}

func allShortestPaths(lt layoutType) map[rune]map[rune][]string {
	pathsFrom := make(map[rune]map[rune][]string)
	l := layouts[lt]
	for start := range l {
		pathsFrom[start] = make(map[rune][]string)
		for end := range l {
			paths := l.shortestPaths(start, end)
			pathsFrom[start][end] = paths
		}
	}
	return pathsFrom
}

type layoutType int

const (
	numeric layoutType = iota
	directional
)

type pathRecord struct {
	position  graph.Vector2
	direction direction.Direction
}

func (l *layout) shortestPaths(start rune, end rune) []string {
	if start == end {
		return []string{""}
	}

	positions := make(map[graph.Vector2]bool)
	for _, v := range *l {
		positions[v] = true
	}
	from := (*l)[start]
	to := (*l)[end]

	getNext := func(p graph.Vector2) []pathRecord {
		var paths []pathRecord
		for _, d := range direction.Cardinals {
			next := *p.Add(d.Delta())
			if !positions[next] || to.ManhattanDistance(next) > to.ManhattanDistance(p) {
				continue
			}
			paths = append(paths, pathRecord{position: next, direction: d})
		}
		return paths
	}
	var queue [][]pathRecord
	for _, p := range getNext(from) {
		queue = append(queue, []pathRecord{p})
	}
	for len(queue) > 0 {
		currentLevel := queue
		for i := 0; i < len(currentLevel); i++ {
			currentPath := currentLevel[i]
			last := currentPath[len(currentPath)-1]
			if last.position == to {
				var shortestPaths []string
				validPaths := validPaths(currentLevel, to)
				for _, path := range validPaths {
					shortestPaths = append(shortestPaths, asString(path))
				}
				return shortestPaths
			} else {
				for _, next := range getNext(last.position) {
					path := append(currentPath, next)
					queue = append(queue, path)
				}
			}
		}
		queue = queue[len(currentLevel):]
	}

	return nil
}

func asString(path []pathRecord) string {
	s := ""
	for _, p := range path {
		switch p.direction {
		case direction.Up:
			s += "^"
		case direction.Right:
			s += ">"
		case direction.Down:
			s += "v"
		case direction.Left:
			s += "<"
		}
	}
	return s
}

func validPaths(allPaths [][]pathRecord, target graph.Vector2) [][]pathRecord {
	var paths [][]pathRecord
	for _, path := range allPaths {
		if path[len(path)-1].position == target {
			paths = append(paths, path)
		}
	}
	return paths
}

type layout map[rune]graph.Vector2

func (l *layout) path(start rune, end rune) string {
	dx := (*l)[end].X - (*l)[start].X
	dy := (*l)[end].Y - (*l)[start].Y

	var s string
	if dx != 0 {
		r := '>'
		if dx < 0 {
			r = '<'
		}
		s = strings.Repeat(string(r), abs(dx))
	}
	if dy != 0 {
		r := 'v'
		if dy < 0 {
			r = '^'
		}
		s += strings.Repeat(string(r), abs(dy))
	}
	return s
}

func press(code string, layoutTypes []layoutType, depth int, cache map[int]map[string]int) int {
	seenValue, ok := cache[depth][code]
	if ok {
		return seenValue
	}
	s := "A" + code
	total := 0
	if depth == len(layoutTypes)-1 {
		return len(s) - 1
	}
	for i := 1; i < len(s); i++ {
		paths := allPaths[layoutTypes[depth]][rune(s[i-1])][rune(s[i])]
		minPath := math.MaxInt
		for _, path := range paths {
			minPath = min(minPath, press(path+"A", layoutTypes, depth+1, cache))
		}
		total += minPath
	}
	cache[depth][code] = total
	return total
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func (d *Day21) pressAll(robots int) int {
	layoutTypes := []layoutType{numeric}
	for i := 0; i < robots; i++ {
		layoutTypes = append(layoutTypes, directional)
	}

	cache := make(map[int]map[string]int)
	for i := 0; i < len(layoutTypes); i++ {
		cache[i] = make(map[string]int)
	}

	total := 0
	for _, code := range d.data {
		presses := press(code, layoutTypes, 0, cache)
		num, _ := strconv.Atoi(code[:3])
		//fmt.Printf("[%s]= %d * %d\n", code, presses, num)
		total += num * presses
	}
	return total
}

func (d *Day21) part1() int {
	return d.pressAll(3)
}

func (d *Day21) part2() int {
	return d.pressAll(26)
}
