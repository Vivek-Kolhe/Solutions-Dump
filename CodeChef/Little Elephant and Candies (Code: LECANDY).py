for _ in range(int(input())):
    n, c = map(int, input().split())
    arr = list(map(int, input().split()))
    sum_ = sum(arr)
    if sum_ <= c:
        print("Yes")
    else:
        print("No")
