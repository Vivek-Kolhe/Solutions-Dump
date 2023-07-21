for _ in range(int(input())):
  a, b = map(int, input().split())
  if a >= 2:
    print(1)
  elif a == 1 and b >=3:
    print(2)
  else:
    print(a + b)
