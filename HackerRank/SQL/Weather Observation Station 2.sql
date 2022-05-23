/*
Enter your query here.
Please append a semicolon ";" at the end of the query and enter your query in a single line to avoid error.
*/
SELECT CAST(ROUND(SUM(LAT_N), 2) AS DECIMAL(10, 2)), CAST(ROUND(SUM(LONG_W), 2) AS DECIMAL(10, 2)) FROM STATION;
