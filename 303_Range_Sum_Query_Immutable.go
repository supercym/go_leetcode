type NumArray struct {
	mynums []int
	d      map[int]int
}

func Constructor(nums []int) NumArray {
	numsarray := &NumArray{}
	numsarray.mynums = append(numsarray.mynums, nums...)
	if len(nums) > 0 {
		numsarray.d = make(map[int]int)
		numsarray.d[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			numsarray.d[i] = nums[i] + numsarray.d[i-1]
		}
	}
	return *numsarray
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.d[j]
	}
	return this.d[j] - this.d[i-1]
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
 