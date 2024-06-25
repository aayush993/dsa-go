package main

import "fmt"

func MaxSum(arr []int, k int) {

	var sum, mSum int
	i := 0

	for j := 0; j < len(arr); j++ {
		sum = sum + arr[j]
		mSum = sum

		if j-i+1 == k {
			sum = sum - arr[i]
			if sum > mSum {
				mSum = sum
			}
			i++
		}
	}

	fmt.Printf("Max sum is %v\n", mSum)
}

func FirstNegativeInt(arr []int, k int) {

	var result int
	i := 0
	var negativeList []int

	for j := 0; j < len(arr); j++ {
		winSize := j - i + 1
		if arr[j] < 0 {
			negativeList = append(negativeList, arr[j])
		}

		if winSize == k {
			if len(negativeList) == 0 {
				result = 0
			} else {
				result = negativeList[0]

				if arr[i] == negativeList[0] {
					negativeList = negativeList[1:] //bug
				}
			}
			fmt.Printf("First negative int is %d\n", result)
			i++
		}
	}

}
