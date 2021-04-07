#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the staircase function below.
def staircase(n):
    for item in range(1, n + 1):
        hashtag = "#" * item
        space = " " * (n - item)
        print(f"{space}{hashtag}")

if __name__ == '__main__':
    n = int(input())

    staircase(n)
