n = int(input())
ticket = input()
first, second = ticket[:(2*n)//2], ticket[(2*n)//2:]
first = list(map(int, list(first)))
second = list(map(int, list(second)))
first.sort()
second.sort()
big, small = False, False
for i in range(n):
  if not first[i] > second[i]:
    small = True
  if not first[i] < second[i]:
    big = True
if big and small:
  print("NO")
else:
  print("YES")
