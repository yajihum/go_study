package main

import (
	"fmt"
	"math/rand"
	"regexp"
	"time"
)

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

	// 条件のところで関数も使えちゃう
	switch time.Now().Weekday().String() {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("It's time to learn some Go.")
	default:
		fmt.Println("It's the weekend, time to rest!")
	}

	fmt.Println(time.Now().Weekday().String())

	// 関数を呼び出すことも可能
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)
	var phone = regexp.MustCompile(`^[(]?[0-9][0-9][0-9][). \-]*[0-9][0-9][0-9][.\-]?[0-9][0-9][0-9][0-9]`)

	contact := "foo@bar.com"

	switch {
	case email.MatchString(contact):
		fmt.Println(contact, "is an email")
	case phone.MatchString(contact):
		fmt.Println(contact, "is a phone number")
	default:
		fmt.Println(contact, "is not recognized")
	}

	// ifで書くよりもswitchがいい場合もある
	// 以下では、if elseで書くこともできるが、switchで書く方が短くかける？
	rand.Seed(time.Now().Unix())
	r := rand.Float64()
	switch {
	case r > 0.1:
		fmt.Println("Common case, 90% of the time")
	default:
		fmt.Println("10% of the time")
	}

	// Goではbreakはいらない、1つのcaseに入るとswitch文は終了する
	// 次のcaseにフォールスルーされるようにしたい場合は「fallthrough」を使う
	// 以下の場合、最初のcaseと一致するがfallthroughがあるため2,3番目のcaseも実行されてしまう
	switch num := 15; {
	case num < 50:
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num > 100:
		fmt.Printf("%d is greater than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200", num)
	}
}
