package main

import (
	"fmt"
	"math"
	"strconv"
)

// RadixSort : this is an implementation of Radix sort that works only for positive integers
func RadixSort(input []int) []int {
	inputArray := make([]int, len(input))
	copy(inputArray, input)
	scanArray := true
	initialCountArray := [10]int{}
	countArray := [10]int{}
	prefixSum := [10]int{}
	scanCounter := 0
	intermediateArray := []int{}
	max := math.Inf(-1)
	for scanArray {
		for i := 0; i < len(inputArray); i++ {
			if scanCounter == 0 {
				if float64(inputArray[i]) > max {
					max = float64(inputArray[i])
				}
				intermediateArray = append(intermediateArray, 0)
			}
			digit := (inputArray[i] / int(math.Pow(10, float64(scanCounter)))) % 10
			countArray[digit]++
		}
		for i := 0; i < len(countArray); i++ {
			if i == 0 {
				prefixSum[i] = countArray[i]
			} else {
				prefixSum[i] = countArray[i] + prefixSum[i-1]
			}
		}
		for i := len(inputArray) - 1; i >= 0; i-- {
			digit := (inputArray[i] / int(math.Pow(10, float64(scanCounter)))) % 10
			intermediateArray[prefixSum[digit]-1] = inputArray[i]
			prefixSum[digit]--
		}
		scanCounter++
		countArray = initialCountArray
		copy(inputArray, intermediateArray)
		if scanCounter == len(strconv.Itoa(int(max))) {
			break
		}
	}
	return inputArray
}
func main() {
	var unsortedArray []int = []int{45, 872, 231, 77, 918, 15, 998, 345, 422, 777}
	sortedArray := RadixSort(unsortedArray)
	fmt.Println(unsortedArray)
	fmt.Println(sortedArray)
}
