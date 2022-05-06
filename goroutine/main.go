package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {

	waitGroup.Add(1) // main goroutinin dışında 1 tane rutin bekle diyoruz

	go printX() // rutini çalıştırıyoz, async

	waitGroup.Wait() // beklenecek async işlemleri waitten önce koyuyoz

	printTire() // sync func

	time.Sleep(time.Second) // main de async olduğu için go routinleri beklemeden kapanıyor, sonucu consolda görebilmek için sleep ediyoz.
}

func printX() {
	for i := 0; i < 100; i++ {
		fmt.Print("X")
	}
	waitGroup.Done()
}
func printTire() {
	for i := 0; i < 1000; i++ {
		fmt.Print("-")
	}
}
