package main

import "math"

type FixedWindow struct {
	arrayOfInputs  []int
	sizeOfSubArray int
}

func (s *FixedWindow) Find(k int) ([]int, int) {
	subArray := make([]int, s.sizeOfSubArray)
	start := 0
	end := s.sizeOfSubArray - 1
	sumOfSubArray := sum(s.arrayOfInputs[start : end+1])
	minMax := math.MaxInt

	for end < len(s.arrayOfInputs) && start < end {
		if sumOfSubArray >= k && sumOfSubArray < minMax {
			minMax = sumOfSubArray
			subArray = s.arrayOfInputs[start : end+1]
		}
		if end == len(s.arrayOfInputs)-1 {
			break
		}
		sumOfSubArray -= s.arrayOfInputs[start]
		sumOfSubArray += s.arrayOfInputs[end+1]
		start += 1
		end += 1
	}

	return subArray, minMax
}

func sum(array []int) int {
	sum := 0
	for _, i := range array {
		sum += i
	}
	return sum
}
