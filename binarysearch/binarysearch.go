package main

import "fmt"

func main() {
	a := []int{2, 3, 5, 6, 7, 8, 9, 22}
	fmt.Println(search(a, 22))
}

func search(a []int, key int) int {
	return searchHelper(a, 0, len(a), key)
}

func searchHelper(a []int, lo, hi, key int) int {
	mid := (hi + lo) / 2
	if hi <= lo {
		return -1
	}
	if a[mid] == key {
		return mid
	} else if a[mid] > key {
		return searchHelper(a, lo, mid, key)
	} else {
		return searchHelper(a, mid, hi, key)
	}
}
