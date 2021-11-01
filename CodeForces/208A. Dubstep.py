string = input()
word_arr = string.split("WUB")
for item in word_arr:
  if not item or " " in item:
    word_arr.remove(item)
print(" ".join(word_arr))
