import math

names = ["Sheldon", "Leonard", "Penny", "Rajesh", "Howard"]
n = int(input())
x, i = 0, 0
while x < n:
  x += 5 * (2 ** i)
  i += 1
diff = n - (x - 5 * (2 ** (i - 1)))
ind = math.ceil(diff / 2 ** (i - 1))
print(names[ind-1])
