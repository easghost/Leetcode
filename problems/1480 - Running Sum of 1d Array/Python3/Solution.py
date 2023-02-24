class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        if len(nums) < 1:
            return nums
        
        res = []
        sum = 0
        for num in nums:
            sum += num
            res.append(sum)
            
        return res