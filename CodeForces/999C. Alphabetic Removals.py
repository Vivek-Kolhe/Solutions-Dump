def counter(s):
  h = {}
  for letter in s:
    if h.get(letter) is None:
      h[letter] = 0
    h[letter] += 1
  return h
  
n, k = map(int, input().split())
s = input()
c = counter(s)
keys = sorted(c.keys())
for key in keys:
  if c[key] >= k:
    s = s.replace(key, "", k)
    k -= k
  else:
    s = s.replace(key, "", c[key])
    k -= c[key]
  if k <= 0:
    break
print(s)
