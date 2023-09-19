n, k = map(int, input().split())
num = "1" + ("0" * (n - 1))
num_last = "9" * n
num, num_last = int(num), int(num_last)
flag = False
for i in range(num, num_last + 1):
  if i % k == 0:
    print(i)
    flag = True
    break
if not flag:
  print(-1)
