x1, y1 = map(int, input().split())
x2, y2 = map(int, input().split())
 
if x1 == x2:
  print(6 + (2 * abs(y1 - y2)))
elif y1 == y2:
  print(6 + (2 * abs(x1 - x2)))
else:
  dist = abs(x1 - x2) + abs(y1 - y2)
  print((2 * (dist)) + 3 + 1)
