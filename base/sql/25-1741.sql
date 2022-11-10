# Write your MySQL query statement below

SELECT event_day AS day,emp_id AS emp_id,sum(out_time-in_time) AS total_time
FROM Employees
GROUP BY event_day,emp_id;
