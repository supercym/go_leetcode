func groupAnagrams(strs []string) [][]string {
    d := make(map[string][]string, 0)
    for _, s := range strs {
        // sort the string
        sl := make([]string, 0)
        for i := 0; i < len(s); i++ {
            sl = append(sl, string(s[i:i+1]))
        }
        sort.Strings(sl)
        ss := strings.Join(sl, "")
        
        if _, ok := d[ss]; ok {
            d[ss] = append(d[ss], s)
        } else {
            d[ss] = []string{s}
        }
    }
    
    ans := make([][]string, 0)
    for _, v := range d {
        ans = append(ans, v)
    }
    return ans
}