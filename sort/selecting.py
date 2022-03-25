def selecting(l):
    for i in range(len(l) - 1):
        min_idx = 0
        for j in range(len(l[i:]) - 1):
            if l[min_idx] > l[j]:
                min_idx = j
        l[i], l[min_idx] = l[min_idx], l[i]
    return l


l = [2, 3, -1, 5, 4, 1, 0, 8]
print(selecting(l))