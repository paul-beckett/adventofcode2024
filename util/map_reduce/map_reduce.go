package map_reduce

type Number interface {
	int | int64 | float64
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	var result = initialValue
	for _, x := range collection {
		result = f(result, x)
	}
	return result
}

func Map[A, B any](collection []A, f func(A) B) []B {
	var mapped []B
	for _, x := range collection {
		result := f(x)
		mapped = append(mapped, result)
	}
	return mapped
}

func Sum[A Number](numbers []A) A {
	add := func(x A) A { return x }
	return SumFunc(numbers, add)
}

func SumFunc[A any, B Number](collection []A, f func(A) B) B {
	add := func(acc B, x A) B { return acc + f(x) }
	return Reduce(collection, add, 0)
}

func Filter[A any](collection []A, f func(A) bool) []A {
	var filtered []A
	for _, x := range collection {
		if f(x) {
			filtered = append(filtered, x)
		}
	}
	return filtered
}

func All[A any](collection []A, f func(A) bool) bool {
	for _, x := range collection {
		if !f(x) {
			return false
		}
	}
	return true
}

func ChunkBy[A any](collection []A, f func(A) bool) [][]A {
	var chunks [][]A
	var chunk []A
	for _, x := range collection {
		if f(x) {
			if chunk != nil {
				chunks = append(chunks, chunk)
			}
			chunk = nil
		} else {
			chunk = append(chunk, x)
		}
	}
	if chunk != nil {
		chunks = append(chunks, chunk)
	}
	return chunks
}
