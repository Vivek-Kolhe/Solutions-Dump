# cook your dish here
for _ in range(int(input())):
    n = int(input())
    arr = list(map(int, input().split()))
    flag = False
    for i in range(n):
        if arr.count(arr[i]) > 1:
            flag = True
            print("ne krasivo")
            break
    if not flag:
        print("prekrasnyy")
