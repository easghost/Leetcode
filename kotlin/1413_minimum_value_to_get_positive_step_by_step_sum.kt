// Runtime: 144 ms
// Memory Usage: 33.5 MB

class Solution {
    fun minStartValue(nums: IntArray): Int {
        var res = 1
        var sum = 1

        for(num in nums) {
            sum += num
            if(sum < 1) {
                res += 1 - sum
                sum = 1
            }
        }

        return res
    }
}