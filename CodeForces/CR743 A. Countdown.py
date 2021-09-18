for _ in range(int(input())):
    n = int(input())
    s = input()
    count = 0
    arr = [int(i) for i in s]
    count += arr[-1]
    arr[n-1] = 0
    for i in range(n-1):
        if arr[i] > 0:
            count += 1 + arr[i]
    print(count)
