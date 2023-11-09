k = int(input())
n = list(map(int, list(input())))
s = sum(n)
if s > k:
  print(0)
else:
  n.sort()
  i = 0
  cnt = 0
  while s < k:
    if n[i] < 9:
      diff = 9 - n[i]
      n[i] += diff
      s += diff
      cnt += 1
    i += 1
  print(cnt)
