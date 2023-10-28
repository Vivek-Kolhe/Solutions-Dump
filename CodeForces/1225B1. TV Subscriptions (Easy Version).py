for _ in range(int(input())):
  n, k, d = map(int, input().split())
  arr = list(map(int, input().split()))
  ans = n
  for i in range(n - d + 1):
    s = set(arr[i:i+d])
    ans = min(len(s), ans)
  print(ans)
