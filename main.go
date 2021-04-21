package main

import (
	"fmt"
	"leetcode-practice/array"
)

func main() {
	var nums []int = []int{
		1, 2, 3, 4,
	}
	fmt.Print(array.TwoSum(nums, 3))
}
