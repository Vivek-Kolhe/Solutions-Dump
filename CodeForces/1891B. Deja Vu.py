for _ in range(int(input())):
  n, q = map(int, input().split())
  arr = list(map(int, input().split()))
  queries = list(map(int, input().split()))
  h = {}
  for query in queries:
    h[query] = False
  for query in queries:
    if h[query]:
      continue
    x = pow(2, query)
    for i in range(n):
      if arr[i] % x == 0:
        arr[i] += x // 2
    h[query] = True
  print(*arr)
