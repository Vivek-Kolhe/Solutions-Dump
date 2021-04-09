#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the viralAdvertising function below.
def viralAdvertising(n):
    total_liked = 0
    liked = 5 // 2
    for i in range(n):
        shared = liked * 3
        total_liked += liked
        liked = shared // 2
    return total_liked
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    result = viralAdvertising(n)

    fptr.write(str(result) + '\n')

    fptr.close()
