package main

import (
	"fmt"
	"go_study/geometry"
	"strings"
)

type triangle struct {
	size int
}

type square struct {
	size int
}

/*
triangleの構造体のみこのメソッドが利用可能になる
引数に構造体を指定
s.perimeterでアクセス可能
*/
func (t triangle) perimeter() int {
	return t.size * 3
}

/*
ラッパーメソッド
引数が違うことによって元のperimeterメソッドをオーバーロードできる
*/
func (t coloredTriangle) perimeter() int {
	return t.triangle.perimeter()
}

/*
squareの構造体のみこのメソッドが利用可能になる
引数に構造体を指定
s.perimeterでアクセス可能
*/
func (s square) perimeter() int {
	return s.size * 4
}

/*
変数を直接更新する場合は前と同様にポインターを使用する
Go の規則では、構造体のいずれかのメソッドにポインターレシーバーがある場合、
その構造体のすべてのメソッドにポインター レシーバーが必要
*/
func (t *triangle) doubleSize() {
	t.size *= 2
}

/*
stringなど基本型に対してメソッドを作成したい場合は
カスタム型にする必要がある
*/
type upperstring string

/*
文字列を小文字から大文字にするメソッド
*/
func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

/*
三角形に色を持たせた構造体
*/
type coloredTriangle struct {
	triangle
	color string
}

func main() {
	t := triangle{3}
	s := square{4}
	t.doubleSize()
	// コンパイラではレシーバーの型に基づいて呼び出す関数が決定
	fmt.Println("Perimeter (triangle):", t.perimeter())
	fmt.Println("Perimeter (square):", s.perimeter())

	str := upperstring("Learning Go!")
	fmt.Println(str)
	fmt.Println(str.Upper())

	tr := coloredTriangle{triangle{3}, "blue"}
	fmt.Println("Size:", tr.size)
	// perimeterメソッドにもアクセスできる
	fmt.Println("Perimeter", tr.perimeter())
	// triangle構造体の方のperimeterメソッドに明示的にアクセス
	fmt.Println("Perimeter (normal)", tr.triangle.perimeter())

	tri := geometry.Triangle{}
	tri.SetSize(3)
	fmt.Println("Size", tri.size) //カプセル化によってアクセスできないためエラー
	fmt.Println("Perimeter", tri.Perimeter())
}
