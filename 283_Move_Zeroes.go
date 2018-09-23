func moveZeroes(nums []int)  {
    i := 0
    start := 0
    move := 0
    sum_move := 0
    length := len(nums)
    for {
        if nums[i] == 0 {
            move++
            i++
        } else {
            if move != 0 {
                nums = append(nums[:start], nums[start + move:]...)
                for j := 0; j < move; j++ {
                    nums = append(nums, 0)
                }
                i -= move
                sum_move += move
                move = 0
            } else {
                start = i + 1
                i++
            }
        }
        if i + sum_move >= length {
            break
        }
    }
}

