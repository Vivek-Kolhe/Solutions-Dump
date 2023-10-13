for _ in range(int(input())):
  l, r, a = map(int, input().split())
  res = r // a + r % a
  m = ((r // a) * a) - 1
  if m >= l:
    res = max(res, m // a + m % a)
  print(res)
