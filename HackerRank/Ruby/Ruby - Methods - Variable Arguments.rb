# Your code here
def full_name(a, *rest)
    rest.reduce(a){|full, rest_| full + " " + rest_}
end
