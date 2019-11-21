package main

import "fmt"

func main() {
	table := booleanArray(10, 10)
	for i := range table {
		fmt.Println(table[i])
	}
}

func booleanArray(rows, columns int) [][]bool {
	arr := make([][]bool, rows)
	for y := range arr {
		arr[y] = make([]bool, columns)
		for x := range arr[y] {
			r := gcd(y, x)
			if r == 1 {
				arr[y][x] = true
			} else {
				arr[y][x] = false
			}
		}
	}
	return arr
}

func gcd(p, q int) int {
	if q == 0 {
		return p
	}
	r := p % q
	return gcd(q, r)
}
