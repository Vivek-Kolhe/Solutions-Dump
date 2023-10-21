a, b, s = map(int, input().split())
a, b = abs(a), abs(b)
buff = s - (a + b)
if s >= a + b:
  if buff % 2 == 0:
    print("Yes")
  else:
    print("No")
else:
  print("No")
