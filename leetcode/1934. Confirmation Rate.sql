SELECT Signups.user_id,
       ROUND(COUNT(CASE Confirmations.action WHEN 'confirmed' THEN 1 END) / COUNT(*), 2) as confirmation_rate
FROM Signups
         LEFT JOIN Confirmations ON Signups.user_id = Confirmations.user_id
GROUP BY Signups.user_id;