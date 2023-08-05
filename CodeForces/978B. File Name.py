n = int(input())
msg = input()
count = 0
res = msg[0:2]
for i in range(2, n):
  if res[-1] == res[-2] == msg[i] == 'x':
    count += 1
  else:
    res += msg[i]
print(count)
