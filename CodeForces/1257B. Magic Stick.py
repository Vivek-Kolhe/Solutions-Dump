for _ in range(int(input())):
  x, y = map(int, input().split())
  if x == 1:
    if y == 1:
      print("YES")
    else:
      print("NO")
  elif x == 2 or x == 3:
    if y < 4:
      print("YES")
    else:
      print("NO")
  else:
    print("YES")
