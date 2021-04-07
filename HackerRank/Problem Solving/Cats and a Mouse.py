#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the catAndMouse function below.
def catAndMouse(x, y, z):
    a = int(xyz[2]) - int(xyz[0])
    b = int(xyz[2]) - int(xyz[1])

    if abs(a) > abs(b):
        return "Cat B"
    elif abs(a) < abs(b):
        return "Cat A"
    elif abs(a) == abs(b):
        return "Mouse C"
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        xyz = input().split()

        x = int(xyz[0])

        y = int(xyz[1])

        z = int(xyz[2])

        result = catAndMouse(x, y, z)

        fptr.write(result + '\n')

    fptr.close()
