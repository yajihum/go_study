package main

import (
	"fmt"
)

func getFibonacciSequence(val int) []int {
	var nums []int
	if val < 2 {
		return nums // 何も入っていない状態であればnilを暗黙的に返せる
	}

	for i := 0; i < val; i++ {
		if i <= 1 {
			nums = append(nums, 1)
			continue
		}
		nums = append(nums, nums[i-1]+nums[i-2])
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
