package main

import (
	"fmt"
)

func main() {
	// map；C#で言うDictionary
	// stringの方がkeyで、intの方がvalueになっている↓
	// マップは動的：追加・削除ができる
	studentsAge := map[string]int{
		"john": 32,
		"bob":  31,
	}
	fmt.Println(studentsAge)

	// 初期化しないとき
	studentsAge1 := make(map[string]int)
	fmt.Println(studentsAge1)

	// 項目の追加
	studentsAge2 := make(map[string]int)
	studentsAge["john"] = 32 // 追加
	studentsAge["bob"] = 31  // 追加
	fmt.Println(studentsAge2)

	// 以下で宣言するとnilが入っているが、この状態で追加しようとするとエラーになる
	// var studentsAge map[string]int
	// make関数を使用した場合は空になる
	// nilマップで検索、削除、またはループの操作を実行した場合、Go がパニックに陥ることはない

	// 項目へのアクセス
	studentsAge3 := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	// mapfunc[key]でkeyに対するvalueを取得できる
	fmt.Println("Bob's age is", studentsAge3["bob"])

	// 存在しないkeyを使ってアクセスしてもパニックにはならない
	// ただし既定値が返される。例えば今回の場合int型なのでint型の規定値である0が返される

	// ↑の場合keyが存在するかどうか確かめる必要がある
	// ageにvalueが、existにあboolが戻り値として返ってくる
	age, exist := studentsAge3["christy"]
	if exist {
		fmt.Println("Christy's age is", age)
	} else {
		fmt.Println("Christy's age couldn't be found")
	}

	// 項目の削除
	delete(studentsAge3, "john")
	fmt.Println(studentsAge3)

	// マップのループ
	studentsAge4 := make(map[string]int)
	studentsAge["john"] = 32
	studentsAge["bob"] = 31
	// rangeをつけることで、nameにkeyが、ageにvalueが入ってくる
	for name, age := range studentsAge4 {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// 「_」を使うと無視できる
	for _, age := range studentsAge4 {
		fmt.Printf("Ages %d\n", age)
	}

	// keyのみ取得することもできる
	for name := range studentsAge {
		fmt.Printf("Names %s\n", name)
	}
}
