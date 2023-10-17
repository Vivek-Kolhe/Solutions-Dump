n = int(input())
a = int(input())
b = int(input())
c = int(input())
if n == 1:
  print(0)
else:
  if c == min(a, b, c):
    print(min(a, b) + (c * (n - 2)))
  else:
    print(min(a, b) * (n - 1))
