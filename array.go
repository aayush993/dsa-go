package main

import "fmt"

func MaxSum(arr []int, k int) {

	var sum, mSum int
	i := 0

	for j := 0; j < len(arr); j++ {
		//fmt.Printf("Sum is  %v\n", sum)
		if j-i+1 <= k {
			sum = sum + arr[j]
			mSum = sum
		} else if j-i+1 > k {
			sum := sum - arr[i] + arr[j]
			if sum > mSum {
				mSum = sum
			}
			i++
		}
	}

	fmt.Printf("Max sum is %v\n", mSum)
}
