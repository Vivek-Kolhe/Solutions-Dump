for _ in range(int(input())):
  n, k = map(int, input().split())
  arr = list(map(int, input().split()))
  even = 0
  ans = k - 1
  for ele in arr:
    rem = ele % k
    if rem == 0:
      print(0)
      break
    ans = min(ans, k - rem)
    if k == 4:
      if ele % 2 == 0:
        even += 1
      ans = min(ans, 2 - even)
      if even == 2:
        print(0)
        break
  else:
    print(ans)
