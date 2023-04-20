package main

import (
	"fmt"
)

func main() {
	// 配列
	var a [3]int      // 配列の長さは変えられない
	a[1] = 10         // 添え字でアクセス
	fmt.Println(a[0]) // 0と出力→int型の配列の場合、規定値が0で初期化されているためエラーにならずアクセスできる
	fmt.Println(a[1]) // 10と出力
	// len()で配列やスライス、マップの要素数を取得するもの
	fmt.Println(a[len(a)-1]) // 0と出力

	// 配列の初期化
	// 全要素を指定しておく必要はない
	// ただし最後の要素は空の配列になるので空白が入ってしまう
	cities := [5]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities)

	// 要素数を省略できる
	// こっちは最後に空白はできない
	cities1 := [...]string{"New York", "Paris", "Berlin", "Madrid"}
	fmt.Println("Cities:", cities1)
}
