package day18

import (
	"adventofcode2024/util/direction"
	"adventofcode2024/util/graph"
	"adventofcode2024/util/ints"
	"fmt"
	"unicode"
)

type Day18 struct {
	bytes      []graph.Vector2
	bytesCount int
	exit       graph.Vector2
}

func newDay18(data []string) *Day18 {
	return newDay18WithSize(data, 1024, 70, 70)
}

func newDay18WithSize(data []string, bytesCount int, exitX int, exitY int) *Day18 {
	var bytes []graph.Vector2
	for _, row := range data {
		nums := ints.ToInts(row, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		bytes = append(bytes, *graph.NewVector2(nums[0], nums[1]))
	}
	return &Day18{bytes: bytes, bytesCount: bytesCount, exit: *graph.NewVector2(exitX, exitY)}
}

func (d *Day18) findPath(bytes int) int {
	walls := make(map[graph.Vector2]bool)
	for _, b := range d.bytes[:bytes] {
		walls[b] = true
	}

	inbounds := func(v graph.Vector2) bool {
		return v.X >= 0 && v.X <= d.exit.X && v.Y >= 0 && v.Y <= d.exit.Y
	}

	visited := make(map[graph.Vector2]bool)
	queue := []graph.Vector2{*graph.NewVector2(0, 0)}
	steps := 0
	for len(queue) > 0 {
		levelCount := len(queue)
		for i := 0; i < levelCount; i++ {
			current := queue[0]
			queue = queue[1:]
			if current == d.exit {
				return steps
			} else if visited[current] {
				continue
			}
			visited[current] = true

			for _, dir := range direction.Cardinals {
				next := *current.Add(dir.Delta())
				if inbounds(next) && !visited[next] && !walls[next] {
					queue = append(queue, next)
				}
			}
		}
		steps++
	}
	return -1
}

func (d *Day18) part1() int {
	return d.findPath(d.bytesCount)
}

func (d *Day18) part2() string {
	low := 0
	high := len(d.bytes)
	for low < high {
		mid := low + (high-low)/2
		//looking for l has valid path, r does not
		l := d.findPath(mid)
		r := d.findPath(mid + 1)
		if l != -1 && r == -1 {
			b := d.bytes[mid]
			return fmt.Sprintf("%d,%d", b.X, b.Y)
		} else if r == -1 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	panic("no path breaker found")
}
