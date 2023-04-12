package main

import "fmt"

func main() {
	// 変数の宣言の仕方

	// 普通のやつ
	// var firstName, lastName string
	// var age int

	// まとめて
	// var (
	// 	firstName, lastName string
	// 	age int
	// )

	// 宣言と同時に初期化
	// 型を指定する必要はない
	// var (
	// 	firstName = "John"
	// 	lastName = "Doe"
	// 	age = 32
	// )

	// まとめて初期化できる
	// var (
	// 	firstName, lastName, age = "John", "Doe", 32
	// )

	// これが一いちばん一般的
	// var 使わなくていい
	// 変数は使用しないと警告ができるのではなくエラーになる
	firstName, lastName := "John", "Doe"
	age := 32
	fmt.Println(firstName, lastName, age)

	// 定数
	// MixedCasedの形式やすべて大文字で記述
	//　const HTTPStatusOK = 200

	// まとめて
	// 定数は変数と違い使用しなくてもエラーにはならない
	const (
		StatusOK              = 0
		StatusConnectionReset = 1
		StatusOtherError      = 2
	)
}
