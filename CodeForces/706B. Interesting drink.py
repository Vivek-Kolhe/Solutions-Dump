n = int(input())
arr = list(map(int, input().split()))
arr.sort()
for _ in range(int(input())):
  c = int(input())
  low, high = 0, n - 1
  ans = -1
  while low <= high:
    mid = (low + high) // 2
    if arr[mid] > c:
      ans = mid
      high = mid - 1
    else:
      low = mid + 1
  print(ans if ans != -1 else n)
