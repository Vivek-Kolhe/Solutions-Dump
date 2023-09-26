def comb(n, target):
  if n > target:
    return min(abs(n - target), abs(10 - n + target))
  return min(abs(n - target), abs(10 + n - target))

n = int(input())
ini = input()
final = input()
count = 0
for i in range(n):
  count += comb(int(ini[i]), int(final[i]))
print(count)
