func fourSumCount(A []int, B []int, C []int, D []int) int {
    n := len(A)
    d1 := make(map[int]int, 0)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if _, ok := d1[A[i]+B[j]]; ok {
                d1[A[i]+B[j]]++
            } else {
                d1[A[i]+B[j]] = 1
            }
        }
    }
    
    ans := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {  
            s := C[i]+D[j]
            if _, ok := d1[-s]; ok {
                ans += d1[-s]
            } 
        }
    }
    
    return ans
}