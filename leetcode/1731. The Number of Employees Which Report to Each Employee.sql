SELECT m.employee_id, m.name, COUNT(*) as reports_count, ROUND(AVG(e.age)) as average_age
FROM Employees m
         INNER JOIN Employees e ON m.employee_id = e.reports_to
GROUP BY m.employee_id, m.name
ORDER BY employee_id;