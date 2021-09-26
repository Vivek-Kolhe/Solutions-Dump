# Enter your code here
def strike(string)
    return string.gsub(string, "<strike>#{string}</strike>")
end

def mask_article(string, arr)
    arr.each do |item|
        if string.include?(item)
            string.gsub!(item, strike(item))
        end
    end
    return string
end
