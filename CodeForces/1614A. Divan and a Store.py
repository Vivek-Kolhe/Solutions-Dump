for _ in range(int(input())):
  n, l, r, k = map(int, input().split())
  arr = sorted(list(map(int, input().split())))
  count, spent = 0, 0
  for i in range(n):
    if arr[i] <= r and arr[i] >= l:
      count += 1
      spent += arr[i]
    if spent >= k:
      if spent > k:
        count -= 1
      break
  print(count)
