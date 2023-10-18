for _ in range(int(input())):
  n, k = map(int, input().split())
  a = list(map(int, input().split()))
  b = list(map(int, input().split()))
  for i in range(n):
    a[i] = [a[i], i]
  a.sort()
  b.sort()
  res = [0] * n
  for i in range(n):
    res[a[i][-1]] = b[i]
  print(*res)
