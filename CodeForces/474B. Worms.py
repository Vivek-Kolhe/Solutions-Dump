n = int(input())
arr = list(map(int, input().split()))
labels = int(input())
label = list(map(int, input().split()))
s = 0
for i in range(n):
  store = arr[i]
  arr[i] += s
  s += store
for l in label:
  low, high = 0, n - 1
  ans = -1
  while low <= high:
    mid = (low + high) // 2
    if l <= arr[mid]:
      ans = mid
      high = mid - 1
    else:
      low = mid + 1
  print(ans + 1)
