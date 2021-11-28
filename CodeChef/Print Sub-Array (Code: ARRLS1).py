# cook your dish here
n = int(input())
arr = list(map(int, input().split()))
l, r = map(int, input().split())
print(*(arr[l - 1:r]))
