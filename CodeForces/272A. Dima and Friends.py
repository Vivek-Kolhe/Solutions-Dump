n = int(input())
arr = list(map(int, input().split()))
s = sum(arr)
cnt = 0
for i in range(1, 6):
  if (s + i) % (n + 1) != 1:
    cnt += 1
print(cnt)
