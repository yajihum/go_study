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

	// 最後の要素のみを指定して初期化もできる
	numbers := [...]int{99: -1}
	fmt.Println("First Position:", numbers[0]) // 0が出力される
	fmt.Println("Last Position:", numbers[99]) // -1と出力
	fmt.Println("Length:", len(numbers))       // 100と出力

	// 多次元配列
	var twoD [3][5]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			twoD[i][j] = (i + 1) * (j + 1)
		}
		fmt.Println("Row", i, twoD[i])
	}
	fmt.Println("\nAll at once:", twoD) // All at once: [[1 2 3 4 5] [2 4 6 8 10] [3 6 9 12 15]]
}
