def min_val(l):
    min_idx = 0
    for i in range(len(l) - 1):
        if l[min_idx] > l[i]:
            min_idx = i
    return l[min_idx]


l = [-1, 2, 3, 5, 4, 1, 0, 8]
print(min_val(l))