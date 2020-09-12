package main

import (
	"fmt"
	"sort"
)

// FindTriplets returns all the triplets from an array such that the sum of the first two elements is equal to the third element
func FindTriplets(input []int) {
	type Triplets struct {
		First, Second, Third int
	}
	result := []Triplets{}
	resultMap := map[Triplets]bool{}
	inputArr := make([]int, len(input))
	copy(inputArr, input)
	for i, v := range inputArr {
		if i == len(inputArr)-1 {
			break
		}
		firstEle := v
		secondEle := inputArr[i+1]
		var diff int
		if firstEle >= secondEle {
			diff = firstEle - secondEle
		} else {
			diff = secondEle - firstEle
		}
		for _, thirdEle := range inputArr {
			if thirdEle == diff {
				set := []int{firstEle, secondEle, thirdEle}
				sort.Ints(set)
				trip := Triplets{set[0], set[1], set[2]}
				if val := resultMap[trip]; val == false {
					resultMap[trip] = true
				}
			}
		}
	}
	for k := range resultMap {
		result = append(result, k)
	}
	fmt.Println(result)
}
func main() {
	FindTriplets([]int{5, 32, 1, 7, 10, 50, 19, 21, 2})
}
