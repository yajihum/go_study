package main

import (
	"fmt"
	"math/rand"
	"time"
)

func givemeanumber() int {
	return -1
}

/*
switch文の条件式に複数を設定
*/
func location(city string) (string, string) {
	var region string
	var continent string
	switch city {
	case "Delhi", "Hyderabad", "Mumbai", "Chennai", "Kochi":
		region, continent = "India", "Asia"
	case "Lafayette", "Louisville", "Boulder":
		region, continent = "Colorado", "USA"
	case "Irvine", "Los Angeles", "San Diego":
		region, continent = "California", "USA"
	default:
		region, continent = "Unknown", "Unknown"
	}
	return region, continent
}

func main() {
	x := 27

	// ifの使い方
	// 条件式部分に括弧は必要ない
	// 三項演算子は使えない
	if x%2 == 0 {
		fmt.Println(x, "is even")
	}

	// Goではifのブロック内で変数を宣言するのが一般的
	// 以下ではnumに関数の戻り値を代入し、そのnum変数を{}内で使用している
	// このnuはifブロックの外側では使えない
	if num := givemeanumber(); num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has only one digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// rand.Seed()によって乱数の初期値（シード）を設定できるようにしている：シード値を設定することで乱数生成の開始点が決まる
	// 以下の例では、現在の時刻ををシード値として設定し、0~9までのランダムな整数を生成している
	sec := time.Now().Unix()
	rand.Seed(sec)
	i := rand.Int31n(10)

	// switch文
	switch i {
	case 0:
		fmt.Print("zero...")
	case 1:
		fmt.Print("one...")
	case 2:
		fmt.Print("two...")
	default:
		fmt.Print("no match...")
	}

	// switch文の条件に複数を設定
	// 最初の%sで一番めの引数のregionを、次の%sで二番目のcontentを出力
	// \nは文字列の末尾に改行を追加する
	region, continent := location("Irvine")
	fmt.Printf("John works in %s, %s\n", region, continent)
}
