SELECT ROUND(COUNT(CASE WHEN d.order_date = d.customer_pref_delivery_date THEN 1 END) * 100 /
             COUNT(*), 2) AS immediate_percentage
FROM Delivery d
         JOIN (SELECT customer_id, MIN(order_date) AS min_date
               FROM Delivery
               GROUP BY customer_id) AS f
              ON d.customer_id = f.customer_id AND d.order_date = f.min_date;