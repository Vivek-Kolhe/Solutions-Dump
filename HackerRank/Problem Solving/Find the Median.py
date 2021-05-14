#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the findMedian function below.
def findMedian(arr):
    arr1 = sorted(arr)
    median = int((len(arr1) + 1) / 2)
    return arr1[median - 1]
  
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    result = findMedian(arr)

    fptr.write(str(result) + '\n')

    fptr.close()