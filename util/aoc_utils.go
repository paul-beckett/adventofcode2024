package util

import (
	"fmt"
	"time"
)

func PrintResultAndTime[A any](name string, f func() A) {
	start := time.Now()
	result := f()
	total := time.Since(start)

	fmt.Printf("%s(%dms)=%v\n", name, total.Milliseconds(), result)
}
