s = list(input())
i = 0
res = []
while i < len(s):
  s.insert(0, s.pop(-1))
  ss = "".join(s)
  if ss not in res:
    res.append(ss)
  i += 1
print(len(res))
