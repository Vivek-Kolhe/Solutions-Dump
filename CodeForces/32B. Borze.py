code = input()
res = ""
i = 0
while i < len(code):
  if code[i] == ".":
    i += 1
    res += "0"
  else:
    if code[i+1] == ".":
      res += "1"
    else:
      res += "2"
    i += 2
print(res)
