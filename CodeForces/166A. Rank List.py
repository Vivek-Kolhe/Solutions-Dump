n, k = map(int, input().split())
arr = []
for _ in range(n):
  arr.append(list(map(int, input().split())))
arr.sort(reverse=True, key=lambda x: (x[0], -x[1]))
count = 0
team = arr[k-1]
for ele in arr:
  if team[0] == ele[0] and team[1] == ele[1]:
    count += 1
print(count)
