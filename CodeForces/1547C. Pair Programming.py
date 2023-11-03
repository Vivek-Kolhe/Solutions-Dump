for _ in range(int(input())):
  input()
  k, n, m = map(int, input().split())
  a = list(map(int, input().split()))
  b = list(map(int, input().split()))
  l, r = 0, 0
  flag = True
  res = []
  while l < n or r < m:
    if l < n and a[l] == 0:
      res.append(a[l])
      k += 1
      l += 1
    elif r < m and b[r] == 0:
      res.append(b[r])
      k += 1
      r += 1
    elif l < n and a[l] <= k:
      res.append(a[l])
      l += 1
    elif r < m and b[r] <= k:
      res.append(b[r])
      r += 1
    else:
      print(-1)
      flag = False
      break
  if flag:
    print(*res)
