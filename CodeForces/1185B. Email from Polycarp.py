from itertools import groupby

for _ in range(int(input())):
  s = ["".join(item) for _, item in groupby(input())]
  t = ["".join(item) for _, item in groupby(input())]
  for i in range(len(s)):
    if not (s[i] in t[i] and len(s) == len(t)):
      print("NO")
      break
  else:
    print("YES")
