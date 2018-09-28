// func intersection(nums1 []int, nums2 []int) []int {
//     ans := []int{}
//     d1 := make(map[int]bool)
//     for _, v := range nums1 {
//         if _, ok := d1[v]; !ok {
//             d1[v] = true
//         }
//     }
    
//     d2 := make(map[int]bool)
//     for _, v := range nums2 {
//         if _, ok := d2[v]; !ok {
//             d2[v] = true
//         }
//     }
    
//     for k, _ := range d1 {
//         if _, ok := d2[k]; ok {
//             ans = append(ans, k)
//         }
//     }
//     return ans
// }

func intersection(nums1 []int, nums2 []int) []int {
    ans := []int{}
    d1 := make(map[int]bool)
    n1 := []int{}
    for _, v := range nums1 {
        if _, ok := d1[v]; !ok {
            d1[v] = true
            n1 = append(n1, v)
        }
    }
    
    d2 := make(map[int]bool)
    n2 := []int{}
    for _, v := range nums2 {
        if _, ok := d2[v]; !ok {
            d2[v] = true
            n2 = append(n2, v)
        }
    }
    
    sort.Ints(n1)
    sort.Ints(n2)
    i, j := 0, 0
    for i < len(n1) && j < len(n2) {
        if n1[i] < n2[j] {
            i++
        } else if n1[i] > n2[j] {
            j++
        } else {
            ans = append(ans, n1[i])
            i++
            j++
        }
    }
    return ans
}

