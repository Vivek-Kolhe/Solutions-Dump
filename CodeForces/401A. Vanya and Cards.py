n, x = map(int, input().split())
cards = abs(sum(list(map(int, input().split()))))
if cards % x == 0:
  print(cards // x)
else:
  print((cards // x) + 1)
