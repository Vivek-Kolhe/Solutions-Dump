#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the superReducedString function below.
def superReducedString(s):
    st = []
    i = 0
    while i < len(s):
        if len(st) != 0 and st[-1] == s[i]:
            i += 1
            st.pop(-1)
        else:
            st.append(s[i])
            i += 1
    if len(st) == 0:
        return "Empty String"
    else:
        return "".join(i for i in st)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    s = input()

    result = superReducedString(s)

    fptr.write(result + '\n')

    fptr.close()
