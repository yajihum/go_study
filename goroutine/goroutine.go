package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkAPI(api string, ch chan string) {
	_, err := http.Get(api)
	if err != nil {
		// 書式設定された文字列をチェネルを通して送信
		ch <- fmt.Sprintf("ERROR: %s is down!\n", api)
		return
	}

	// 書式設定された文字列をチェネルを通して送信
	ch <- fmt.Sprintf("SUCCESS: %s is up and running!\n", api)
}

func send(ch chan string, message string) {
	ch <- message
}

// send-onlyのチャネルを引数として指定
func send2(ch chan<- string, message string) {
	fmt.Printf("Sending: %#v\n", message)
	ch <- message
}

// receive-onlyのチャネルを引数として指定
// このチャネルに対してデータを送信するとコンパイル時にエラーが発生
func read(ch <-chan string) {
	fmt.Printf("Receiving: %#v\n", <-ch)
}

func process(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "Done processing!"
}

func replicate(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "Done replicating!"
}

func main() {
	// make関数の使用によるチャネルの作成はバッファー無しのもの
	// =受信操作がある場合のみ送信操作が受け付けられる、それ以外の場合プログラムは無期限に待機しブロックされる
	// 誰かがデータを受信する準備ができるまで送信操作はバッファー無しのチャネルによってブロックされる
	ch := make(chan string) // チャネルの作成
	start := time.Now()

	apis := []string{
		"https://management.azure.com",
		"https://dev.azure.com",
		"https://api.github.com",
		"https://outlook.office.com/",
		"https://api.somewhereintheinternet.com/",
		"https://graph.microsoft.com",
	}

	// 一つずるapiに繋げられるかの処理を行うのではなく並列で処理できるようにしたい
	for _, api := range apis {
		// 関数の前にgoをつけることでcheckAPI関数がどう実行される
		go checkAPI(api, ch)
	}

	// データを受信して出力
	// apisの分だけfmt.Printを行いチャネルからのデータを出力する
	for i := 0; i < len(apis); i++ {
		fmt.Print(<-ch)
	}

	elapsed := time.Since(start)
	fmt.Printf("Done! It took %v seconds!\n", elapsed.Seconds())

	// バッファーありの場合10のようにキューのサイズを追加する
	// バッファーありの場合はプログラムがブロックされずにデータが送受信される
	// チェネルがいっぱいになるとデータを保持する領域が確保できるまで送信操作は待機
	// チャネルが空の状態で読み取り操作が起きると読み取り対象が発生するまでチャネルはブロックされる
	//size := 2 // このサイズを送信データの数より小さくすると（2とか）デットロックされる
	ch2 := make(chan string, 10)
	send(ch2, "one")
	send(ch2, "two")
	// チャネルを使用するときは常にgoroutineを使用するようにする！
	go send(ch2, "three")
	go send(ch2, "four")
	fmt.Println("All data sent to the channel ...")

	for _, api := range apis {
		go checkAPI(api, ch2)
	}

	for i := 0; i < 4; i++ {
		fmt.Println(<-ch2)
	}

	fmt.Println("Done!")

	// バッファーありとなしチャネルの使い分け
	// バッファーなし：データを送信するたびに、誰かがチャネルから読み取るまでプログラムがブロックされることが保証される
	// バッファーあり；送信操作と受信操作は分離されています。 プログラムはブロックされないがデッドロックが発生する可能性があるので注意する必要がある

	// recieveまたはsend-onlyなチャネルの利用
	ch3 := make(chan string, 1)
	send2(ch3, "Hello World!")
	read(ch3)

	ch4 := make(chan string)
	ch5 := make(chan string)
	go process(ch4)
	go replicate(ch5)

	for i := 0; i < 2; i++ {
		// selectはチャネル用
		// 処理対象のイベントを受信するまでプログラムの実行はブロックされる
		// イベントの処理が完全に終わった後実行が終了される
		select {
		case process := <-ch4:
			fmt.Println(process)
		case replicate := <-ch5:
			fmt.Println(replicate)
		}
	}
}
