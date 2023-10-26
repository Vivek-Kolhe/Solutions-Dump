months = [
  "January", 
  "February", 
  "March", 
  "April", 
  "May", 
  "June", 
  "July", 
  "August", 
  "September", 
  "October", 
  "November", 
  "December"
  ]
m = input()
n = int(input())
rem = n % 12
diff = rem + 1 - (12 - months.index(m))
print(months[diff-1])
