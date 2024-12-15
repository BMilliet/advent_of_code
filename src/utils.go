package src

import "fmt"

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		fmt.Println("Index out of bounds")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}
