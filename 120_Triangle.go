func minimumTotal(triangle [][]int) int {
    h = len(triangle)
    record = make([][]int, h)
    for i = 0; i  h; i++ {
        record[i] = make([]int, i+1)
    }
    
    min = func(a, b int) int {
        if a  b {
            return a
        }
        return b
    }
    
    for i = h-1; i = 0; i-- {
        for j = 0; j  i+1; j++ {
            if i == h-1 {
                record[i][j] = triangle[i][j]
            } else {
                record[i][j] = triangle[i][j] + min(record[i+1][j], record[i+1][j+1])
            }
        }
    } 
    
    return record[0][0]
}