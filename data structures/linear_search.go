package main

import "fmt"

func linearsearch(datalist []int, key int) bool {
	for _, item := range datalist {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	data := []int{45, 454, 445, 566, 867, 254, 748, 756, 488, 998, 219}
	if linearsearch(data, 748) {
		fmt.Println("list contains item")
	} else {
		fmt.Println("list doesn't contain item")
	}
}