n, m = map(int, input().split())
if n == 1 and m == 1:
  print(2)
else:
  count = 0
  for i in range(0, max(n, m)):
    for j in range(0, max(n, m)):
      if (i ** 2) + j == n and i + (j ** 2) == m:
        count += 1
  print(count)
