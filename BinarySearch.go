package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func partition(a []int) int {
	hi := len(a) - 1
	pivot := a[hi]
	i := -1

	for j := 0; j < hi; j++ {
		if a[j] <= pivot {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	i++
	a[i], a[hi] = a[hi], a[i]
	return i
}

func quicksort(a []int) {
	if len(a) <= 1 {
		return
	}
	p := partition(a)
	quicksort(a[:p])
	quicksort(a[p:])
}

func make_random_slice(num_items, max int) []int {
	var rand_slice []int

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num_items; i++ {
		rand_slice = append(rand_slice, rand.Intn(max))
	}

	return rand_slice
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func print_slice(arr []int, num_items int) {
	num_elems := Min(len(arr), num_items)
	fmt.Println(arr[:num_elems])
}

func binary_search(values []int, target int) (index, num_tests int) {
	l := 0
	r := len(values) - 1
	n_tests := 0

	for l <= r {
		n_tests++
		m := int(math.Floor(float64((l + r) / 2)))
		if values[m] < target {
			l = m + 1
		} else if values[m] > target {
			r = m - 1
		} else {
			return m, n_tests
		}
	}
	return -1, len(values)
}

func main() {
	var num_items, max, target int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted slice.
	values := make_random_slice(num_items, max)

	quicksort(values)

	print_slice(values, len(values))
	fmt.Println()

	fmt.Printf("Target: ")
	fmt.Scanln(&target)

	i, v := binary_search(values, target)

	if i == -1 {
		fmt.Printf("Target %d not found, %d tests", target, v)
	} else {
		fmt.Printf("values[%d] = %d, %d tests", i, target, v)
	}
}
