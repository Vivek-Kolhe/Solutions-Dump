for _ in range(int(input())):
  k, x = map(int, input().split())
  low, high = 1, (2 * k) - 1
  ans = (2 * k) - 1
  while low <= high:
    mid = (low + high) // 2
    emo = 0
    if mid < k:
      emo = count(mid)
    else:
      emo = ((k * (k + 1)) // 2) + (((mid - k) * (k - 1 + k - (mid % k))) // 2)
    if emo >= x:
      ans = mid
      high = mid - 1
    else:
      low = mid + 1
  print(ans)
