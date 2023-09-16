for _ in range(int(input())):
  n, m = map(int, input().split())
  res = ["B" * m for i in range(n)]
  res[0] = "W" + "B" * (m - 1)
  for r in res:
    print(r)
