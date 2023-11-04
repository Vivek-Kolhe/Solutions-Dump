for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  if arr[0] != n:
    print("NO")
    continue
  a = arr.copy()
  l = n
  diff = 0
  res = [0] * max(arr)
  i, j = 0, 1
  while i < n:
    res[i] = l
    diff += 1
    while j <= n and a[-j] - diff == 0:
      j += 1
      l -= 1
    i += 1
  for i in range(n):
    if arr[i] != res[i]:
      print("NO")
      break
  else:
    print("YES")
