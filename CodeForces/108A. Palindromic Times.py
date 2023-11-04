h, m = input().split(":")
if (int(h) >= 16 and int(h) <= 19) or (int(h) == 15 and int(m) == 51):
  print("20:02")
elif (int(h) >= 6 and int(h) <= 9) or (int(h) == 5 and int(m) == 50):
  print("10:01")
else:
  if int(h) == 23 and int(m) > 32:
    print("00:00")
  else:
    currp = h[::-1]
    if currp <= m:
      h = str(int(h) + 1)
      h = "0" + h if int(h) <= 9 else h
    print(h + ":" + h[::-1])
