SELECT user_id,
       CONCAT(
               UPPER(LEFT(name, 1)),
               LOWER(RIGHT(name, CHAR_LENGTH(name) - 1))
       ) as name
FROM Users
ORDER BY user_id;