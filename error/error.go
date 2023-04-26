package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func main() {
	// エラーを判別
	employee, err := getInformation(1001)
	if err != nil {
		// Something is wrong. Do something.
	} else {
		fmt.Print(employee)
	}
}

// エラーインスタンスを作成できる
var ErrNotFound = errors.New("Employee not found!")

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(1000)
	// errorを受け取ってそれがnilかどうかを判別する
	if err != nil {
		// Errorfメソッドが利用できる
		//return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
		return nil, ErrNotFound // エラー変数は最初にErrをつける必要がある
	}

	// errors.Is関数を使うと発生しているエラーの種類を比較できる
	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT FOUND: %v\n", err)
	} else {
		fmt.Print(employee)
	}
	return employee, nil
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
