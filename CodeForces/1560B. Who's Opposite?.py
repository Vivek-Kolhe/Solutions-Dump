for _ in range(int(input())):
  a, b, c = map(int, input().split())
  if abs(a - b) > 1 or ((a == 1 or a == 2) and (b == 1 or b == 2)):
    diff = abs(a - b)
    total = diff * 2
    if c > total or a > total or b > total:
      print(-1)
    else:
      if c > total // 2:
        print(c - diff)
      else:
        print(c + diff)
  else:
    print(-1)
