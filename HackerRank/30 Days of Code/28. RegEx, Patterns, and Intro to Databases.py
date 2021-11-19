#!/bin/python3

import math
import os
import random
import re
import sys

def solution(inputs):
    result = []
    for i in range(len(inputs)):
        match = re.search(r"^[\w.+\-]+@gmail\.com$", inputs[i][-1])
        if match:
            result.append(inputs[i][0])
    for i in sorted(result):
        print(i)

if __name__ == '__main__':
    N = int(input().strip())
    inputs = []
    for N_itr in range(N):
        first_multiple_input = input().rstrip().split()
        inputs.append(first_multiple_input)
    solution(inputs)
