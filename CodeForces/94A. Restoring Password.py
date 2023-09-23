binary = input()
h = {}
for i in range(10):
  h[input()] = str(i)
res = ""
for i in range(0, 80, 10):
  res += h[binary[i:i+10]]
print(res)
