package main

import "fmt"

func insertionSort(array []int) {
	for pos := 1; pos < len(array); pos++ {
		i := pos -1
		value := array[pos]
		for i >= 0 && array[i] > value {
			array[i + 1] = array[i]
			i--
		}
		array[i + 1] = value
	}
}

func main() {
	array := []int{3, 2, 1, 5, 10, 7, 9}
	fmt.Println(array)
	insertionSort(array)
	fmt.Println(array)
}
