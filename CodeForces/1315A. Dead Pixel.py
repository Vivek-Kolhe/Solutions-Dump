for _ in range(int(input())):
  a, b, x, y = map(int, input().split())
  x += 1
  y += 1
  top = (x - 1) * b
  bottom = (a - x) * b
  left = (y - 1) * a
  right = (b - y) * a
  print(max([top, bottom, left, right]))
