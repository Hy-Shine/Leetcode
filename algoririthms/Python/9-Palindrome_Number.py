class Solution1:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        x = str(x)
        i, j = 0, len(x) - 1
        while i <= j:
            if x[i] != x[j]:
                return False
            i += 1
            j -= 1
        return True


class Solution2:
    def isPalindrome(self, x: int) -> bool:
        if x < 0:
            return False
        return str(x) == str(x)[::-1]


# testing cases
number = Solution1()
number.isPalindrome(0)  # True
number.isPalindrome(123)  # False
number.isPalindrome(-12)  # False
number.isPalindrome(121)  # True