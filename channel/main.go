package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func chanelExample(value int, c chan int) {
	c <- value
}

func doubleChannel(c1 chan int, c2 chan int) {
	value := <-c1
	c2 <- value * 2
}

func randomValueWithChannel(c chan int, n int) {
	for i := 0; i < n; i++ {
		c <- i
	}
}

func stringListChannel(c chan string, s []string) {

	for i := len(s) - 1; i >= 0; i-- {
		c <- s[i]
	}
}

func createNChannel(c chan<- int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		c <- rand.Intn(100)
	}
}

var wg sync.WaitGroup

func createNChannel1(c chan int, n int) {
	wg.Add(1)
	defer wg.Done()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		c <- rand.Intn(100)
	}
}

func main() {
	//ele conecta a função main com a função chanelExample
	// channel := make(chan int)

	// go chanelExample(10, channel)

	// value := <- channel

	// println(value)

	//double channel
	// channel1 := make(chan int)
	// channel2 := make(chan int)

	// go doubleChannel(channel1, channel2)

	// channel1 <- 10

	// value := <- channel2

	// println(value)

	//random value with channel

	// channel := make(chan int)

	// go randomValueWithChannel(channel, 10)

	// for i := 0; i < 10; i++ {
	// 	println(<-channel)
	// }

	// //string list channel
	// channel := make(chan string)

	// go stringListChannel(channel, []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"})

	// for i := 0; i < 10; i++ {
	// 	//show in decrescent order
	// 	println(<-channel)

	// }

	//create n channel
	ch := make(chan int)
	n := 10

	// Função para enviar valores aleatórios para o channel
	send := func(ch chan<- int) {
		rand.Seed(time.Now().UnixNano())
		for i := 0; i < 10; i++ {
			ch <- rand.Intn(10)
		}
	}

	// Criação das goroutines para enviar valores aleatórios para o channel
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			send(ch)
		}()
	}

	// Função para receber os valores do channel e calcular a soma total
	receive := func(ch <-chan int) {
		var sum int
		for v := range ch {
			sum += v
		}
		fmt.Println("Soma total:", sum)
	}

	// Criação da goroutine para receber os valores do channel e calcular a soma total
	go func() {
		wg.Wait()
		close(ch)
	}()
	receive(ch)

}
