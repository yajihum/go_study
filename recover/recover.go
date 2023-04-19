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
	// 匿名関数
	defer func() {
		// recover関数でパニックの後で制御を取り戻すことができる
		// クラッシュする前にそれを避けることができる
		// recoverの呼び出しはdeferを呼び出している関数ないでのみ行う
		handler := recover()
		if handler != nil {
			fmt.Println("main(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")

	// Goではpanicとrecoverを使ってtry catchを行う
}
