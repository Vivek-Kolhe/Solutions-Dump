for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  s = sum(arr)
  ans, c = 0, 0
  arr.sort()
  if s % 2 != 0:
    arr[0] -= 1
    ans += 1
  k = s // 2
  i = 0
  while i < n and c + arr[i] <= k:
    c += arr[i]
    i += 1
  ans += k + (n - i)
  print(ans)
