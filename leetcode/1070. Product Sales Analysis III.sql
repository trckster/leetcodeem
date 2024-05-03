SELECT s1.product_id, s1.year as first_year, s1.quantity, s1.price FROM Sales as s1
    JOIN (SELECT product_id, MIN(year) as min_year FROM Sales GROUP BY product_id) as s2
WHERE s1.product_id=s2.product_id AND s1.year = s2.min_year;