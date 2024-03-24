MOD = 1000000007
for _ in range(int(input())):
    n, k = map(int, input().split())
    arr = list(map(int, input().split()))
    s, maxs, curr = 0, 0, 0
    for i in range(n):
        s += arr[i]
        curr += arr[i]
        maxs = max(maxs, curr)
        if curr < 0:
            curr = 0
    s -= maxs
    for i in range(k):
        maxs += maxs
    print((s + maxs) % MOD)
