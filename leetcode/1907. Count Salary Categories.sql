SELECT 'Low Salary' as category, COUNT(1) as accounts_count
FROM Accounts
WHERE income < 20000

UNION ALL
SELECT 'Average Salary' as category, COUNT(1) as accounts_count
FROM Accounts
WHERE income >= 20000 AND income <= 50000

UNION ALL
SELECT 'High Salary' as category, COUNT(1) as accounts_count
FROM Accounts
WHERE income > 50000;
