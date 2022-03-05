# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def firstBadVersion(self, n: int) -> int:
        l,r = 0, n
        while l != r:
            m = (l+r)
            if isBadVersion(m-1) == False and isBadVersion(m)== True:
                return m
            else:
                if isBadVersion(m) == False:
                    l = m +1
                else:
                    r = m -1
        return -1