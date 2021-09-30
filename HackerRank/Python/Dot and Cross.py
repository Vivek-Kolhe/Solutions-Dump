import numpy

n = int(input())
matrix_1 = [list(map(int, input().split())) for _ in range(n)]
matrix_2 = [list(map(int, input().split())) for _ in range(n)]

print(numpy.dot(matrix_1, matrix_2))
