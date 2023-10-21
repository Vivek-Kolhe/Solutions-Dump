n, m = map(int, input().split())
s = input()
t = input()
h = {}
for i in range(m - n + 1):
  temp = []
  cost = 0
  ss = t[i:i+n]
  for j in range(n):
    if s[j] != ss[j]:
      cost += 1
      
      temp.append(j + 1)
  h[cost] = temp
key = min(h.keys())
print(key)
print(*h[key])
