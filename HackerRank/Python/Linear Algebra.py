import numpy

n = int(input())
arr = [list(map(float, input().split())) for _ in range(n)]
print(str(round(numpy.linalg.det(arr), 2)))
