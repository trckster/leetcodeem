SELECT uid as id, COALESCE(cnt1, 0) + COALESCE(cnt2, 0) as num
FROM (
    SELECT requester_id as uid, COUNT(1) AS cnt1
    FROM RequestAccepted
    GROUP BY requester_id
) q1 FULL JOIN (
    SELECT accepter_id as uid, COUNT(1) AS cnt2
    FROM RequestAccepted
    GROUP BY accepter_id
) q2
USING (uid)
ORDER BY num DESC
LIMIT 1;
