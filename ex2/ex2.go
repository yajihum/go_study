package main

import (
	"fmt"
)

func getFibonacciSequence(val int) []int {
	//var nums []int
	if val < 2 {
		//return nums  何も入っていない状態であればnilを暗黙的に返せる
		return make([]int, 0) // リファクタリング
	}

	// リファクイズ
	// makeでスライスを作成　:=との違いは要素数を指定しなくていい
	nums := make([]int, val)
	nums[0], nums[1] = 1, 1

	// for i := 0; i < val; i++ {
	// 	if i <= 1 {
	// 		nums = append(nums, 1)
	// 		continue
	// 	}
	// 	nums = append(nums, nums[i-1]+nums[i-2])
	// }

	// リファクタリング
	for i := 2; i < val; i++ {
		nums[i] = nums[i-1] + nums[i-2]
	}
	return nums
}

func main() {
	// フィボナッチ数列
	val := 0
	fmt.Print("Enter number: ")
	fmt.Scanf("%d", &val)

	fmt.Print(getFibonacciSequence(val))
}
