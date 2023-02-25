n, m = map(int, input().split())
count = 0
building = []
for _ in range(n):
  building.append(list(map(int, input().split())))
floor_pairs = []
for floor in building:
  temp = []
  for i in range(0, len(floor), 2):
    temp.append([floor[i], floor[i+1]])
  floor_pairs.append(temp)
for floor in floor_pairs:
  for flat in floor:
    if 1 in flat:
      count += 1
print(count)
