def insertSort(List):
    if len(List) <= 1:
        return List
    for i in range(1, len(List)):
        for j in range(i, 0, -1):
            if List[j] < List[j - 1]:
                List[j], List[j - 1] = List[j - 1], List[j]
    return List


# testing cases
L = [7, 6, 1, 8, 2, 4, 10, -1, 0, 34]  #OK
L = []  #OK
L = [2, 0]  #OK