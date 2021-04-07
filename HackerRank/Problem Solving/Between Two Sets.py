#!/bin/python3

import math
import os
import random
import re
import sys

#
# Complete the 'getTotalX' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER_ARRAY a
#  2. INTEGER_ARRAY b
#

def getTotalX(arr_1, arr_2):
    factor_arr_1, factor_arr_2 = [], []

    for i in range(min(arr_1), min(arr_2) + 1):
        count = 0
        for j in arr_1:
            if i % j == 0:
                count += 1
        if count == len(arr_1):
            if i not in factor_arr_1:
                factor_arr_1.append(i)

    for i in factor_arr_1:
        count = 0
        for j in arr_2:
            if j % i == 0:
                count += 1
        if count == len(arr_2):
            if i not in factor_arr_2:
                factor_arr_2.append(i)
    return len(factor_arr_2)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    m = int(first_multiple_input[1])

    arr = list(map(int, input().rstrip().split()))

    brr = list(map(int, input().rstrip().split()))

    total = getTotalX(arr, brr)

    fptr.write(str(total) + '\n')

    fptr.close()
