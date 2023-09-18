import string

alphabets = string.ascii_lowercase
for _ in range(int(input())):
  s = input()
  min_cost = int(input())
  costs = [alphabets.find(letter) + 1 for letter in s]
  cost = sum(costs)
  arr = sorted([(costs[i], i) for i in range(len(s))])
  while cost > min_cost:
    cost -= arr.pop()[0]
  new = sorted([(val, i) for i, val in arr])
  for i, val in new:
    print(alphabets[val-1], end="")
  print()
