#!/bin/python3

import math
import os
import random
import re
import sys

# Complete the timeInWords function below.
def timeInWords(h, m):
    hour = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve"]
    minutes = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "quarter", "sixteen", "seventeen", "eighteen", "nineteen", "twenty", "twenty one", "twenty two", "twenty three", "twenty four", "twenty five", "twenty six", "twenty seven", "twenty eight", "twenty nine"]

    if m == 00:
        return f"{hour[h-1]} o' clock"
    elif m == 1:
        return f"{minutes[m-1]} minute past {hour[h-1]}"
    elif 1 < m <= 14 or 15 < m <= 29:
        return f"{minutes[m-1]} minutes past {hour[h-1]}"
    elif m == 15:
        return f"{minutes[m-1]} past {hour[h-1]}"
    elif m == 30:
        return f"half past {hour[h-1]}"
    elif 30 < m <= 44 or 45 < m <= 59:
        return f"{minutes[-(m-30)]} minutes to {hour[h]}"
    else:
        return f"quarter to {hour[h]}"

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    h = int(input())

    m = int(input())

    result = timeInWords(h, m)

    fptr.write(result + '\n')

    fptr.close()
