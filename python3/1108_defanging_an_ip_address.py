# Runtime: 32 ms
# Memory Usage: 14.3 MB

class Solution:
    def defangIPaddr(self, address: str) -> str:
        return address.replace(".", "[.]")