import string
alphabets = string.ascii_lowercase
for _ in range(int(input())):
  n = int(input())
  s = input()
  i = 0
  res = ""
  while i < n:
    if i + 2 <= n - 1:
      if s[i+2] == "0":
        if i + 3 <= n - 1:
          if s[i+3] == "0":
            res += alphabets[int(s[i])-1]
            i += 1
            continue
      if s[i+2] == "0":
        res += alphabets[int(s[i:i+2])-1]
        i += 3
      else:
         res += alphabets[int(s[i])-1]
         i += 1
    else:
      res += alphabets[int(s[i])-1]
      i += 1
  print(res)
