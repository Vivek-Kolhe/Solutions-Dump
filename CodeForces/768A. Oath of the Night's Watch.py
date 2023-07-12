n = int(input())
arr = list(map(int, input().split()))
max, min = max(arr), min(arr)
maxc, minc = arr.count(max), arr.count(min)
res = n - (maxc + minc)
if res >= 0:
  print(res)
else:
  print(0)
