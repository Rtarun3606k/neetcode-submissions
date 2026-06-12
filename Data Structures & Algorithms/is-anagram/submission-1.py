class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        s1 = sorted(s)
        s2 = sorted(t)
        if len(s) != len(t):
            return  False
        if s1 ==s2:
            return True
        return  False
        