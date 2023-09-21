# cook your dish here
for _ in range(int(input())):
    n, p, pos = map(int, input().split())
    arr = sorted(list(map(int, input().split())))[::-1]
    delish = 0
    curr = 1
    while len(arr) > 0:
        if curr == pos:
            delish += arr[0]
        arr.pop(0)
        curr += 1
        if curr > p:
            curr = 1
    print(delish)
