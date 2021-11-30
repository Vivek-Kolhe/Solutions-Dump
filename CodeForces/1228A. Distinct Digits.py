def distinct(n):
  for i in range(len((n))):
    if n.count(n[i]) > 1:
      return False
  return True
 
a, b = map(int, input().split())
flag = False
for i in range(a, b + 1):
  if distinct(str(i)):
    flag = True
    print(i)
    break
if not flag:
  print(-1)
