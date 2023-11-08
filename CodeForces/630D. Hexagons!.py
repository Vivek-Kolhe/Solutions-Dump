n = int(input())
if n == 0:
  print(1)
else:
  nth = n * 6
  first = 6
  s = n * (first + nth)
  s //= 2
  print(s + 1)
