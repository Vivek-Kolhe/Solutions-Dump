def process_text(arr)
    new_arr = []
    arr.each{|item| new_arr.append(item.strip)}
    return new_arr.join(" ")
end
