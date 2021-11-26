n = int(input())
arr = list(map(int, input().split()))
count, new_arr = 0, []
for i in range(n):
  if arr[i] > 0:
    count += arr[i] - 1
    new_arr.append(1)
  elif arr[i] < 0:
    count += -(arr[i]) - 1
    new_arr.append(-1)
  else:
    new_arr.append(0)
if new_arr.count(-1) % 2 != 0:
  if 0 in new_arr:
    count += 1
    new_arr[new_arr.index(0)] = -1
  else:
    count += 2
    try:
      new_arr[new_arr.index(1)] = -1
    except ValueError:
      new_arr[new_arr.index(-1)] = -1
if 0 in new_arr:
  count += new_arr.count(0)
print(count)
