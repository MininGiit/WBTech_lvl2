package main

import (
	"fmt"
	"strategy/types"
)

func main() {
    arr := []int{5, 2, 8, 1, 9, 4}

    bubbleSorter := types.NewSorter(&types.QickSort{})
    sortedArr1 := bubbleSorter.Sort(arr)
    fmt.Println("Qick Sort:", sortedArr1) 

    mergeSorter := types.NewSorter(&types.MergeSort{})
    sortedArr2 := mergeSorter.Sort(arr)
    fmt.Println("Merge Sort:", sortedArr2)
}