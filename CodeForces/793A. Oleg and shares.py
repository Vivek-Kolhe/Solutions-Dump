n, k = map(int, input().split())
arr = list(map(int, input().split()))
flag = True
for i in range(1, n):
  flag = flag and (arr[0] % k == arr[i] % k)
if not flag:
  print(-1)
else:
  min_ = min(arr)
  tot = 0
  for ele in arr:
    tot += ele - min_
  print(tot // k)
