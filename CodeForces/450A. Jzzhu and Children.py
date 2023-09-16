rounds = lambda count, n: count // n if count % n == 0 else (count // n) + 1
n, m = map(int, input().split())
arr = list(map(int, input().split()))
max_rounds = rounds(max(arr), m)
kids = []
for i in range(n):
  if rounds(arr[i], m) == max_rounds:
    kids.append(i)
print(kids[-1] + 1)
