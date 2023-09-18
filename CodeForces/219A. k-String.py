k = int(input())
s = input()
flag = True
for letter in s:
  if s.count(letter) % k != 0:
    print(-1)
    flag = False
    break
if flag:
  uniq = "".join(list(set(s)))
  res = ""
  for letter in uniq:
    res += letter * (s.count(letter) // k)
  print(res * k)
