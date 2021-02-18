package main

import "fmt"

func main() {
  // Let's assume that we have a slice of integers (numbers may repeat) then we need to sort the integers based on their counts
  // with a no with more occurance coming before a no with lesser occurance. 
  // If two integers have the same count then the greater no will be placed before the smaller one (in terms of their values). 
  // The sorted output should be repetition free.

  li := []int{4,1,1,2,2,2,8,9,16,16,5,5,5,5,11}
  sortBasedOnCount(li)
}

func sortBasedOnCount(li []int) []int{
  liCopy := make([]int, len(li))
  nonRepeatedLi := make([]int, 0)
  copy(liCopy, li)
  counts := make(map[int]int)
  for _, num := range li{
    if _,present := counts[num]; present{
      counts[num] += 1
    }else{
      counts[num] = 1
      nonRepeatedLi = append(nonRepeatedLi, num)
    }
  }
  for i,_ := range nonRepeatedLi{
    for j:=i+1; j<len(nonRepeatedLi); j++{
      if counts[nonRepeatedLi[i]]<counts[nonRepeatedLi[j]]{
        //Swap
        var temp int
        temp = nonRepeatedLi[i]
        nonRepeatedLi[i] = nonRepeatedLi[j]
        nonRepeatedLi[j] = temp
      }else if counts[nonRepeatedLi[i]] == counts[nonRepeatedLi[j]]{
        if nonRepeatedLi[i] < nonRepeatedLi[j]{
          //Swap
          var temp int
          temp = nonRepeatedLi[i]
          nonRepeatedLi[i] = nonRepeatedLi[j]
          nonRepeatedLi[j] = temp
        }
      }
    }
  } 
  fmt.Println(counts)
  fmt.Println(nonRepeatedLi)
  return []int{}
}
