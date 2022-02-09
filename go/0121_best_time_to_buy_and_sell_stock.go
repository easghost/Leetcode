func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    maxProfit := 0
    minPrice := prices[0]
    for i := range prices {
        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if prices[i]-minPrice > maxProfit {
            maxProfit = prices[i] - minPrice
        }
    }
	
    return maxProfit
}