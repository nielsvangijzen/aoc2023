package util

func MaxValue[T Ordered](slice []T) T {
	mx := slice[0]

	for _, item := range slice {
		if item > mx {
			mx = item
		}
	}

	return mx
}

func MinValue[T Ordered](slice []T) T {
	mn := slice[0]

	for _, item := range slice {
		if item < mn {
			mn = item
		}
	}

	return mn
}

func Sum[T Addable](slice []T) (sum T) {
	for _, value := range slice {
		sum += value
	}

	return
}

func ExistsInSlice[T comparable](haystack []T, needle T) bool {
	for _, item := range haystack {
		if item == needle {
			return true
		}
	}

	return false
}

func Unique[T comparable](haystack []T) bool {
	set := map[T]int{}

	for _, value := range haystack {
		set[value] += 1
	}

	return len(set) == len(haystack)
}
