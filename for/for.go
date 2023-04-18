package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// for文
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum of 1..100 is", sum)

	// for文で前ステートメントと後ステートメントを使わない場合はwhile文と同様のことができる
	// 以下ではnumが5になるまでfor文の中身の処理が実行される
	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}

	// 無限ループ
	// 抜けるにはbreakを使う
	var num1 int32
	sec := time.Now().Unix()
	rand.Seed(sec)

	for {
		fmt.Print("Writing inside the loop...")
		if num1 = rand.Int31n(10); num1 == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(num1)
	}

	// continueを使うことでループの現在の反復をスキップできる
	// 以下ではnumが5で割り切れる場合はsum1に加算されないようにしている
	sum1 := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum1 += num
	}
	fmt.Println("The sum of 1 to 100, but excluding numbers divisible by 5, is", sum1)
}
