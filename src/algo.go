package src

func createFrequencyMap(list []int) map[int]int {
	fr := make(map[int]int)

	for _, e := range list {
		fr[e]++
	}

	return fr
}

func quickSort(list []int) []int {
	if len(list) == 0 {
		return list
	}

	pivotIndex := len(list) / 2
	pivot := list[pivotIndex]

	rawList := append(list[:pivotIndex], list[pivotIndex+1:]...)

	left := []int{}
	right := []int{}
	for i := 0; i < len(rawList); i++ {
		e := rawList[i]

		if e > pivot {
			right = append(right, e)
		} else {
			left = append(left, e)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}
