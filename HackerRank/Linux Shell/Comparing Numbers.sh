read X
read Y
if [[ $X < $Y ]]
then
    echo "X is less than Y"
elif [[ $X == $Y ]]
then
    echo "X is equal to Y"
else
    echo "X is greater than Y"
fi
