#!/bin/python3
import sys

n = int(input().strip())
a = list(map(int, input().strip().split(' ')))
swaps = 0

for i in range(n):
    for j in range(n - i - 1):
        if a[j] > a[j + 1]:
            a[j], a[j + 1] = a[j + 1], a[j]
            swaps += 1

print(f"Array is sorted in {swaps} swaps.\nFirst Element: {a[0]}\nLast Element: {a[-1]}")
