for _ in range(int(input())):
  x1, p1 = map(int, input().split())
  x2, p2 = map(int, input().split())
  min_p = min(p1, p2)
  p1, p2 = p1 - min_p, p2 - min_p
  p1, p2 = min(p1, 10), min(p2, 10)
  x, y = x1 * (10 ** p1), x2 * (10 ** p2)
  if x > y:
    print('>')
  elif x < y:
    print('<')
  else:
    print('=')
