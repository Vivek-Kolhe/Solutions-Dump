n = int(input())
stairs = list(map(int, input().split()))
if n == 1:
  count, staircases = 1, [1]
else:
  count, staircases = stairs.count(1), []
  for i in range(1, n):
    if stairs[i] == 1:
      staircases.append(stairs[i - 1])
    if i == n - 1:
      staircases.append(stairs[i])
print(count)
print(*staircases)
