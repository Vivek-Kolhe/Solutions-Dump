n, p1, p2, p3, t1, t2 = map(int, input().split())
periods = []
for _ in range(n):
  periods.append(list(map(int, input().split())))
c = 0
for i in range(n - 1):
  curr = periods[i]
  next = periods[i+1]
  c += (curr[-1] - curr[0]) * p1
  gap = next[0] - curr[-1]
  if gap <= t1:
    c += gap * p1
  elif gap <= t1 + t2:
    c += (t1 * p1) + (gap - t1) * p2
  else:
    c += (t1 * p1) + (t2 * p2) + (gap - t1 - t2) * p3
c += (periods[-1][-1] - periods[-1][0]) * p1
print(c)
