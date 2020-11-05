package hard

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

Follow up: The overall run time complexity should be O(log (m+n)).
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var index1, index2 int
	var len1, len2 = len(nums1), len(nums2)
	var current float64
	if len1+len2 == 1 {
		nums1 = append(nums1, nums2...)
		return float64(nums1[0])
	}
	if len1+len2 == 2 {
		nums1 = append(nums1, nums2...)
		return float64(nums1[0]+nums1[1]) / 2
	}

	if (len1+len2)%2 == 0 {
		// we are looking for 2 numbers
		seedNo := (len1 + len2) / 2
		for index1 < len1 || index2 < len2 {
			n1, n2 := 1000001, 1000001
			if len1 != index1 {
				n1 = nums1[index1]
			}
			if len2 != index2 {
				n2 = nums2[index2]
			}
			if n1 < n2 {
				current = float64(nums1[index1])
				// increase index
				index1++
				seedNo--
			} else {
				current = float64(nums2[index2])
				// increase index
				index2++
				seedNo--
			}
			if seedNo == 0 {
				last1, last2 := 1001, 1001
				if len1 != index1 {
					last1 = nums1[index1]
				}
				if len2 != index2 {
					last2 = nums2[index2]
				}
				if last1 < last2 {
					current += float64(nums1[index1])
				} else {
					current += float64(nums2[index2])
				}
				break
			}
		}
		current /= 2
	} else {
		seedNo := (len1+len2)/2 + 1
		for index1 < len1 || index2 < len2 {
			n1, n2 := 1001, 1001
			if len1 != index1 {
				n1 = nums1[index1]
			}
			if len2 != index2 {
				n2 = nums2[index2]
			}
			if n1 < n2 {
				current = float64(nums1[index1])
				// increase index
				index1++
				seedNo--
			} else {
				current = float64(nums2[index2])
				// increase index
				index2++
				seedNo--
			}
			if seedNo == 0 {
				break
			}
		}
	}

	return current
}
