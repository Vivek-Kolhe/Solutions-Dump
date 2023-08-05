for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  res = []
  count = 0
  for i in range(n):
    if arr[i] == 0:
      count += 1
    else:
      res.append(count)
      count = 0
  res.append(count)
  print(max(res))
