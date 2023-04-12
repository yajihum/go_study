package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	// os.Argsでコマンドライン引数を利用できる
	// 例えばgo run function.go 3 5
	// で実行すれば Sum: 8 が出力される
	number1, _ := strconv.Atoi(os.Args[1])
	number2, _ := strconv.Atoi(os.Args[2])
	fmt.Println("Sum:", number1+number2)
}
