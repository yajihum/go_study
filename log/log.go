package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.SetPrefix("main(): ")           // プログラムのログエッセー時に接頭語を付けられる　今回の場合「main():」と言うのがログメッセージ毎に先頭につく
	log.Print("Hey, I'm a log!")        // log.Print：最初に日付と時刻が出力される
	log.Fatal("Hey, I'm an error log!") // log.Fatal：os.Exit(1)を使用した時と同様にエラーをログに記録してプログラムを終了できる
	log.Panic("Hey, I'm an error log!") // panicを発生させてFatalと同様にプログラムを終了できる
	fmt.Print("Can you see me?")

	// ファイルにログを記録する
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// ファイルが開ける状態になったら全ての出力をファイルに送信するようにする
	log.SetOutput(file)
	log.Print("Hey, I'm a log!")

	// ログのライブラリ：zerolog
	// JSON形式でログが出力される
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Hey! I'm a log message!")

	log.Debug().
		Int("EmployeeID", 1001).
		Msg("Getting employee information")

	log.Debug().
		Str("Name", "John").
		Send()
}
