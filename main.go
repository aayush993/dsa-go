package main

func main() {
	// Given an array
	// find max sum of subarray of size k

	arr := []int{1, 4, 2, 10, 2, 3, 1, 0, 20}
	k := 4
	MaxSum(arr, k)

	arr = []int{100, 200, 300, 400}
	k = 2
	MaxSum(arr, k)
	arr = []int{100, 200, 300, 400}
	k = 4
	MaxSum(arr, k)

}
