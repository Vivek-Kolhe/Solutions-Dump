for _ in range(int(input())):
  n, k = map(int, input().split())
  arr = list(map(int, input().split()))
  ind, res = [], []
  for i in range(n):
    arr[i] = arr[i] % k
    if arr[i] == 0:
      arr[i] = k
    ind.append([-arr[i], i + 1])
  ind.sort()
  for i in ind:
    res.append(i[-1])
  print(*res)
