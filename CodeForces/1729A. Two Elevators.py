for _ in range(int(input())):
  a, b, c = map(int, input().split())
  at = abs(a - 1)
  bt = abs(b - c) + abs(c - 1)
  if at < bt: print(1)
  elif bt < at: print(2)
  else: print(3)
