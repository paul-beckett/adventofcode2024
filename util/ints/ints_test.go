package ints

import (
	"slices"
	"testing"
)

func TestToInts(t *testing.T) {
	comma := func(r rune) bool { return r == ',' }
	var testCases = []struct {
		name       string
		data       string
		fieldsFunc func(rune) bool
		want       []int
	}{
		{
			name:       "csv",
			data:       "1,2,34,,567",
			fieldsFunc: comma,
			want:       []int{1, 2, 34, 567},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			got := ToInts(testCase.data, testCase.fieldsFunc)

			if !slices.Equal(got, testCase.want) {
				t.Errorf("got %v, want %v", got, testCase.want)
			}
		})
	}
}
