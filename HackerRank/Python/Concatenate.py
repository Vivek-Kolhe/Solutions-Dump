import numpy

n, m, p = map(int, input().split())
arr_1 = numpy.array([list(map(int, input().split())) for _ in range(n)])
arr_2 = numpy.array([list(map(int, input().split())) for _ in range(m)])

print(numpy.concatenate((arr_1, arr_2), axis = 0))
