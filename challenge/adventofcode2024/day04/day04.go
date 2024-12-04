package day04

type Day04 struct {
	data []string
}

func newDay04(data []string) *Day04 {
	return &Day04{data: data}
}

func (d *Day04) part1() int {
	count := 0
	word := "XMAS"
	deltas := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}
	for y, line := range d.data {
		for x := range line {
			for _, delta := range deltas {
				for i := 0; i < len(word); i++ {
					nextY := y + delta[0]*i
					nextX := x + delta[1]*i
					if nextX < 0 || nextX >= len(line) || nextY < 0 || nextY >= len(d.data) {
						break
					}
					if d.data[nextY][nextX] != word[i] {
						break
					}
					if i == len(word)-1 {
						count++
					}
				}
			}
		}
	}
	return count
}

func (d *Day04) part2() int {
	masCount := 0
	for y := 1; y < len(d.data)-1; y++ {
		for x := 1; x < len(d.data[y])-1; x++ {
			if d.data[y][x] != 'A' {
				continue
			}
			freq := make(map[uint8]int)
			freq[d.data[y-1][x-1]]++
			freq[d.data[y-1][x+1]]++
			freq[d.data[y+1][x-1]]++
			freq[d.data[y+1][x+1]]++
			if freq['M'] == 2 && freq['S'] == 2 && d.data[y-1][x-1] != d.data[y+1][x+1] {
				masCount++
			}
		}
	}
	return masCount
}
