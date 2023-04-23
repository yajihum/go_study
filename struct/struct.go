package main

import "fmt"

// 構造体
type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

// ↑のperson型を再利用
type Employee1 struct {
	Information Person
	ManagerID   int
}

func main() {
	// 変数の生成
	//var john Employee

	// 変数の宣言と初期化
	employee := Employee{1001, "John", "Doe", "Doe's Street"}
	fmt.Println(employee)

	// フィールドを明示的に指定する場合は他のフィールドは絞り込んでおk
	employee1 := Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee1)

	// 各フィールドへのアクセスは「.」を使用する
	employee.ID = 1001

	employeeCopy := &employee1 // ポインターを生成
	employeeCopy.FirstName = "David"
	fmt.Println(employee1)

	// 構造体を再利用することができる
	var employee2 Employee1
	employee2.Information.FirstName = "John"
}
