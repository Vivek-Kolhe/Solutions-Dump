#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the plusMinus function below.
def plusMinus(arr):
    p = []
    n = []
    z = []

    for item in arr:
        if item > 0:
            p.append(item)
        elif item < 0:
            n.append(item)
        else:
            z.append(item)
    ratio1 = len(p)/len(arr)
    print(ratio1)
    ratio2 = len(n)/len(arr)
    print(ratio2)
    ratio3 = len(z)/len(arr)
    print(ratio3)
if __name__ == '__main__':
    n = int(input())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
