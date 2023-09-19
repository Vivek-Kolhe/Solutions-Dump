n = int(input())
pat, uniq = [], []
flag = True
for _ in range(n):
  temp = input()
  for letter in temp:
    if letter not in uniq:
      uniq.append(letter)
  pat.append(temp)
if not 3 > len(uniq) > 1:
  flag = False
if flag:
  x = pat[0][0]
  for i in range(n):
    if i == n // 2:
      if pat[i].count(x) != 1:
        flag = False
        print("NO")
        break
    else:
      if pat[i].count(x) != 2:
        flag = False
        print("NO")
        break
  if flag:
    print("YES")
else:
  print("NO")
