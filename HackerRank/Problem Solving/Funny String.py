#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the funnyString function below.
def funnyString(s):
    r = s[::-1]
    ascii_s = []
    for letter in s:
        ascii_s.append(ord(letter))
    ascii_r = ascii_s[::-1]
    check_s = []
    check_r = []
    for i in range(0, len(ascii_s) - 1):
        check_s.append(abs(ascii_s[i] - ascii_s[i + 1]))
        check_r.append(abs(ascii_r[i] - ascii_r[i + 1]))
    if check_s == check_r:
        return "Funny"
    else:
        return "Not Funny"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        s = input()

        result = funnyString(s)

        fptr.write(result + '\n')

    fptr.close()
