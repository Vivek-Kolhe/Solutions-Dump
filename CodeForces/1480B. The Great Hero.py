def ceil(a, b):
  if a % b == 0:
    return a // b
  return (a // b) + 1

for _ in range(int(input())):
  a, b, n = map(int, input().split())
  atk = list(map(int, input().split()))
  hp = list(map(int, input().split()))
  x = list(zip(atk, hp))
  x.sort(key=lambda x: x[0])
  for i in range(n):
    curr = x[i]
    rounds = ceil(curr[-1], a)
    b -= rounds * curr[0]
  if b < -x[-1][0]:
    print("NO")
  else:
    print("YES")
