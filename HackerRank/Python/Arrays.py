import numpy

def arrays(arr):
    rev = arr[::-1]
    return numpy.array(rev, float)

arr = input().strip().split(' ')
result = arrays(arr)
print(result)
