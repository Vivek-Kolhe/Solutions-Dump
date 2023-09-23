n, m = map(int, input().split())
arr = list(map(int, input().split()))
missing = sorted(list(map(int, input().split())))[::-1]
for i in range(n):
  if arr[i] == 0:
    arr[i] = missing.pop(0)
if arr == sorted(arr):
  print("No")
else:
  print("Yes")
