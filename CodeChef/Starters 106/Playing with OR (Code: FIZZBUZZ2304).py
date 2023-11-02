# cook your dish here res
for _ in range(int(input())):
    n, k = map(int, input().split())
    arr = list(map(int, input().split()))
    cnt = 0
    ind = -1
    found = False
    window = arr[:k]
    for i in range(k):
        if window[i] & 1:
            ind = i
            found = True
    if found:
        cnt += 1
    for i in range(k, n):
        if arr[i] & 1:
            ind = i
        if i - k + 1 <= ind <= i:
            cnt += 1
    print(cnt)
