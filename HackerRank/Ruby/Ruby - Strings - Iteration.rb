# Your code here
def count_multibyte_char(string)
    count = 0
    string.each_char do |x|
        if x.bytesize > 1
            count += 1
        end
    end
    return count
end
