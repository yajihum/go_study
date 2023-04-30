package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
)

// インターフェースの定義
// このShapeインターフェースに必要なPerimeterとAreaメソッドを継承先で実装する必要がある
type Shape interface {
	Perimeter() float64
	Area() float64
}

// Square
type Square struct {
	size float64
}

// Square
func (s Square) Area() float64 {
	return s.size * s.size
}

// Square
func (s Square) Perimeter() float64 {
	return s.size * 4
}

// Circle
type Circle struct {
	radius float64
}

// Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// インターフェースを作ることで
// この関数にShapeを継承する全ての構造体のオブジェクトがアクセスできる
func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
	fmt.Println()
}

type Person struct {
	Name, Country string
}

// カスタムでStringメソッドを作れる
func (p Person) String() string {
	// Sprintfはstring型とそれ以外の型（今回の場合Person型）を
	// 一緒にstring型にして出力してくれるもの
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}

// 空の構造体
type customWriter struct{}

type GitHubResponse []struct {
	FullName string `json:"full_name"`
}

// p：APIからとってきたJSONデータからfull_nameのところだけ取得して
func (w customWriter) Write(p []byte) (n int, err error) {
	var resp GitHubResponse
	json.Unmarshal(p, &resp)
	// rangeはsliceとかを回せるやつ
	// 最初の変数の方にはindexが取れるが今回は不要
	for _, r := range resp {
		fmt.Println(r.FullName)
	}
	return len(p), nil
}

func main() {
	var s Shape = Square{3}
	printInformation(s)

	c := Circle{6} // Shapeオブジェクトと指定してなくてもいける
	printInformation(c)

	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)

	// Github APIを使用してMicrosoftから3つリポジトリを取得する
	resp, err := http.Get("https://api.github.com/users/microsoft/repos?page=15&per_page=5")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	writer := customWriter{}
	// 最初の引数にWriterがあるioの中身を見るとWriterインターフェースがある
	io.Copy(writer, resp.Body)
}
