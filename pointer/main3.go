package main

import "fmt"

func func1(s *int) {
	fmt.Println("s value:", 2**s)
}

func func2(s *string) {
	// invert string

	invertString := []rune(*s)

	for i, j := 0, len(invertString)-1; i < j; i, j = i+1, j-1 {
		invertString[i], invertString[j] = invertString[j], invertString[i]
	}

	*s = string(invertString)

	fmt.Println("s value:", *s)

}

func func3(s *int, slice []int) {
	higher := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] > higher {
			higher = slice[i]
		}
	}
	*s = higher
	fmt.Println("s value:", *s)
}
func main() {
	// s := 10
	// func1(&s)

	// s := "hello"
	// func2(&s)

	// s := 0

	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// func3(&s, slice)
}
