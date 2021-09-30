import numpy

n, m = map(int, input().split())
arr_1 = numpy.array([list(map(int, input().split())) for _ in range(n)])
arr_2 = numpy.array([list(map(int, input().split())) for _ in range(n)])

print(
        numpy.add(arr_1, arr_2),
        numpy.subtract(arr_1, arr_2),
        numpy.multiply(arr_1, arr_2),
        numpy.floor_divide(arr_1, arr_2),
        numpy.mod(arr_1, arr_2),
        numpy.power(arr_1, arr_2),
        sep = "\n"
     )
