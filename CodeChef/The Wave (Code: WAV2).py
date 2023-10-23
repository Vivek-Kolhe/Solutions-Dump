# cook your dish here
n, q = map(int, input().split())
arr = list(map(int, input().split()))
arr.sort()
for _ in range(q):
    query = int(input())
    low, high = 0, n - 1
    ans = n
    while low <= high:
        mid = (low + high) // 2
        if query == arr[mid]:
            print(0)
            break
        elif arr[mid] >= query:
            ans = mid
            high = mid - 1
        else:
            low = mid + 1
    else:
        l, r = ans, n - ans
        if l % 2 == 0 and r % 2 == 0:
            print("POSITIVE")
        else:
            print("NEGATIVE")
