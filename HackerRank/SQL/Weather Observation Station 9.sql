/*
Enter your query here.
Please append a semicolon ";" at the end of the query and enter your query in a single line to avoid error.
*/
SELECT DISTINCT CITY FROM STATION WHERE CITY NOT LIKE '[AEIOU]%' ORDER BY CITY ASC;
