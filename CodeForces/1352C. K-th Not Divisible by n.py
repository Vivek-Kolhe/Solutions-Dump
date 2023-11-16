import math

for _ in range(int(input())):
  n, k = map(int, input().split())
  end = n * math.ceil(k / (n - 1))
  start = end - n
  mod = k % (n - 1)
  if mod == 0:
    mod = n - 1
  print(start + mod)
