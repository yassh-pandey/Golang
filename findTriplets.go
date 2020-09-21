package main

import (
	"fmt"
	"sort"
)

type Triplets struct {
	First, Second, Third int
}

// FindTriplets returns all the triplets from an array of positive integers such that the sum of the first two elements is equal to the third element
func FindTriplets(input []int) []Triplets {
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
	return result
}
func main() {
	answer := FindTriplets([]int{5, 32, 1, 7, 10, 50, 39, 21, 2})
	fmt.Println(answer)
}
