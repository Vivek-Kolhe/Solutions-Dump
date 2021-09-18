n = int(input())
total = sum(list(map(int, input().split())))
arr = sum(list(range(1, n + 1)))
print(arr - total)
