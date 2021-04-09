#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the fairRations function below.
def fairRations(B):
    B = [i for i, element in enumerate(B) if element % 2]
    if len(B) % 2:
        return "NO"
    counter = 0
    for i in range(0, len(B), 2):
        counter+=(B[i+1] - B[i]) * 2
    return counter
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    N = int(input())

    B = list(map(int, input().rstrip().split()))

    result = fairRations(B)

    fptr.write(str(result) + '\n')

    fptr.close()
