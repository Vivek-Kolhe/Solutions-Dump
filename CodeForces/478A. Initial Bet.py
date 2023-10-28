arr = list(map(int, input().split()))
s = sum(arr)
if s % 5 == 0 and s > 0:
  print(s // 5)
else:
  print(-1)
