# Runtime: 28 ms
# Memory Usage: 14.3 MB

class Solution:
    def minStartValue(self, nums: List[int]) -> int:
        res = 1
        cur = res
        for n in nums:
            cur += n
            if cur < 1:
                res += 1 - cur
                cur = 1
        return res