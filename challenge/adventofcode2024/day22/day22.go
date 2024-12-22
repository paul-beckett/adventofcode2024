package day22

import (
	"fmt"
	"strconv"
)

type Day22 struct {
	secrets []int
}

func newDay22(data []string) *Day22 {
	var secrets []int
	for _, line := range data {
		num, _ := strconv.Atoi(line)
		secrets = append(secrets, num)
	}
	return &Day22{secrets: secrets}
}

func mix(a, b int) int {
	return a ^ b
}

func prune(n int) int {
	return n % 16_777_216
}

func nextSecret(n int) int {
	n = mix(n, n*64)
	n = prune(n)

	n = mix(n, n/32)
	n = prune(n)

	n = mix(n, n*2048)
	n = prune(n)
	return n
}

func (d *Day22) part1() int {
	total := 0
	for _, secret := range d.secrets {
		current := secret
		for i := 0; i < 2000; i++ {
			current = nextSecret(current)
		}
		total += current
	}
	return total
}

func price(n int) int {
	return n % 10
}

func toKey(nums []int) string {
	return fmt.Sprint(nums[len(nums)-4:])
}

func (d *Day22) part2() int {
	pricesByKey := make(map[string]int)
	for _, secret := range d.secrets {
		seenForMonkey := make(map[string]bool)
		current := secret
		var priceChanges []int
		for i := 0; i < 2000; i++ {
			before := price(current)
			current = nextSecret(current)
			after := price(current)
			priceChanges = append(priceChanges, after-before)
			if i >= 3 {
				key := toKey(priceChanges)
				if !seenForMonkey[key] {
					pricesByKey[key] += after
					seenForMonkey[key] = true
				}
			}
		}
	}

	m := 0
	for _, p := range pricesByKey {
		m = max(m, p)
	}
	return m
}
