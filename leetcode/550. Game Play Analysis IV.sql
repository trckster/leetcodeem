SELECT ROUND(COUNT(a.player_id) / (SELECT COUNT(DISTINCT player_id) FROM Activity), 2) AS fraction
FROM (SELECT player_id, MIN(event_date) as min_date
      FROM Activity
      GROUP BY player_id) as a
WHERE DATE_ADD(a.min_date, INTERVAL 1 DAY) IN (SELECT event_date
                                               FROM Activity b
                                               WHERE a.player_id = b.player_id);