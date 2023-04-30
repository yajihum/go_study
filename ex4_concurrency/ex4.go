package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func fib(ch chan string, number float64) {
	x, y := 1.0, 1.0
	for i := 0; i < int(number); i++ {
		x, y = y, x+y
	}

	r := rand.Intn(3)
	time.Sleep(time.Duration(r) * time.Second)

	ch <- fmt.Sprintf("Fib(%v): %v\n", number, x)
}

var quit = make(chan bool)

func calcFib(c chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
			log.Print(x, y)
		case <-quit:
			fmt.Println("Done calculating Fibonacci!")
			return
		}
	}
}

func main() {
	start1 := time.Now()

	command := ""
	data := make(chan int)

	log.Print("start")
	go calcFib(data)

	/*
		Q：このプログラムで最初の1回目以降、チャネルからのデータが、フィボナッチ数の計算より先に行われるのはなぜでしょうか？
		A：calcFibゴルーチンがチャネルにデータを送信すると、そのデータが他のゴルーチンによって受け取られるまでcalcFibゴルーチンはブロック（つまり停止）される。
		calcFibゴルーチンがフィボナッチ数を計算し、それをチャネルに送信した後、次のフィボナッチ数の計算を開始する前に、その数値がmainゴルーチンによって受け取られなければなりません。
		これにより、フィボナッチ数の計算が次の計算に進む前に、毎回mainゴルーチンでその数値が表示されることが保証される。
	*/
	for {
		num := <-data
		log.Print(num)
		fmt.Scanf("%s", &command)
		if command == "quit" {
			quit <- true
			break
		}
	}

	// calcFib関数が完全に終了するのを待つために1秒間停止
	time.Sleep(1 * time.Second)

	elapsed1 := time.Since(start1)
	fmt.Printf("Done! It took %v seconds!\n", elapsed1.Seconds())
}
