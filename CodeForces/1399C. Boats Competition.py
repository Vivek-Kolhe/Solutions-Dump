def wt(mid, n, arr):
  l, r = 0, n - 1
  cnt = 0
  while l < r:
    if arr[l] + arr[r] > mid:
      r -= 1
    elif arr[l] + arr[r] < mid:
      l += 1
    else:
      cnt += 1
      l += 1
      r -= 1
  return cnt
  
for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  arr.sort()
  if n < 2:
    print(0)
  elif n == 2:
    print(1)
  else:
    low = arr[0] + arr[1]
    high = arr[n-1] + arr[n-2]
    ans = 0
    for i in range(low, high + 1):
      ans = max(ans, wt(i, n, arr))
    print(ans)
