package main

import "math"

type FlexWindow struct {
	arrayOfInputs []int
}

func (fw *FlexWindow) Find(k int) ([]int, int) {
	queue := NewQueue(0, true)
	subArray := make([]int, 0)
	endIndex := 0
	sumOfSubArray := 0
	minSumOfSubArray := math.MaxInt
	for endIndex <= len(fw.arrayOfInputs) {
		if sumOfSubArray >= k {
			if sumOfSubArray <= minSumOfSubArray {
				subArray = queue.queue
				minSumOfSubArray = sumOfSubArray
			}

			startItem, _ := queue.Pop()
			sumOfSubArray -= startItem
		} else if endIndex >= len(fw.arrayOfInputs) {
			break
		} else {
			sumOfSubArray += fw.arrayOfInputs[endIndex]
			queue.Push(fw.arrayOfInputs[endIndex], false)
			endIndex += 1
		}
	}
	return subArray, sumOfSubArray
}
