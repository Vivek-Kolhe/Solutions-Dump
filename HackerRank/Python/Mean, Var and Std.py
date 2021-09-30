import numpy

arr = []
n, m = map(int, input().split())
for _ in range(n):
    arr.append(list(map(int, input().split())))

print(numpy.mean(arr, axis = 1), numpy.var(arr, axis = 0), str(round(numpy.std(arr), 11)), sep = "\n")
