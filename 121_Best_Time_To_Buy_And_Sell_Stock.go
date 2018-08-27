func maxProfit(prices []int) int {
    if len(prices) == 0 || len(prices) == 1 {
        return 0
    }
    max_profit := 0
    min_price := prices[0]
    for i := 0; i < len(prices); i++ {
        if prices[i] < min_price {
            min_price = prices[i]
        }
        
        if max_profit < prices[i] - min_price {
            max_profit = prices[i] - min_price
        }
        
    }
    return max_profit
}
