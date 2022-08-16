for _ in range(int(input())):
  string = input()
  count = 0
  order = "".join([i for i in string if i == i.upper()])
  for i in order:
    if string.find(i.lower()) < string.find(i):
      count += 1
  if count == 3:
    print("YES")
  else:
    print("NO")
