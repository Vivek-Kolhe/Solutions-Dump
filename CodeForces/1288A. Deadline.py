def ceil_val(d, x):
  x += 1
  if d % x == 0:
    return d // x
  return (d // x) + 1

for _ in range(int(input())):
  n, d = map(int, input().split())
  low, high = 0, n
  ans = 0
  while low <= high:
    mid = (low + high) // 2
    if mid + ceil_val(d, mid) <= n:
      print("YES")
      break
    else:
      high = mid - 1
  else:
    print("NO")
