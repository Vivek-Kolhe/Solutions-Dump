from collections import Counter
 
for _ in range(int(input())):
  string = ""
  n = int(input())
  for i in range(n):
    string += input()
  count = Counter(string)
  flag = False
  for i in count.values():
    if i % n != 0:
      flag = True
      print("NO")
      break
  if not flag:
    print("YES")
