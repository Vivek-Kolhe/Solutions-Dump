for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  ans = 0
  for i in range(n):
    if arr[i] == ans + 1:
      ans += 2
    else:
      ans += 1
  print(ans)
