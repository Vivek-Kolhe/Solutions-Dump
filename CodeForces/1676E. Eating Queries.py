for _ in range(int(input())):
  n, q = map(int, input().split())
  arr = list(map(int, input().split()))
  arr.sort(reverse=True)
  s = 0
  for i in range(n):
    store = arr[i]
    arr[i] += s
    s += store
  for i in range(q):
    query = int(input())
    ans = -1
    low, high = 0, n - 1
    while low <= high:
      mid = (low + high) // 2
      if arr[mid] >= query:
        high = mid - 1
        ans = mid
      else:
        low = mid + 1
    print(ans + 1 if ans > -1 else -1)
