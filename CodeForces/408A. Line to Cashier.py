n = int(input())
arr = list(map(int, input().split()))
min_ = 99999
for i in range(n):
  count = 0
  temp = sum(list(map(int, input().split())))
  count += (temp * 5) + (arr[i] * 15)
  if min_ > count:
    min_ = count
print(min_)
