package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// defer関数；defer ステートメントが含まれる関数が終了するまで関数の実行が延期される
	// 以下では、「regular 1〜regular 4」まで出力された後に「deferred -4」〜「deferred-1」が順に出力される
	// 後入先出しになる、遅延実行
	for i := 1; i <= 4; i++ {
		defer fmt.Println("deferred", -i)
		fmt.Println("regular", i)
	}

	// 例えば「ファイルの使用が終了した後でファイルを閉じる」ことなどをしたいときに使われる

	// Create関数の戻り値は(File, error)なのでそれを新しい変数に代入している
	newfile, error := os.Create("learnGo.txt")
	// nilとは値が存在しないことを表す値であり、型も区別される
	// 以下のエラーに対するnilは、エラー型の変数errorが存在しないことを表す
	if error != nil {
		fmt.Println("Error: Could not create file.")
		return
	}
	// ファイルを作成したり開いたりした後にファイルを閉じることを忘れないように遅延実行させる
	defer newfile.Close()

	if _, error = io.WriteString(newfile, "Learning Go!"); error != nil {
		fmt.Println("Error: Could not write to file.")
		return
	}
	newfile.Sync()
}
