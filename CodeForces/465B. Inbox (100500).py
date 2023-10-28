n = int(input())
arr = list(map(int, input().split()))
ind = []
inc = 0
for i in range(n):
  if arr[i] == 1:
    ind.append(i)
    inc += 1
gaps = 0
for i in range(inc - 1):
  if ind[i+1] - ind[i] > 1:
    gaps += 1
print(inc + gaps)
