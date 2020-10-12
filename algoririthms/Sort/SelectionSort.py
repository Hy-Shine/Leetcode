# Selection Sort
class Solution:
    def findMin(self, array):
        if len(array) > 0:
            minvalue = array[0]
            j = 0
            for i in range(1, len(array)):
                if array[i] < minvalue:
                    minvalue = array[i]
                    j = i
            return j
        return None

    def selectionSort(self, array2):
        length = len(array2)
        for i in range(len(array2) - 1):
            v = self.findMin(array2[i:length])
            array2[i], array2[v+i] = array2[v+i], array2[i]
        return array2

# testing cases
arrays = Solution()
print(arrays.selectionSort([10, 46, 51, 90, 43, 56]))
print(arrays.selectionSort([]))