package main

import (
	"fmt"
	"strconv"
)

func findprimes(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}

	if number != 1 {
		return true
	} else {
		return false
	}
}

func main() {
	// FizzBuzz
	for i := 1; i <= 100; i++ {
		switch {
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		}
		fmt.Println(strconv.Itoa(i))
	}

	// 素数かどうかのチェック
	fmt.Println("Prime numbers less than 20:")

	for number := 1; number <= 20; number++ {
		if findprimes(number) {
			fmt.Printf("%v ", number)
		}
	}

	// パニックにする
	val := 0
	for {
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)

		// switchのがいいかも
		if val < 0 {
			fmt.Println("Panic!")
			// 引数にエラーメッセージを入れる
			panic("val is under 0")
		} else if val == 0 {
			fmt.Println(val, "is neither negative nor positive")
		} else {
			fmt.Println("You entered:", val)
		}
	}
}
