n = int(input())
home = input()
l, r = 0, 0
for _ in range(n):
  trav = input().split("->")
  if trav[0] == home:
    l += 1
  if trav[-1] == home:
    r += 1
if l == r:
  print("home")
else:
  print("contest")
