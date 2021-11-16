for _ in range(int(input())):
  a, b = map(int, input().split())
  if a == b:
    print(0)
  else:
    diff = abs(a - b)
    if diff % 10 != 0:
      print((diff // 10) + 1)
    else:
      print(diff // 10)
