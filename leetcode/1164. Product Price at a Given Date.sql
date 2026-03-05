SELECT DISTINCT p.product_id, COALESCE(latest.new_price, 10) as price FROM Products p
LEFT JOIN (
    SELECT product_id, new_price,
        ROW_NUMBER() OVER (PARTITION BY product_id ORDER BY change_date DESC) AS rn
    FROM Products
    WHERE change_date <= CAST('2019-08-16' AS DATE)
) latest
ON p.product_id = latest.product_id
WHERE latest.rn = 1 OR latest.rn IS NULL
ORDER BY p.product_id;
