"""
Kind of dumb solution but works lmao!
"""

import math

n = int(input())
arr = list(map(int, input().split()))
h = {
  1 : 0,
  2 : 0,
  3 : 0,
  4 : 0
  }
for ele in arr:
  h[ele] += 1
cnt = h[4]
h[4] = 0
mini = min(h[3], h[1])
cnt += mini
h[3] -= mini
h[1] -= mini
cnt += h[3]
h[3] -= h[3]
twos = h[2] // 2
cnt += twos
h[2] -= (twos * 2)
if h[2] > 0 and h[1] > 0:
  h[2] -= 1
  if h[1] < 2:
    h[1] -= 1
  else:
    h[1] -= 2
  cnt += 1
cnt += h[4] + h[3] + h[2]
cnt += math.ceil(h[1] / 4)
print(cnt)
