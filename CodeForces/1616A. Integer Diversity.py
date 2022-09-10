for _ in range(int(input())):
  n = int(input())
  arr = list(map(int, input().split()))
  uniq = list(set(arr))
  if n == len(uniq):
    print(n)
  else:
    new_arr = []
    for i in range(n):
      if arr[i] not in new_arr:
        new_arr.append(arr[i])
      elif arr[i] in new_arr:
        if -(arr[i]) not in new_arr:
          new_arr.append(-(arr[i]))
    print(len(new_arr))
