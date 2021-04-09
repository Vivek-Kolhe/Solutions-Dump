#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the cutTheSticks function below.
def cutTheSticks(arr):
    new_arr = []
    while True:
        x = min(arr)
        for i in range(len(arr)):
            arr[i] = arr[i] - x
        new_arr.append(len(arr))
        for i in range(arr.count(0)):
            arr.remove(0)
        
        if len(arr) == 0:
            return new_arr
            break

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    result = cutTheSticks(arr)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
