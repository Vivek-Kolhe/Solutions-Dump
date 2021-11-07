n = int(input())
arr = set(map(int, input().split()))
print(list(set(range(1, n + 1)) - arr)[0])