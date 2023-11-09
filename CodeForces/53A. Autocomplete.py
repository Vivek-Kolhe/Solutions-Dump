s = input()
same = []
for _ in range(int(input())):
  t = input()
  if t.startswith(s):
    same.append(t)
if len(same) > 0:
  print(min(same))
else:
  print(s)
