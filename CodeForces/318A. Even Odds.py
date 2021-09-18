n, pos = map(int, input().split())
if n % 2 == 0:
  mid = n // 2
else:
  mid = (n // 2) + 1
if pos > mid:
  print(2 * (pos - mid))
else:
  print((2 * pos) - 1)
