package main

import (
	"fmt"
	"math"
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

func main() {
	var s Shape = Square{3}
	printInformation(s)

	c := Circle{6} // Shapeオブジェクトと指定してなくてもいける
	printInformation(c)
}
