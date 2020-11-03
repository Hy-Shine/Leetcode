class Solution:
    def twoSum(self, numbers, target: int):
        i, j = 0, len(numbers) - 1
        while i <= j:
            if numbers[i] + numbers[j] == target:
                return i + 1, j + 1
            elif numbers[i] + numbers[j] < target:
                i += 1
            else:
                j -= 1
        return -1, -1

# Testing case
L = Solution
print(L.twoSum(1, 3, 5, 9, 25), 14) -> [3,4]
print(L.twoSum([], 2)) -> [-1, -1]