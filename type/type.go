package main

import (
	"fmt"
	"strconv"
)

func main() {
	// データ型
	// 参考；https://go.dev/src/builtin/builtin.go

	// int 整数
	var integer8 int8 = 127
	var integer16 int16 = 32767
	var integer32 int32 = 2147483647
	var integer64 int64 = 9223372036854775807
	fmt.Println(integer8, integer16, integer32, integer64)

	// 異なるデータ型はエラー
	// var integer16 int16 = 127
	// var integer32 int32 = 32767
	// fmt.Println(integer16 + integer32)

	// rune ： Unicode文字
	// Gではなく71が出力される
	// ""ではなく''
	rune := 'G'
	fmt.Println(rune)

	// 浮動小数点
	var float32Num float32 = 2147483647
	var float64Num float64 = 9223372036854775807
	fmt.Println(float32Num, float64Num)

	// ブール値
	var featureFlag bool = true
	fmt.Println((featureFlag))

	// 文字列
	// ''は1つの文字に使用する
	var firstName string = "John"
	lastName := "Doe"
	fmt.Println(firstName, lastName)

	// 既定値
	// 0 0 0 false "" が出力される
	var defaultInt int
	var defaultFloat32 float32
	var defaultFloat64 float64
	var defaultBool bool
	var defaultString string
	fmt.Println(defaultInt, defaultFloat32, defaultFloat64, defaultBool, defaultString)

	// 型変換
	var integer16Num int16 = 127
	var integer32Num int32 = 32767
	fmt.Println(int32(integer16Num) + integer32Num)

	// strconvを使ってstringをintへ
	// _はその変数の値を使用しないことを指す
	i, _ := strconv.Atoi("-42")
	s := strconv.Itoa(-42)
	fmt.Println(i, s)
}
