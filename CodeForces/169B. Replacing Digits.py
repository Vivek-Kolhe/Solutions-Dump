a = list(map(int, list(input())))
b = list(map(int, list(input())))
n, m = len(a), len(b)
b.sort(reverse=True)
j = 0
for i in range(n):
  if j < m and a[i] < b[j]:
    a[i] = b[j]
    j += 1
for i in range(n):
  print(a[i], end="")
print()
