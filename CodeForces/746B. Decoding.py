n = int(input())
word = input()
res = word[0]
if n % 2 != 0:
  for i in range(1, n):
    if i % 2 != 0:
      res = word[i] + res
    else:
      res += word[i]
else:
  for i in range(1, n):
    if i % 2 != 0:
      res += word[i]
    else:
      res = word[i] + res
print(res)
