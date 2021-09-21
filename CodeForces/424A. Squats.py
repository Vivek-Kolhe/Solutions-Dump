n = int(input())
s = input()
sitting, standing = s.count("x"), s.count("X")
x = (n // 2) - min(standing, sitting)
print(x)
if standing > sitting:
  s = s.replace("X", "x", x)
elif standing < sitting:
  s = s.replace("x", "X", x)
print(s)
