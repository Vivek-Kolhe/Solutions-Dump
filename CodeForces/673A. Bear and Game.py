n = int(input())
arr = list(map(int, input().split()))
flag = 0
if arr[0] > 15:
    print(15)
else:
    if len(arr) == 1:
        print(arr[0] + 15)
    else:
        for i in range(n - 1):
            if arr[i + 1] - arr[i] > 15:
                print(arr[i] + 15)
                flag = 1
                break
        if not flag:
            if arr[-1] > 89:
                print(90)
            else:
                if arr[-1] + 15 > 90:
                    print(90)
                else:
                    print(arr[-1] + 15)
