package main

import (
	"fmt"
)

func givemeanumber() int {
	return -1
}

func main() {
	x := 27

	// ifの使い方
	// 条件式部分に括弧は必要ない
	// 三項演算子は使えない
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	// Goではifのブロック内で変数を宣言するのが一般的
	// 以下ではnumに関数の戻り値を代入し、そのnum変数を{}内で使用している
	// このnuはifブロックの外側では使えない
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
