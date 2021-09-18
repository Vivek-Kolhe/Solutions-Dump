n, a, b = map(int, input().split())
arr = sorted(list(map(int, input().split())))
print(arr[b] - arr[b - 1])
