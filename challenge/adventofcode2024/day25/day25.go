package day25

type Day25 struct {
	keys  [][]int
	locks [][]int
}

func newDay25(data []string) *Day25 {

	var keys [][]int
	var locks [][]int

	for i := 0; i < len(data); i += 8 {

		if data[i] == "#####" && data[i+6] == "....." {
			lock := make([]int, 5)
			for j := i; j < i+6; j++ {
				for k := 0; k < 5; k++ {
					if data[j][k] == '.' {
						lock[k]++
					}
				}
			}
			locks = append(locks, lock)
		} else {
			key := make([]int, 5)
			for j := i; j < i+6; j++ {
				for k := 0; k < 5; k++ {
					if data[j][k] == '#' {
						key[k]++
					}
				}
			}
			keys = append(keys, key)
		}

	}

	return &Day25{keys: keys, locks: locks}
}

func keyFits(key []int, lock []int) bool {
	for i, k := range key {
		if lock[i] < k {
			return false
		}
	}
	return true
}

func (d *Day25) part1() int {
	total := 0
	for _, key := range d.keys {
		for _, lock := range d.locks {
			if keyFits(key, lock) {
				total++
			}
		}
	}
	return total
}

func (d *Day25) part2() string {
	return "there is no part 2!"
}
