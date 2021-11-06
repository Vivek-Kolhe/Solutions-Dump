read n
total = 0
for (( i = 0; i <= n; i++ ))
do
    read temp
    total=$(( total + temp ))
done
printf "%.3f" `echo "$total / $n" | bc -l`
