package main

import (
	"fmt"
	"sort"
)

`
	Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? 
	Find all unique triplets in the array which gives the sum of zero.

	Notice that the solution set must not contain duplicate triplets.
`
func threeSum(nums []int) [][]int {
	resultMap := map[[3]int]bool{}
	resultList := make([][]int, 0)
	if len(nums) < 3 {
		return nil
	}
	for i, num1 := range nums {
		for j := i + 1; j < len(nums); j++ {
			num2 := nums[j]
			for k := j + 1; k < len(nums); k++ {
				num3 := nums[k]
				sum := num1 + num2 + num3
				if sum == 0 {
					list := []int{num1, num2, num3}
					sort.Ints(list)
					key := [3]int{}
					copy(key[:], list)
					if _, okay := resultMap[key]; !okay {
						resultMap[key] = true
					}
				}
			}
		}
	}
	for k := range resultMap {
		key := make([]int, len(k))
		copy(key, k[:])
		resultList = append(resultList, key)
	}
	return resultList
}
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
