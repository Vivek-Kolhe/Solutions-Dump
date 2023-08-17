for _ in range(int(input())):
  x, k = map(int, input().split())
  if x % k != 0:
    print(1)
    print(x)
  else:
    print(2)
    if x % 2 == 0 and k % 2 == 0:
      print(1, x - 1)
    else:
      print(2, x - 2)
