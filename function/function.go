package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 関数について

	// os.Argsでコマンドライン引数を利用できる
	// 例えばgo run function.go 3 5
	// で実行すれば Sum: 8 が出力される
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Sum:", number1+number2)

	sum := sum(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum)

	sum2 := sum2(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum2)

	// 複数の戻り値がある関数の受け取り方
	sum3, mul := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum3)
	fmt.Println("Mul:", mul)

	// 戻り値が複数ある場合で「_」を使うとそれを無視できる
	// ↓の例ではmulを無視できる
	sum4, _ := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum4)
}

// カスタム関数
func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi((number1))
	int2, _ := strconv.Atoi((number2))
	return int1 + int2
}

// 戻り値に名前をつけられる（result）
func sum2(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi((number1))
	int2, _ := strconv.Atoi((number2))
	result = int1 + int2
	return
}

// 複数の戻り値を設定できる
// 戻り値の名前は省略可能
func calc(number1 string, number2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}
