def counter(s):
  h = {}
  for letter in s:
    if letter != " ":
      if h.get(letter) is None:
        h[letter] = 0
      h[letter] += 1
  return h
  
head = input()
text = input()
c1, c2 = counter(head), counter(text)
keys = c2.keys()
for key in keys:
  if c1.get(key) is None:
    print("NO")
    break
  if c1[key] < c2[key]:
    print("NO")
    break
else:
  print("YES")
