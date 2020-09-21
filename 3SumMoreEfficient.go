package main

import (
	"fmt"
	"sort"
)

type number struct {
	value      int
	occurrence int
}

func (n *number) incrementOccurrence(val int) (newOccurrence int) {
	n.occurrence += val
	newOccurrence = n.occurrence
	return
}
func count(inputSlice []int, element int) (c int) {
	c = 0
	for _, el := range inputSlice {
		if el == element {
			c++
		}
	}
	return
}
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	resultMap := map[[3]int]bool{}
	numMap := make(map[int]*number)
	resultList := make([][]int, 0)
	//Populate numMap
	for _, num := range nums {
		if _, okay := numMap[num]; okay {
			numMap[num].incrementOccurrence(1)
		} else {
			n := number{num, 1}
			numMap[num] = &n
		}
	}
	for i, num1 := range nums {
		for j, num2 := range nums {
			if i == j {
				continue
			} else {
				twoSum := num1 + num2
				searchKey := -twoSum
				if num3, ok := numMap[searchKey]; ok {
					list := []int{num1, num2, num3.value}
					sameNumReused := false
					for _, item := range list {
						if c := count(list, item); c > numMap[item].occurrence {
							sameNumReused = true
							break
						}
					}
					if !sameNumReused {
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
	}
	for k := range resultMap {
		key := make([]int, len(k))
		copy(key, k[:])
		resultList = append(resultList, key)
	}
	return resultList
}
func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}))
}
