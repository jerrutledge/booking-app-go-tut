package main

import (
	"fmt"
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	total_count := len(nums1) + len(nums2)
	digit := 0
	for i, j, count := 0, 0, 0; i < len(nums1) && j < len(nums2); count++ {
		if nums1[i] <= nums2[j] {
			digit = nums1[i]
			i++
		} else {
			digit = nums2[j]
			j++
		}
		if total_count%2 == 0 {
			if count == total_count/2 {
				// average of this and the next number
			}
		} else if count == int(math.Ceil(float64(total_count/2))) {
			// odd number, so this is the one
		}
	}
	return float64(digit)
}

// Where to start the program - compiler signal
func main() {
	nums1 := [2]int{1, 3}
	nums2 := [1]int{2}
	fmt.Print(findMedianSortedArrays(nums1[:], nums2[:]))
}
