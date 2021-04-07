#!/bin/python3

import os
import sys

#
# Complete the pageCount function below.
#
def pageCount(n, p):
    book = []
    for i in range(0, n + 1, 2):
        book.append((i, i + 1))
    rev_book = book[::-1]
    if p % 2 == 0:
        x = book.index((p, p + 1))
        y = rev_book.index((p, p + 1))
    else:
        x = book.index((p - 1, p))
        y = rev_book.index((p - 1, p))
    return min(x, y)

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    n = int(input())

    p = int(input())

    result = pageCount(n, p)

    fptr.write(str(result) + '\n')

    fptr.close()
