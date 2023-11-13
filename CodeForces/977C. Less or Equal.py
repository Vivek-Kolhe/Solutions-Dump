n, k = map(int, input().split())
arr = list(map(int, input().split()))
arr.sort()
if k == 0:
  if arr[0] > 1:
    print(arr[0] - 1)
  else:
    print(-1)
else:
  if n == 1:
    print(arr[0])
  elif n == k:
    print(arr[n-1])
  else:
    if arr[k] == arr[k-1]:
      print(-1)
    else:
      print(arr[k-1])
