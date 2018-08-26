func generate(numRows int) [][]int {
    ans := [][]int{}
    if numRows == 0 {
        return ans
    }
    
    tmp := []int{1}
    ans = append(ans, tmp)
    for numRows != 1 {
        tmp = gene_row(tmp)
        ans = append(ans, tmp)
        numRows--
    }
    return ans
}

func gene_row(row []int) []int {
    next_row := []int{1}
    
    for i := 0; i < len(row)-1; i++ {
        next_row = append(next_row, row[i] + row[i+1])
    }
    
    next_row = append(next_row, 1)
    
    return next_row
}
