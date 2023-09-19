from collections import Counter

arr = []
for _ in range(int(input())):
  h, m = map(int, input().split())
  arr.append((h, m))
c = Counter(arr)
print(c.most_common(1)[0][-1])
