for _ in range(int(input())):
  text = input()
  if text.startswith("miao.") and text.endswith("lala."):
    print("OMG>.< I don't know!")
  elif text.endswith("lala."):
    print("Freda's")
  elif text.startswith("miao."):
    print("Rainbow's")
  else:
    print("OMG>.< I don't know!")
