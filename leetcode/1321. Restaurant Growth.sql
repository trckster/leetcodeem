WITH totals AS (
    SELECT visited_on,
        SUM(day_total) OVER (ORDER BY visited_on) as total
    FROM (
        SELECT visited_on, SUM(amount) as day_total
        FROM Customer
        GROUP BY visited_on
    ) _
)

SELECT t1.visited_on,
    t1.total - COALESCE(t2.total, 0) AS amount,
    ROUND((t1.total - COALESCE(t2.total, 0)) / 7, 2) AS average_amount
FROM totals t1
LEFT JOIN totals t2
ON t1.visited_on = DATE_ADD(t2.visited_on, INTERVAL 7 DAY)
LIMIT 18446744073709551615 OFFSET 5;
