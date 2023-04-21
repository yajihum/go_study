package main

import (
	"fmt"
)

func main() {
	// sliceとは：同じ型の要素が連続していることを表す Go のデータ型
	// 配列との違い：スライスのサイズは固定ではなく動的
	// sliceの初期化：要素数を指定していないことに注意、これが元になる配列
	months := []string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	fmt.Println(months)
	fmt.Println("Length:", len(months))   // Length: 12
	fmt.Println("Capacity:", cap(months)) // Capacity: 12

	// s[i:p]という形でsliceできる（s：元になる配列、i：新しい配列をスライスする最初のindex、p：新しい配列をスライスする最後のindex）
	quarter1 := months[0:3]
	quarter2 := months[3:6]
	quarter3 := months[6:9]
	quarter4 := months[9:12]
	// スライスで指定する最初のindexから数えてどれくらいの要素を指定できるかがcapの出力数になる
	fmt.Println(quarter1, len(quarter1), cap(quarter1)) // [January February March] 3 12
	fmt.Println(quarter2, len(quarter2), cap(quarter2)) // [April May June] 3 9
	fmt.Println(quarter3, len(quarter3), cap(quarter3)) // [July August September] 3 6
	fmt.Println(quarter4, len(quarter4), cap(quarter4)) // [October November December] 3 3

	// スライスした配列を拡張できる
	// quarter2の4番目の要素から4つ分を取得
	quarter2Extended := quarter2[:4]                                            // 逆に最後の位置を省略すると最後までという意味になる
	fmt.Println(quarter2Extended, len(quarter2Extended), cap(quarter2Extended)) // [April May June July] 4 9

	// 要素を追加する：appendを使用
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("%d\tcap=%d\t%v\n", i, cap(numbers), numbers)
	}
	// cap()の出力は順にcap=1,2,4,4,8,8,8,8,16,16になる
	// これはcapを超えた要素数になってしまった場合その時点でのcapの2倍をcapに設定するようになるため
	// ただし、ある時点でスライスの容量が必要以上に多くなり、メモリが無駄に消費されるおそれがある

	// 要素の削除
	// 削除できる関数は用意されていないため、新しいスライスを作成する形で対応
	letters := []string{"A", "B", "C", "D", "E"}
	remove := 2
	if remove < len(letters) {
		fmt.Println("Before", letters, "Remove ", letters[remove])
		// 0~1番目までのスライスに対して3番目以降のスライスを追加している
		letters = append(letters[:remove], letters[remove+1:]...) // 「...」はスライスの要素を個々に渡すことができる
		fmt.Println("After", letters)
	}

	// 削除のもう1つの方法はスライスのコピーを作成すること
	slice2 := make([]string, 3) // make関数で長さ5容量5のint型配列を作成 ちなみにmake([]int, 3, 5)では長さ3容量5の配列ができる
	letters1 := []string{"A", "B", "C", "D", "E"}
	// jsと同じで3章型になるため同じ配列を使っている方に影響を与えないためにもコピーを作成する
	copy(slice2, letters1[1:4])

	fmt.Println("Before", letters1)

	slice1 := letters1[0:2]
	copy(slice2, letters1[1:4])

	slice1[1] = "Z"

	// slice1のZになった影響をslice2は受けていない
	fmt.Println("After", letters1) // After [A Z C D E]
	fmt.Println("Slice2", slice2)  // Slice2 [B C D]
}
