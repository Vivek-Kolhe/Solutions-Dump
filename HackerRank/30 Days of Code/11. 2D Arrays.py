#!/bin/python3
import math
import os
import random
import re
import sys

if __name__ == '__main__':
    arr = []
    for _ in range(6):
        arr.append(list(map(int, input().rstrip().split())))

count = 0
new = []

for x in range(0, 4):
    for y in range(0, 4):
        for i in range(x, x + 3):
            for j in range(y, y + 3):
                if i == x + 1 and (j == y or j == y + 2):
                    continue
                else:
                    count += arr[i][j]
        new.append(count)
        count = 0
print(max(new))
