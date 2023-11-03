for _ in range(int(input())):
  n, x = map(int, input().split())
  arr = list(map(int, input().split()))
  arr.sort(reverse=True)
  cnt = 1
  s = 0
  for i in range(n):
    if (s + arr[i]) / cnt >= x:
      s += arr[i]
      cnt += 1
  print(cnt - 1)
