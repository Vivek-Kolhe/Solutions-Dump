n = input()
res = ""
for i in range(len(n)):
  digit = int(n[i])
  if i == 0 and digit == 9:
    res += n[i]
    continue
  res += str(min(digit, 9 - digit))
print(int(res))
