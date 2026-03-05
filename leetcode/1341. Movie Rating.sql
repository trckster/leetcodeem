SELECT q1.name as results
FROM (
    SELECT name, COUNT(*) AS ratings_count
    FROM Users
    JOIN MovieRating
    USING (user_id)
    GROUP BY user_id
    ORDER BY ratings_count DESC, name
    LIMIT 1
) q1

UNION ALL

SELECT q2.title as results
FROM (
    SELECT title, AVG(MovieRating.rating) as avg_rating
    FROM Movies
    JOIN MovieRating
    USING (movie_id)
    WHERE YEAR(created_at) = 2020 AND MONTH(created_at) = 2
    GROUP BY movie_id
    ORDER BY avg_rating DESC, title
    LIMIT 1
) q2;
