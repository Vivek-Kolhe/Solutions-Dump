n = int(input())
string = input().split("W")
arr = [len(item) for item in string if item]
print(len(arr))
print(*arr)
