package main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d recebeu %d\n", workerID, x)
		time.Sleep(time.Second)
	}
}

func main() {
	var qtdWorker, qtdExecutions int

	fmt.Print("Digite a quantidade de workers: ")
	fmt.Scan(&qtdWorker)

	fmt.Print("Digite a quantidade de execuções: ")
	fmt.Scan(&qtdExecutions)
	
	canal := make(chan int)

	for i := 1; i <= qtdWorker; i++ {
		go worker(i, canal)
	}

	for i := 0; i < qtdExecutions; i++ {
		canal <- i
	}
}
