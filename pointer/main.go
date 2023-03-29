package main

import "fmt"

func PointerInt(s *int) int {
	*s = *s + 1
	return *s
}

func PointerBoll(s *bool) bool {
	*s = !*s
	return *s
}


func PointerSlicer(s *[]int) {
	 for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
				fmt.Println(i, j)
        (*s)[i], (*s)[j] = (*s)[j], (*s)[i]
    }

}

func PointerString(s *string) string{
	runes := []rune(*s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*s = string(runes)
	return *s
}
func PointerMatrizMulti(s *[2][2]int) {

	var soma int
	var media int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s[i]); j++ {
			soma += s[i][j]
		}
	}
	media = soma / (len(s) * len(s[0]))
	fmt.Println(media)
}



func main() {
	// s := 0
	// s = PointerInt(&s)
	// print(s)

	// s := true
	// s = PointerBoll(&s)
	// print(s) 

	// s := []int{1, 2, 3}
	// PointerSlicer(&s)
	// fmt.Println(s)

	// s := "hello"
	// s = PointerString(&s)

	// fmt.Println(s)

	// s := [2][2]int{{1, 1}, {0, 0}}
	// PonterMatrizMulti(&s)
	// fmt.Println(s)

}