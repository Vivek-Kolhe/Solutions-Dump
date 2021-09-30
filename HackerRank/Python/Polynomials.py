import numpy

arr = list(map(float, input().split()))
val = float(input())

print(numpy.polyval(arr, val))
