pos = input()
c, d = pos[0], pos[1]
if (c == 'a' or c == 'h') and (d == '1' or d == '8'):
  print(3)
elif 'a' < c < 'h' and 1 < int(d) < 8:
  print(8)
else:
  print(5)
