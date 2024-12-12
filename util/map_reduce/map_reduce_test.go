package map_reduce

import (
	"reflect"
	"slices"
	"strings"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum ints", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("sum int64", func(t *testing.T) {
		numbers := []int64{1, 2, 3}
		got := Sum(numbers)
		want := int64(6)
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	t.Run("sum float64", func(t *testing.T) {
		numbers := []float64{1, 2, 3}
		got := Sum(numbers)
		want := 6.0
		if got != want {
			t.Errorf("got %f want %f given, %v", got, want, numbers)
		}
	})
}

func TestSumFunc(t *testing.T) {
	stringLengthFunc := func(s string) int {
		return len(s)
	}
	t.Run("sum string length", func(t *testing.T) {
		strings := []string{"sum", "of", "string", "lengths"}
		got := SumFunc(strings, stringLengthFunc)
		want := 18
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, strings)
		}
	})
}

func TestMap(t *testing.T) {
	t.Run("map to uppercase string", func(t *testing.T) {
		toUpper := func(upper string) string { return strings.ToUpper(upper) }
		words := []string{"abc", "def", "ghi"}
		got := Map(words, toUpper)
		want := []string{"ABC", "DEF", "GHI"}
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v given, %v", got, want, words)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("join words", func(t *testing.T) {
		join := func(acc string, word string) string { return acc + "," + word }
		words := []string{"abc", "def", "ghi"}
		got := Reduce(words, join, "_")
		want := "_,abc,def,ghi"
		if got != want {
			t.Errorf("got %s want %s given, %v", got, want, words)
		}
	})
}

func TestFilter(t *testing.T) {
	t.Run("filter even numbers", func(t *testing.T) {
		isEven := func(num int) bool { return num%2 == 0 }
		nums := []int{1, 2, 3, 4, 5}
		got := Filter(nums, isEven)
		want := []int{2, 4}
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v given, %v", got, want, nums)
		}
	})
}

func TestAll(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	t.Run("all values match predicate", func(t *testing.T) {
		allLessThan10 := func(num int) bool { return num < 10 }
		got := All(nums, allLessThan10)
		want := true
		if got != want {
			t.Errorf("got %v want %v given %v", got, want, nums)
		}
	})
	t.Run("all values do not match predicate", func(t *testing.T) {
		allLessThan4 := func(num int) bool { return num < 4 }
		got := All(nums, allLessThan4)
		want := false
		if got != want {
			t.Errorf("got %v want %v given %v", got, want, nums)
		}
	})
}

func TestChunkBy(t *testing.T) {
	var testCases = []struct {
		name string
		data []string
		want [][]string
	}{
		{
			name: "split in middle",
			data: []string{"abc", "def", "", "ghi"},
			want: [][]string{{"abc", "def"}, {"ghi"}},
		},
		{
			name: "multiple splits",
			data: []string{"abc", "", "def", "", "ghi"},
			want: [][]string{{"abc"}, {"def"}, {"ghi"}},
		},
		{
			name: "split at start",
			data: []string{"", "abc", "def", "ghi"},
			want: [][]string{{"abc", "def", "ghi"}},
		},
		{
			name: "split at end",
			data: []string{"abc", "def", "ghi", ""},
			want: [][]string{{"abc", "def", "ghi"}},
		},
	}
	matchWhiteSpace := func(s string) bool { return strings.TrimSpace(s) == "" }
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := ChunkBy(testCase.data, matchWhiteSpace)

			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
