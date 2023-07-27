a, b = map(int, input().split())
if a == 9 and b == 1:
  print(a, a + 1)
elif a > b:
  print(-1)
else:
  if b - a > 1:
    print(-1)
  else:
    if b - a == 0:
      a = str(a) + '0'
      b = int(a) + 1
    else:
      b = str(b) + '0'
      a = int(b) - 1
    print(a, b)
