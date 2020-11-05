package easy

/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.
*/

/*
Runtime: 16 ms, faster than 95.87% of Go online submissions for Majority Element.
Memory Usage: 5.9 MB, less than 100.00% of Go online submissions for Majority Element.
*/
func majorityElement(nums []int) int {
	mMap := make(map[int]int)
	for _, v := range nums {
		if _, ok := mMap[v]; ok {
			mMap[v] = mMap[v] + 1
			continue
		}
		mMap[v] = 1
	}
	var result, prev int
	for k, v := range mMap {
		if prev < v {
			prev = v
			result = k
		}
	}

	return result
}
