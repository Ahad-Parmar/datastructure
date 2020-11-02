package main

import "fmt"

func binarysearch(array []int, val int) bool {

	low := 0
	high := len(array) - 1

	for low <= high {
		median := (low + high) / 2

		if array[median] < val {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(array) || array[low] != val {
		return false
	}

	return true
}

func main() {
	items := []int{8, 7, 10, 20, 30, 47, 45, 80, 50}
	if binarySearch(items, 30) {
		fmt.Println("list contains item")
	} else {
		fmt.Println("list doesn't contains item")
	}
}