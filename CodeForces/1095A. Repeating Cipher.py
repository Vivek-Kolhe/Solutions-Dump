n = int(input())
s = input()
res = ""
i, itr = 0, 1
while i < n:
  res += s[i]
  for j in range(itr):
    i += 1
  itr += 1
print(res)
