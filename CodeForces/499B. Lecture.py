n, m = map(int, input().split())
first, second = [], []
for _ in range(m):
  x, y = input().split()
  first.append(x)
  second.append(y)
record =input().split()
res = []
for word in record:
  ind = first.index(word)
  if len(first[ind]) <= len(second[ind]):
    res.append(first[ind])
  else:
    res.append(second[ind])
print(*res)
