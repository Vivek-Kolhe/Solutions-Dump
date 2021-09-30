import numpy

dimensions = tuple(map(int, input().split()))

print(numpy.zeros(dimensions, dtype = numpy.int), numpy.ones(dimensions, dtype = numpy.int), sep = "\n")
