SELECT (
    CASE
        WHEN id % 2 = 0
            THEN id - 1
        WHEN id = (SELECT MAX(id) FROM Seat)
            THEN id
        ELSE
            id + 1
    END
) as id, student
FROM Seat
ORDER BY id;
