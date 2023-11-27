for _ in range(int(input())):
  a, b, c = map(int, input().split())
  if (a == b and b == c) or (a % 2 == b % 2 == c % 2):
    print(1, 1, 1)
  else:
    res = [0, 0, 0]
    if a == b:
      res[2] = 1
    elif b == c:
      res[0] = 1
    elif c == a:
      res[1] = 1
    else:
      if abs(b - c) % 2 == 0:
        res[0] = 1
      elif abs(c - a) % 2 == 0:
        res[1] = 1
      else:
        res[2] = 1
    print(*res)
