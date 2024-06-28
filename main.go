package main

func main() {
	// Given an array
	// find max sum of subarray of size k

	// arr := []int{1, 4, 2, 10, 2, 3, 1, 0, 20}
	// k := 4
	// MaxSum(arr, k)

	// arr = []int{100, 200, 300, 400}
	// k = 2
	// MaxSum(arr, k)
	// arr = []int{100, 200, 300, 400}
	// k = 4
	// MaxSum(arr, k)

	// arr := []int{12, -1, -7, 8, -15, 30, 16, 28}
	// k := 3

	// FirstNegativeInt(arr, k)

	// str := "aabaabaa"
	// ptr := "aaba"

	// countAnagramInString(str, ptr)

	// arr := "aayuᏝh"
	// for j := 0; j < len(arr); j++ {
	// 	fmt.Println(arr[j])
	// }
	// for k, v := range arr {
	// 	fmt.Printf("key %v and value %v\n", k, v)
	// }

	arr1 := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3

	maxOfAllSubArray(arr1, k)

}
