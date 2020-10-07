## Python3
class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        s = s.lstrip().rstrip()
        if len(s) > 0:
            listString = s.split(" ")
            return len(listString[-1])
        return 0


# testing cases
string = Solution()
print(string.lengthOfLastWord("hello world "))
print(string.lengthOfLastWord("  "))
