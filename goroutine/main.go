package main

import (
	"fmt"
	"sync"
)

func fun1() {
	fmt.Printf("hello word")
}

func reciveNumber1(value int, ch chan<- int) {
	ch <- value
}

func reciveNumber2(value int, ch chan<- int) {
	ch <- value
}

func sum(ch1 <-chan int, ch2 <-chan int) {
	fmt.Println(<-ch1 + <-ch2)
}

func printString(str string, wg *sync.WaitGroup) {
	defer wg.Done() // sinaliza ao WaitGroup que a goroutine foi concluída
	fmt.Println(str)
}

func sendNumberOdd(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i += 2 {
		ch <- i
	}
	close(ch)
}

func sendNumberEven(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= 10; i += 2 {
		ch <- i
	}
	close(ch)
}

func printNumber(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		fmt.Println(v)
	}
}

func sumNumbers(ch <-chan int, wg *sync.WaitGroup) int {
	defer wg.Done() // sinaliza ao WaitGroup que a goroutine foi concluída

	sum := 0
	for i := range ch {
		sum += i
	}

	return sum
}

func main() {
	// trigger func1 every 5 sec with gourtine
	// go fun1()
	// wait for 5 sec

	// create a channel
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go reciveNumber1(10, ch1)
	// go reciveNumber2(20, ch2)

	// go sum(ch1, ch2)

	// time.Sleep(5 * time.Second)

	// strs := []string{"Olá", "Mundo", "Goroutines", "São", "Demais"}

	// var wg sync.WaitGroup // cria um WaitGroup para aguardar a conclusão das goroutines

	// for _, str := range strs {
	// 	wg.Add(1) // incrementa o WaitGroup para cada goroutine criada
	// 	go printString(str, &wg)
	// }
	// ch := make(chan int)

	// var wg sync.WaitGroup

	// wg.Add(2)

	go sendNumberOdd(ch, &wg)
	go sendNumberEven(ch, &wg)
	go printNumber(ch, &wg)

	// wg.Wait() // aguarda a conclusão de todas as goroutines antes de encerrar o programa
	ch := make(chan int) // cria um canal de inteiros

	var wg sync.WaitGroup // cria um WaitGroup para aguardar a conclusão das goroutines

	// cria várias goroutines para enviar valores para o canal
	for i := 1; i <= 10; i++ {
		wg.Add(1) // incrementa o WaitGroup para cada goroutine criada

		go func(num int) {
			ch <- num
			wg.Done() // sinaliza ao WaitGroup que a goroutine foi concluída
		}(i)
	}

	go func() {
		wg.Wait() // aguarda a conclusão de todas as goroutines antes de fechar o canal
		close(ch) // fecha o canal quando todas as goroutines terminarem
	}()

	sum := sumNumbers(ch, &wg) // chama a função que calcula a soma dos valores

	fmt.Println("Soma:", sum) // impr

}
