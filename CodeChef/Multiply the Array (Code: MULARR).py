# cook your dish here
for _ in range(int(input())):
    n = int(input())
    arr = list(map(int, input().split()))
    product = 1
    for i in range(n):
        product *= arr[i]
    print(product)
