SELECT person_name FROM (
    SELECT person_name, weight, turn, SUM(weight) OVER (ORDER BY turn) as running_weight FROM Queue
) q
WHERE q.running_weight <= 1000
ORDER BY q.running_weight DESC
LIMIT 1;
