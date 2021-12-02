n = int(input())
text = input().split()
arr = []
for i in range(len(text)):
  vol = 0
  for j in range(len(text[i])):
    if text[i][j].isupper():
      vol += 1
  arr.append(vol)
print(max(arr))
