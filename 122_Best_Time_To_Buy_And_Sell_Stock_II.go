func maxProfit(prices []int) int {
    if len(prices) == 0 || len(prices) == 1 {
        return 0
    }
    
    total_profit := 0
    buy := prices[0]
    sold := prices[0]
    for i := 1; i < len(prices); i++ {
        if prices[i] < buy && buy == sold {
            buy = prices[i]
            sold = prices[i]
        } else if prices[i] >= sold {
            sold = prices[i]
            if i == len(prices) -1 {
                total_profit += (sold - buy)
            }
        } else if prices[i] < sold && prices[i-1] == sold {
            total_profit += (sold - buy)
            buy = prices[i]
            sold = prices[i]
        }
    } 
    return total_profit
}
