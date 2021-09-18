# Your code here
def serial_average(s)
    temp = s.split("-")
    return "#{temp[0]}-#{((temp[1].to_f + temp[2].to_f) / 2).round(2)}"
end
