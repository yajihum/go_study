package main

import "fmt"

type Person2 struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee2 struct {
	Person2   // 直接埋め込むことができる
	ManagerID int
}

type Contractor struct {
	Person2
	CompanyID int
}

func struct1() {
	employee2 := Employee2{
		Person2: Person2{
			FirstName: "John",
		},
	}
	employee2.LastName = "Doe"
	// 直接FirstNameにアクセスできる
	fmt.Println(employee2.FirstName)
}
