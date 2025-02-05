package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// Ensure nums1 is the shorter array
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	left := 0
	half := (len(nums1) + len(nums2) + 1) / 2
	right := len(nums1)

	for left <= right {
		x := (left + right) / 2
		y := half - x

		if x > 0 && nums1[x-1] > nums2[y] {
			right = x - 1
		} else if x < len(nums1) && nums2[y-1] > nums1[x] {
			left = x + 1
		} else {
			var maxLeft float64
			if x == 0 {
				maxLeft = float64(nums2[y-1])
			} else if y == 0 {
				maxLeft = float64(nums1[x-1])
			} else {
				maxLeft = math.Max(float64(nums1[x-1]), float64(nums2[y-1]))
			}

			// if total length is odd, maxLeft is the median
			if (len(nums1)+len(nums2))%2 != 0 {
				return maxLeft
			}

			var minRight float64
			if x == len(nums1) {
				minRight = float64(nums2[y])
			} else if y == len(nums2) {
				minRight = float64(nums1[x])
			} else {
				minRight = math.Min(float64(nums1[x]), float64(nums2[y]))
			}

			return (maxLeft + minRight) / 2.0
		}
	}

	return -1.0
}

func main() {
	println("4. Median of Two Sorted Arrays")

	tc1 := findMedianSortedArrays([]int{1, 2}, []int{2})
	println(tc1) // Expected: 2

	tc2 := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	println(tc2) // Expected: 2.5

	tc3 := findMedianSortedArrays([]int{1, 7}, []int{3, 10, 12})
	println(tc3) // Expected: 7

	tc4 := findMedianSortedArrays([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5, 6, 7, 8})
	println(tc4) // Expected 4
}
