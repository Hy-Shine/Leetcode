# Solution1
class Solution:
    def reverseString(self, s) -> None:
        """
        Do not return anything, modify s in-place instead.
        """
        length = len(s)
        for i in range(len(s) // 2):
            s[i], s[length - i - 1] = s[length - i - 1], s[i]


# testing cases
s1 = []
s2 = ['a','c','g','x']
lists = Solution()
lists.reverseString(s1)
lists.reverseString(s2)
print(s1)
print(s2)