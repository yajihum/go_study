package main

import (
	"fmt"
	"math/rand"

	"github.com/yajium/calculator"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))

	// パッケージの使用
	total := calculator.Sum(3, 5)
	fmt.Println((total))
	fmt.Println("Version: ", calculator.Version)
}
