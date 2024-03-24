for _ in range(int(input())):
    n, k = map(int, input().split())
    arr = list(map(int, input().split()))
    w = list(map(int, input().split()))
    arr.sort(reverse=True)
    w.sort()
    l, r = 0, n - 1
    breaker, s = k, 0
    ans = 0
    for i in range(k):
        if w[i] > 1:
            breaker = i
            break
        s += w[i]
        ans += (arr[l] * 2)
        l += 1
    arr, w = arr[s:], w[breaker:]
    arr.reverse()
    w.reverse()
    l, r = 0, len(arr) - 1
    for i in range(0, len(w)):
        ans += (arr[l] + arr[r])
        r -= 1
        l += 1
        l += (w[i] - 2)
    print(ans)
