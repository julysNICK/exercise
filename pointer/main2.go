package main

import "fmt"

func func1(s *int) {
	fmt.Println("s value:", *s)
	fmt.Println("s address:", &s)
}

func func2(s *int) int {
	return 2 * *s
}

func func3(slice *[]int) int {
	lower := 0

	size := len(*slice)

	for i := 0; i < size; i++ {
		valueActual := (*slice)[i]
		if valueActual < (*slice)[lower] {
			lower = i
		}
	}
	return (*slice)[lower]
}

func func4(slice *[]int, value int) []int {

	newSlice := make([]int, 0, len(*slice)+1)
	size := len(*slice)

	for i := 0; i < size; i++ {
		if (*slice)[i] < value {
			newSlice = append(newSlice, (*slice)[i])
		}

	}
	return newSlice

}

func binarySearchTree(slice *[]int, value int) int {

	lower := 0
	upper := len(*slice) - 1
	middle := 0

	for lower <= upper {
		fmt.Println("lower:", lower, "upper:", upper)

		middle = (lower + upper) / 2
		fmt.Println("middle:", middle)
		fmt.Println("value:", (*slice)[middle])
		if (*slice)[middle] == value {
			return middle
		} else if (*slice)[middle] < value {
			lower = middle + 1
		} else {
			upper = middle - 1
		}
	}
	return -1
}

func main() {
	// s := 2

	// func1(&s)

	// s = func2(&s)
	// fmt.Println(s)

	// slice := []int{9, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Println(func3(&slice))

	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// fmt.Println(func4(&slice, 5))

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(binarySearchTree(&slice, 9))
}
