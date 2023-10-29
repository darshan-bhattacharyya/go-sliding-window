package main

import "fmt"

func main() {
	arrayOfInputs := []int{5, 2, 8, 9, 1, 2, 3, 2, 4}

	suite := FixedWindow{
		arrayOfInputs:  arrayOfInputs,
		sizeOfSubArray: 3,
	}

	subArray, subArraySum := suite.Find(6)

	fmt.Println(subArray)
	fmt.Println(subArraySum)
	fmt.Println("------ flex window-------")

	flexWindow := FlexWindow{
		arrayOfInputs: arrayOfInputs,
	}
	subArray, subArraySum = flexWindow.Find(5)
	fmt.Println(subArray)
	fmt.Println(subArraySum)
}
