package helpers

func GetPermutations[T any](collection []T) [][]T {
	return permute(collection, 0)
}

func permute[T any](collection []T, position int) [][]T {
	l := len(collection)
	if l == position {
		tmp := make([]T, l)
		copy(tmp, collection)
		return [][]T{tmp}
	}

	result := [][]T{}
	for i := position; i < l; i++ {
		swap(collection, position, i)
		result = append(result, permute(collection, position+1)...)
		swap(collection, position, i)
	}

	return result
}

func swap[T any](collection []T, p1, p2 int) {
	tmp := collection[p1]
	collection[p1] = collection[p2]
	collection[p2] = tmp
}
