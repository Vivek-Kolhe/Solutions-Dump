rats, women_child, men, captain, final = [], [], [], "", []
for _ in range(int(input())):
  name, crew = input().split()
  if crew == "woman" or crew == "child":
    women_child.append(name)
  elif crew == "man":
    men.append(name)
  elif crew == "rat":
    rats.append(name)
  else:
    captain = name
final = rats + women_child + men
final.append(captain)
for _ in range(len(final)):
  print(final[_])
