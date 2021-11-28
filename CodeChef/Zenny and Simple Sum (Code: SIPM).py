# cook your dish here
for _ in range(int(input())):
    n = int(input())
    arr = list(set(list(map(int, input().split()))))
    pos, neg = 0, 0
    for i in range(len(arr)):
        if arr[i] > 0:
            pos += arr[i]
        else:
            neg += arr[i]
    print(pos, neg)
