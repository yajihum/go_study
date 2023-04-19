package main

import "fmt"

func highlow(high int, low int) {
	// 範囲外のインデックスを使用して配列にアクセスしようとしたり、nil ポインターを逆参照したりして、実行時エラーが発生すると、Go プログラムはパニック状態になる
	// 逆にパニック状態を引き起こすこともできる。パニック状態を引き起こすことで遅延実行される関数が呼び出されエラーメッセージを吐き出し、クラッシュできる
	if high < low {
		// lowの値がheighの値より大きくなた時点でパニックが発生
		fmt.Println("Panic!")
		// 引数にエラーメッセージを入れる
		panic("highlow() low greater than high")
	}
	defer fmt.Println("Deferred: highlow(", high, ",", low, ")")
	// 正常な時にCallが実行される
	fmt.Println("Call: highlow(", high, ",", low, ")")

	highlow(high, low+1)
}

func main() {
	highlow(2, 0)
	// 例外が起きた場合は以下は実行されない
	fmt.Println("Program finished successfully!")
}
