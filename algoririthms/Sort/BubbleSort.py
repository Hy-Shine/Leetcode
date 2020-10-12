class Solution:
    def bubbleSort(self, array):
        length = len(array)
        for i in range(length):
            for j in range(length - i - 1):
                if array[j] > array[j + 1]:
                    array[j], array[j +1] = array[j + 1], array[j]
        return array


#  testing cases
array1 = Solution()
list1 = [1, 5, 9, 3, 5, 2, 10, 3, 90, 43, 21]
print(array1.bubbleSort(list1))
list2 = [2, 1, 0]
print(array1.bubbleSort(list2))
list3 = [1]
print(array1.bubbleSort(list3))
list4 = []
print(array1.bubbleSort(list4))