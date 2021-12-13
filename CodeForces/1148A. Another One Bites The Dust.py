a, b, c = map(int, input().split())
if a == b and b == c:
  print(a * 4)
elif a == b:
  print((a * 2) + (c * 2))
else:
  print((2 * min([a, b])) + 1 + (2 * c))
