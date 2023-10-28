a, b, c, n = map(int, input().split())
if c > a or c > b:
  print(-1)
else:
  bug, king = a - c, b - c
  s = bug + king + c
  if s < n:
    print(n - s)
  else:
    print(-1)
