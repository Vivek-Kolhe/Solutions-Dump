n = int(input())
num = input()[::-1]
res = bin(int(num, 2) + int("1", 2))[2:]
res = ("0" * (n - (len(res))) + res)[::-1][:n]
count = 0
for i in range(n):
  if res[i] == num[::-1][i]:
    count += 1
print(n - count)
