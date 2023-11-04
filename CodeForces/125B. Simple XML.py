s = input().split(">")
s.pop()
for i in range(len(s)):
  s[i] += ">"
st = []
lvl = 0
for i in range(len(s)):
  if "/" not in s[i]:
    spaces = " " * (2 * lvl)
    print(spaces + s[i])
    lvl += 1
  else:
    lvl -= 1
    spaces = " " * (2 * lvl)
    print(spaces + s[i])
