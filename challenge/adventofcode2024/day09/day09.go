package day09

import (
	"strconv"
)

type Day09 struct {
	data []int
}

func newDay09(data []string) *Day09 {
	var nums []int
	for _, c := range data[0] {
		num, _ := strconv.Atoi(string(c))
		nums = append(nums, num)
	}
	return &Day09{
		data: nums,
	}
}

type fs []int

func toFs(data []int) fs {
	var fs []int
	for i := 0; i < len(data); i++ {
		for j := 0; j < data[i]; j++ {
			fs = append(fs, i/2)
		}
		i++
		if i < len(data) {
			for j := 0; j < data[i]; j++ {
				fs = append(fs, -1)
			}
		}
	}
	return fs
}

func (f fs) checksum() int {
	total := 0
	for block, id := range f {
		if id != -1 {
			total += id * block
		}
	}

	return total
}

func (f fs) firstEmpty(size int) int {
	for i := 0; i < len(f); i++ {
		for f[i] != -1 {
			i++
		}

		start := i
		for i < len(f) && f[i] == -1 {
			i++
		}
		end := i
		if (end - start) >= size {
			return start
		}
	}
	return len(f)
}

func (d *Day09) part1() int {
	fs := toFs(d.data)

	l := 0
	r := len(fs) - 1

	for l < r {
		for fs[l] != -1 {
			l++
		}
		for fs[r] == -1 {
			r--
		}
		if l < r {
			fs[l] = fs[r]
			fs[r] = -1
			l++
			r--
		}
	}
	return fs.checksum()
}

func (d *Day09) part2() int {
	fs := toFs(d.data)

	r := len(fs) - 1
	id := fs[r]
	for id > 0 {
		for fs[r] != id {
			r--
		}
		end := r
		for fs[r-1] == id {
			r--
		}
		start := r
		fileLen := end - start + 1

		insertion := fs.firstEmpty(fileLen)
		if insertion < r {
			for i := 0; i < fileLen; i++ {
				fs[insertion+i] = fs[start+i]
				fs[start+i] = -1
			}
		}
		id--
	}
	return fs.checksum()
}
