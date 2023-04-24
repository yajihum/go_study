package main

import (
	"encoding/json" // JSONデータをエンコードしたりデコードできるパッケージ
	"fmt"
)

type Person3 struct {
	ID        int
	FirstName string `json:"name"`
	LastName  string
	Address   string `json:"address,omitempty"`
}

type Employee3 struct {
	Person3
	ManagerID int
}

type Contractor struct {
	Person3
	CompanyID int
}

func struct2() {
	employees := []Employee3{
		Employee3{
			Person3: Person3{
				LastName: "Doe", FirstName: "John",
			},
		},
		Employee3{
			Person3: Person3{
				LastName: "Campbell", FirstName: "David",
			},
		},
	}

	data, _ := json.Marshal(employees) // JSON文字列をデータ構造にエンコード　FirstNameはnameとして出力される
	fmt.Printf("%s\n", data)

	var decoded []Employee3
	json.Unmarshal(data, &decoded) // デコード　＆はポインタで&をつけると関数に渡した変数と宣言した変数を同期できる（復習）
	fmt.Printf("%v", decoded)
}
