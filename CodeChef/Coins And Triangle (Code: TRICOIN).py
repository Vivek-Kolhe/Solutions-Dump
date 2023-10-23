# cook your dish here
for _ in range(int(input())):
    n = int(input())
    low, high = 1, n
    ans = n
    while low <= high:
        mid = (low + high) // 2
        s = mid * (mid + 1) // 2
        if s <= n:
            ans = mid
            low = mid + 1
        else:
            high = mid - 1
    print(ans)
