import numpy

arr = []
n, m = map(int, input().split())
for _ in range(n):
    arr.append(list(map(int, input().split())))

print(numpy.max(numpy.min(arr, axis = 1)))
