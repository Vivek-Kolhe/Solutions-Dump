#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the miniMaxSum function below.
def miniMaxSum(arr):
    i=0
    j=0
    for item in arr:
        i = i + item
        x = i - max(arr)
    for item in arr:
        j = j + item
        y = i - min(arr)
    print(x, y)

if __name__ == '__main__':
    arr = list(map(int, input().rstrip().split()))

    miniMaxSum(arr)
