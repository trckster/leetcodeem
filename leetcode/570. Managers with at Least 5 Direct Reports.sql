SELECT managers.name
FROM Employee managers
         LEFT JOIN Employee reports ON managers.id=reports.managerId
GROUP BY managers.id
HAVING COUNT(reports.id) >= 5;