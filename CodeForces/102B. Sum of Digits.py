n = input()
length= len(n)
count = 0
while length != 1:
  n = str(sum(map(int, list(n))))
  length = len(n)
  count += 1
print(count)
