for _ in range(int(input())):
  n = int(input())
  s = list(input())
  l, r = 0, n - 1
  while s[l] != s[r] and l < r:
    l += 1
    r -= 1
  print(r - l + 1)
