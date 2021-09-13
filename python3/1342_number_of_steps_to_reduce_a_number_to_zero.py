# Runtime: 32 ms
# Memory Usage: 14.3 MB

class Solution:
    def numberOfSteps(self, num: int) -> int:
        if num == 0 | num < 0:
            return 0
        
        counter = 0
        while num > 0:
            if num % 2 != 0:
                num -= 1
                counter += 1
            if num > 0:
                num = num / 2
                counter += 1
        
        return counter