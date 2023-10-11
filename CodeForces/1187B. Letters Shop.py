from collections import *
n = int(input())
string = input()
d = defaultdict(list)
i = 1
for letter in string:
  d[letter].append(i)
  i += 1
for _ in range(int(input())):
  name = Counter(input())
  ans = -1
  for key in name: 
    ans = max(d[key][name[key]-1], ans)
  print(ans)
