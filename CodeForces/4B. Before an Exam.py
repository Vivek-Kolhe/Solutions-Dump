import math

d, s = map(int, input().split())
mins, maxs = [], []
for _ in range(d):
  a, b = map(int, input().split())
  mins.append(a)
  maxs.append(b)
mini = sum(mins)
maxi = sum(maxs)
if mini <= s <= maxi:
  print("YES")
  s -= mini
  for i in range(d):
    diff = min(s, maxs[i] - mins[i])
    print(mins[i] + diff, end=" ")
    s -= diff
  print()
else:
  print("NO")
