import math
import sys

for _ in range(int(input())):
  n = int(input())
  sq = int(math.sqrt(n))
  ans = sq -  1 + (n // sq) - 1
  if n % sq:
    ans += 1
  print(ans)
