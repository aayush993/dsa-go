package main

import (
	"fmt"
	"math/rand"
)

func selectionSort(arr []int) {

	n := len(arr)
	for i := 0; i < n-1; i++ {
		// lets test with n-1 first
		// it only goes up until n-2

		iMin := i

		for j := i + 1; j < n-1; j++ {
			if arr[j] < arr[iMin] {
				iMin = j
			}
		}

		// cpu uses registers to keep track of the value while swapping
		arr[iMin], arr[i] = arr[i], arr[iMin]
	}

	fmt.Printf("sorted arr %v\n", arr)
}

func bubbleSort(arr []int) {
	n := len(arr)

	isSwapPass := false
	for i := 1; i < n-1; i++ {

		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				isSwapPass = true
			}
		}

		if !isSwapPass {
			break
		}
	}

	fmt.Printf("sorted arr %v\n", arr)
}

func insertionSort(arr []int) {

	n := len(arr)

	for i := 1; i < n-1; i++ {

		val := arr[i]
		hole := i
		for hole > 0 && val < arr[hole-1] {
			arr[hole] = arr[hole-1] // cannot do swap here
			hole--
		}
		arr[hole] = val

	}

	fmt.Printf("sorted arr %v\n", arr)
}

func mergeSort(arr []int) []int {

	var k int
	if len(arr) < 2 {
		return arr
	} else {
		n := len(arr)
		k = n / 2

	}

	lArr := mergeSort(arr[:k])
	rArr := mergeSort(arr[k:])
	arr = mergeArr(lArr, rArr)

	return arr

}

func mergeArr(lArr []int, rArr []int) []int {
	i := 0
	j := 0
	k := 0
	nLeft := len(lArr)
	nRight := len(rArr)

	retArr := make([]int, nLeft+nRight)

	for i < nLeft && j < nRight {

		if lArr[i] <= rArr[j] {
			retArr[k] = lArr[i]
			i++
		} else {
			retArr[k] = rArr[j]
			j++
		}

		k++
	}

	for i < nLeft {
		retArr[k] = lArr[i]
		k++
		i++
	}
	for j < nRight {
		retArr[k] = rArr[j]
		k++
		j++
	}

	return retArr

}

func quickSort(arr []int, start int, end int) {

	if start < end {
		pIndex := randomPartition(arr, start, end)

		quickSort(arr, start, pIndex-1)
		quickSort(arr, pIndex+1, end)
	}

}

func partition(arr []int, start, end int) int {

	pIndex := start // start from zero and increase as we populate left most list
	// consider this pivot arr[end]
	for i := start; i < end; i++ {
		if arr[i] < arr[end] {
			arr[pIndex], arr[i] = arr[i], arr[pIndex]
			pIndex++ // make sure to increase this

		}

	}

	// our pIndex will be somewhere in between
	// last thing is we swap our pivot at pIndex
	arr[pIndex], arr[end] = arr[end], arr[pIndex]

	return pIndex
}

func randomPartition(arr []int, start, end int) int {

	// Generate a random number in the range [min, max]
	randomNum := rand.Intn(end-start+1) + start
	arr[randomNum], arr[end] = arr[end], arr[randomNum]

	return partition(arr, start, end)
}
