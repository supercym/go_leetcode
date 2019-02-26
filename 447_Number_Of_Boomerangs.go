func numberOfBoomerangs(points [][]int) int {
    ans := 0
    for i := 0; i < len(points); i++ {
        dis := make(map[int]int, 0)
        for j := 0; j < len(points); j++ {
            x2 := (points[i][0]-points[j][0])*(points[i][0]-points[j][0])
            y2 := (points[i][1]-points[j][1])*(points[i][1]-points[j][1])
            d := x2+y2
            if _, ok := dis[d]; ok {
                dis[d]++
            } else {
                dis[d] = 1
            }
        }
        for _, v := range dis {
            ans += v*(v-1)
        }
    }
    
    return ans
}