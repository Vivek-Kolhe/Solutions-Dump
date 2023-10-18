def divisors(n):
  res = []
  for i in range(2, n):
    if n % i == 0:
      res.append(i)
  return res

n = int(input())
s = input()
divi = divisors(n)
for d in divi:
  s = s[:d][::-1] + s[d:]
s = s[::-1]
print(s)
