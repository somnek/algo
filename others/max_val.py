def max_val(l):
    for i in range(len(l) - 1):
        if l[i] > l[i+1]:
            l[i], l[i+1] = l[i+1], l[i]
    return l[-1]
    

l = [2, 3, 5, 0, 4, 1]
print(max_val(l))