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

func countAnagramInString(str string, ptrString string) {

	k := len(ptrString)
	i := 0
	ans := []int{}

	ptrMap := make(map[rune]int)

	for _, val := range ptrString {
		ptrMap[val]++
	}

	count := len(ptrMap)

	for j, v := range str {
		winSize := j - i + 1
		fmt.Printf("window %v\n", i)
		ptrMap[v]--
		if ptrMap[v] == 0 {
			count--
		}

		if winSize == k {
			if count == 0 {
				ans = append(ans, i)
			}

			ptrMap[v]++
			if ptrMap[v] > 0 {
				count++
			}
			i++
		}
	}

	fmt.Printf("Number of anagrams %v", ans)

}

func maxOfAllSubArray(arr []int, k int) {

	var tempArr, maxOfAll []int

	i := 0
	for j := 0; j < len(arr); j++ {
		winSiz := j - i + 1

		if len(tempArr) != 0 {
			for n := len(tempArr) - 1; n >= 0; n-- {
				if arr[j] < tempArr[n] {
					break
				}
				tempArr = tempArr[0:n]
			}

		}
		tempArr = append(tempArr, arr[j])

		if winSiz == k {

			maxOfAll = append(maxOfAll, tempArr[0])

			if arr[i] == tempArr[0] {
				tempArr = tempArr[1:]
			}
			i++

		}
	}

	fmt.Printf("max of all %v  \n", maxOfAll)
}
