package main

func func1(s *int) {
	*s += 1
}


func main(){
	var x int
	func1(&x)
	println(x)

}