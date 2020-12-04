package medium

// NumArray : structure of array
type NumArray struct {
	nums []int
}

// Constructor : returns a creation of NumArray
func Constructor(nums []int) NumArray {
	return NumArray{
		nums: nums,
	}
}

// Update : update the number in specific position of NumArray
func (a *NumArray) Update(i int, val int) {
	a.nums[i] = val
}

// SumRange : return the sum of numbers within a given range
func (a *NumArray) SumRange(i int, j int) int {
	total := 0
	for start := i; start <= j; start++ {
		total += a.nums[start]
	}
	return total
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
