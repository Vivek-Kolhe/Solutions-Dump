a = int(input())
b = int(input())
c = a + b
a = str(a).replace("0", "")
b = str(b).replace("0", "")
c = str(c).replace("0", "")
if int(a) + int(b) == int(c):
  print("YES")
else:
  print("NO")
