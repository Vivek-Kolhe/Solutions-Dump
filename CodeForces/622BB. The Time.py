hrs, mins = map(int, input().split(":"))
after = int(input())
mins += after % 60
hrs += after // 60 + mins // 60
mins %= 60
hrs %= 24
if hrs < 10:
  hrs = f"0{hrs}"
else:
  hrs = str(hrs)
if mins < 10:
  mins = f"0{mins}"
else:
  mins = str(mins)
print(f"{hrs}:{mins}")
