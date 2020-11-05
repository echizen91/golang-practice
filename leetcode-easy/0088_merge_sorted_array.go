package easy

/////////// ### KIV
/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is equal to m + n) to hold additional elements from nums2.
*/

// MergeSortedArray : merge nums2 into nums1 as one sorted array
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) []int {
	// start from the end
	var endM, endN, total = m - 1, n - 1, m + n - 1
	for i := endN; i >= 0; i-- {
		// zeros, compare m and n values

		if nums1[endM] <= nums2[endN] {
			nums1[total] = nums2[endN]
			endN--
			total--
		} else {
			nums1[total] = nums1[endM]
			endM--
			total--
			i++
		}

	}
	return nums1
}
