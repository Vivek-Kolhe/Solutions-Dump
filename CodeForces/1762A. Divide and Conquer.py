def parity_change(n):
  count = 0
  if n % 2:
    while n % 2:
      n = n // 2
      count += 1
  else:
    while not n % 2:
      n = n // 2
      count += 1
  return count
  
for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  if sum(arr) % 2 == 0:
    print(0)
  else:
    count = []
    for i in range(n):
      count.append(parity_change(arr[i]))
    print(min(count))
