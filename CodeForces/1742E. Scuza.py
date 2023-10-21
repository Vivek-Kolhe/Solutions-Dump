for _ in range(int(input())):
  n, q = map(int, input().split())
  arr = list(map(int, input().split()))
  que = list(map(int, input().split()))
  pre = [0]
  pref = []
  for i in range(n):
    pre.append(pre[-1] + arr[i])
    if i == 0:
      pref.append(arr[i])
    else:
      pref.append(max(pref[-1], arr[i]))
  for qu in que:
    low, high = 0, n - 1
    ans = -1
    while low <= high:
      mid = (low + high) // 2
      if pref[mid] > qu:
        high = mid - 1
      else:
        ans = mid
        low = mid + 1
    print(pre[ans+1], end=" ")
  print()
