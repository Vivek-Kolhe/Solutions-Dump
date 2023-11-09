def counter(s):
  h = {}
  for letter in s:
    if h.get(letter) is None:
      h[letter] = 0
    h[letter] += 1
  return h
  
t = input()
tl = len(t)
hate = ""
for letter in t:
  if letter != "a":
    hate += letter
tc = counter(t)
for key in tc.keys():
  if key != "a" and tc[key] % 2 != 0:
    print(":(")
    break
else:
  n = len(hate)
  hate = hate[:n//2]
  i, j = 0, tl - (n // 2)
  while i < n // 2 and j < tl:
    if hate[i] != t[j]:
      print(":(")
      break
    i += 1
    j += 1
  else:
    print(t[:tl-(n//2)])
