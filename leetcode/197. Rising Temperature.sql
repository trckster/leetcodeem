SELECT today.id
FROM Weather today
         LEFT JOIN Weather yesterday ON today.recordDate = ADDDATE(yesterday.recordDate, INTERVAL 1 DAY)
WHERE today.temperature > yesterday.temperature;