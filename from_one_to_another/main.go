package main

import "fmt"

func reverse(ar []int) []int {
	i := 0
	j := len(ar) - 1
	for i < j {
		ar[i], ar[j] = ar[j], ar[i]
		i++
		j--
	}
	return ar
}

func main() {

	a := []int{1, 2, 3, 4, 5}
	b := a
	fmt.Println(a)
	fmt.Println(reverse(b))
}
