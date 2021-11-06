for _ in range(int(input())):
  keyboard = input()
  word = input()
  time = 0
  if len(word) <= 1:
    print(time)
  else:
    for i in range(len(word) - 1):
      time += abs(keyboard.find(word[i]) - keyboard.find(word[i + 1]))
    print(time)
