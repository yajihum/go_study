package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

// priceを出力するときにカスタム文字列出力メソッドを使うことで$と一緒に出力される
func (d dollars) String() string {
	return fmt.Sprintf("$%.2f", d)
}

type database map[string]dollars

// これもカスタム型
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		// Fが先頭についているものは書き込み先を明示的に指定できる
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func main() {
	db := database{"Go T-Shirt": 25, "Go Jacket": 55}
	// localhost:8000を開くと「Go T-Shirt: $25.00」とかが出力されている
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
