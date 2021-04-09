#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the pangrams function below.
def pangrams(s):
    x = s.lower()
    s1 = "".join(x.split())
    s2 = sorted(list(set(s1)))
    string1 = "".join(s2)
    string2 = "abcdefghijklmnopqrstuvwxyz"
    if string1 == string2:
        return "pangram"
    return "not pangram"
  
if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = pangrams(s)

    fptr.write(result + '\n')

    fptr.close()
