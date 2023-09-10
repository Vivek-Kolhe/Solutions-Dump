n = int(input())
arr = list(map(int, input().split()))
diff = max(arr)
for ele in arr:
  temp = ele if ele > 0 else ele * -1
  if diff > temp:
    diff = temp
print(diff if diff >= 0 else diff * -1 )
