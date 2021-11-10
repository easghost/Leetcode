// Runtime: 176 ms
// Memory Usage: 35.5 MB

class Solution {
    fun maxProfit(prices: IntArray): Int {
        var res = 0
        var i = 0
        var count = prices.count()-2
        while(i <= count) {
            if(prices[i] < prices[i+1]) {
                res += prices[i+1]-prices[i]
            }
            i++
        }
        
        return res
    }
}