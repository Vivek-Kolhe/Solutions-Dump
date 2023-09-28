def get_area(arr, h):
  area = 0
  for i in range(len(arr)):
    if arr[i] <= h:
      area += h - arr[i]
  return area
  
for _ in range(int(input())):
  n, w = map(int, input().split())
  arr = list(map(int, input().split()))
  low, high = 0, (2 * (10 ** 9)) + 7
  ans = 0
  while low <= high:
    mid = (low + high) // 2
    if get_area(arr, mid) <= w:
      low = mid + 1
      ans = max(ans, mid)
    else:
      high = mid - 1
  print(ans)
