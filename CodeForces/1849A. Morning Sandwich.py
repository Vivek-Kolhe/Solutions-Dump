for _ in range(int(input())):
  b, c, h = map(int, input().split())
  if b > c + h:
    print(2 * (c + h) + 1)
  else:
    print((2 * b) - 1)
