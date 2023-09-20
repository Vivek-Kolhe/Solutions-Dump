for _ in range(int(input())):
  n, m = map(int, input().split())
  if n == m == 1:
    print(0)
  else:
    print(n - 1 + m - 1 + min(n, m))
