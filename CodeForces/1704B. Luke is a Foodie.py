import sys

for _ in range(int(input())):
  n, x = map(int, input().split())
  arr = list(map(int, input().split()))
  maxi, mini = arr[0], arr[0]
  res = 0
  for i in range(1, n):
    if arr[i] > maxi:
      maxi = arr[i]
    if arr[i] < mini:
      mini = arr[i]
    if maxi - mini > 2 * x:
      res += 1
      maxi = arr[i]
      mini = arr[i]
  print(res)
