n = int(input())
arr = list(map(int, input().split()))
a = arr.copy()
if arr[0] == 0 or arr[n-1] == 0:
  print(0)
  exit(0)
maxi, mini = max(arr), min(arr)
for i in range(mini, maxi + 1):
  for j in range(n):
    if arr[j] == i:
      a[j] = 0
    if j > 0:
      if a[j] == 0 and a[j-1] == 0:
        print(i)
        exit(0)
    elif j < n - 1:
      if a[j] == 0 and a[j+1] == 0:
        print(i)
        exit(0)
  if a[0] == 0 or a[n-1] == 0:
    print(i)
    break
