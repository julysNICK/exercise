package main

import (
	"fmt"
	"sync"
	"time"
)

func func1() {
	fmt.Println("hello world")
}

func func2(arrayString []string) {
	for _, v := range arrayString {
		fmt.Println(v)
	}
}

func func3(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		ch <- i
	}
	close(ch)
}

func func4(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 10; i += 2 {
		ch <- i
	}
	close(ch)
}

func func5(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

}

func main() {

	go func1()
	time.Sleep(1 * time.Second)

	go func2([]string{"hello", "world", "!"})

	// fmt.Scanln()

	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go func3(ch, &wg)
	go func4(ch, &wg)
	go func5(ch, &wg)

	wg.Wait()

}
