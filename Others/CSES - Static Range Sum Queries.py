n, m = map(int, input().split())
arr = list(map(int, input().split()))
pref = [0]
for i in range(n):
    pref.append(pref[-1] + arr[i])
for _ in range(m):
    l, r = map(int, input().split())
    print(pref[r] - pref[l-1])
